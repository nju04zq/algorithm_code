#!/usr/bin/env python
# encoding: utf-8

import os
import re
import bs4
import sys
import codecs
import requests

data_fname = "codeforces"

link = "http://codeforces.com/problemset/page/"

class Problem(object):
    CSV_SEP = "#"

    def __init__(self, no, level, title, ptags, ac):
        self.no = no
        self.level = level
        self.title = title.replace("#", ",")
        self.ptags = ptags
        self.ac = ac
        try:
            int(self.ac)
        except:
            print "{0} {1} ac: {2}".format(no, level, ac)
            self.ac = "0"

    @staticmethod
    def load_csv(s):
        fields = s.split(Problem.CSV_SEP)
        tbl = {}
        tbl["no"] = fields[0]
        tbl["level"] = fields[1]
        tbl["title"] = fields[2]
        tbl["ac"] = fields[3]
        tbl["ptags"] = fields[4:]
        return tbl

    def csv(self):
        fields = [self.no, self.level, self.title, self.ac]
        for ptag in self.ptags:
            fields.append(ptag)
        return self.CSV_SEP.join(fields)

    def __repr__(self):
        return "{0}{1} {2} {3} {4}".format(self.no, self.level,
                                           self.title, self.ac, self.ptags)

def get_maxpage():
    r = requests.get(link + "1")
    s = r.text
    soup = bs4.BeautifulSoup(s, "html.parser")
    maxpage = 0
    res = soup.find_all("div", class_="pagination")
    res = res[0].find_all("span", class_="page-index")
    for tag in res:
        maxpage = max(maxpage, int(tag["pageindex"]))
    if maxpage == 0:
        raise Exception("Get 0 pages")
    return maxpage

def get_onepage(page):
    problems = []
    r = requests.get(link + str(page))
    s = r.text
    soup = bs4.BeautifulSoup(s, "html.parser")
    tags = soup.find_all("table", class_="problems")
    tags = tags[0].find_all("tr")
    for tag in tags[1:]:
        problem = get_oneproblem(tag)
        if problem is not None:
            problems.append(problem)
    return problems

def get_oneproblem(tag):
    #print str(tag)
    # problem no and level
    res = tag.find_all("td", class_="id")
    res = res[0].find_all("a")
    href = res[0]["href"] # "/problemset/problem/932/G"
    no, level = href.split("/")[-2:]
    # problem title and tags
    res = tag.find_all("div", style="float: left;")
    res = res[0].find_all("a")
    title = res[0].string.rstrip().lstrip()
    res = tag.find_all("div", style=re.compile("float: right;"))
    res = res[0].find_all("a")
    ptags = []
    for ptag in res:
       ptags.append(ptag.string)
    # ac
    res = tag.find_all("a", title="Participants solved the problem")
    if len(res) == 0:
        ac = "0"
    else:
        res = re.findall("(\d+)</a>", str(res[0]))
        ac = res[0]
    return Problem(no, level, title, ptags, ac)

def get_problems():
    maxpage = get_maxpage()
    problems = []
    for page in xrange(1, maxpage+1):
        sys.stdout.write("\r[{0}/{1}]".format(page, maxpage))
        sys.stdout.flush()
        problems += get_onepage(page)
    with codecs.open(data_fname, "w", encoding="utf-8") as fp:
        for problem in problems:
            fp.write(problem.csv()+"\n")
    sys.stdout.write("\n")
    return problems

def readin_problems():
    problems = []
    with codecs.open(data_fname, "r", encoding="utf-8") as fp:
        for line in fp.readlines():
            line = line.rstrip()
            params = Problem.load_csv(line)
            problems.append(Problem(**params))
    return problems

def load_problems():
    if not os.path.exists(data_fname):
        problems = get_problems()
    else:
        problems = readin_problems()
    return problems

def setup_db():
    problems = load_problems()
    db_ptags = {}
    for problem in problems:
        for ptag in problem.ptags:
            if ptag not in db_ptags:
                db_ptags[ptag] = {}
            level = problem.level[0]
            if level not in db_ptags[ptag]:
                db_ptags[ptag][level] = []
            db_ptags[ptag][level].append(problem)
    return problems, db_ptags

def show_ptags(problems, db_ptags):
    levels = set()
    for problem in problems:
        if len(problem.ptags) > 0:
            levels.add(problem.level[0])
    levels = sorted(list(levels))
    header = ["Tag", "Total"] + levels
    lines = []
    for ptag in db_ptags:
        line = [ptag]
        cnts = [0 for i in xrange(len(levels))]
        for level in db_ptags[ptag]:
            for i, header_level in enumerate(levels):
                if level == header_level:
                    cnts[i] += len(db_ptags[ptag][level])
        total = sum(cnts)
        line.append(total)
        for cnt in cnts:
            line.append(str(cnt))
        lines.append(line)
    lines.sort(key=lambda x:x[1], reverse=True)
    for line in lines:
        line[1] = str(line[1])
    PrettyTable(header, lines).show()

def query_db(problems, db_ptags, ptag):
    if ptag not in db_ptags:
        print "tag {0} not found".format(ptag)
        return
    ptag_data = db_ptags[ptag]
    levels = sorted(ptag_data.keys())
    header = ["#", "ID", "AC", "Title"]
    lines = []
    for level in levels:
        problems = sorted(ptag_data[level], key=lambda x:int(x.ac), reverse=True)
        for problem in problems:
            line = [level]
            line.append("{0}/{1}".format(problem.no, problem.level))
            line.append(problem.ac)
            line.append(problem.title[:51])
            lines.append(line)
    PrettyTable(header, lines).show()

class PrettyTable(object):
    def __init__(self, header, lines):
        self.header = header
        self.lines = lines
        self.col_limit = self.get_table_col_limit()
        # pad the seperator between columns
        self.col_seperator = "  "

    # print the whole table
    def show(self):
        sys.stdout.write(self.format())

    # format the whole table, return string
    def format(self):
        output = ""
        output += self.format_table_one_line(self.header)
        output += self.format_table_seperator()
        for oneline in self.lines:
            output += self.format_table_one_line(oneline)
        return output

    # calculate the width limit for each column in table
    def get_table_col_limit(self):
        self.lines.append(self.header)
        col_cnt = len(self.header)
        col_limit = [0 for i in xrange(col_cnt)]
        for line in self.lines:
            if len(line) != col_cnt:
                raise Exception("Table line {0} not match header {1}".format(\
                                line, self.header))
            for i in xrange(len(col_limit)):
                col_limit[i] = max(col_limit[i], len(line[i]))
        self.lines.pop()
        return col_limit

    # format one line in the table, each line is defined by a tuple containing
    # column values. If column value string length is less than the column width
    # limit, extra spaces will be padded
    def format_table_one_line(self, line):
        output = ""
        cols = []
        for i in xrange(len(line)):
            s = ""
            s += line[i]
            s += (" " * (self.col_limit[i]-len(line[i])))
            cols.append(s)
        output += (self.col_seperator.join(cols) + "\n")
        return output

    # format the seperator as -------
    def format_table_seperator(self):
        sep_cnt = sum(self.col_limit)
        # count in column seperators, why -1?, 2 columns only have one
        sep_cnt += (len(self.col_limit) - 1)*len(self.col_seperator)
        # one extra sep to make it pretty
        sep_cnt += 1
        return "-" * sep_cnt + "\n"

def main():
    problems, db_ptags = setup_db()
    if len(sys.argv) <= 1:
        show_ptags(problems, db_ptags)
    else:
        query_db(problems, db_ptags, sys.argv[1])

if __name__ == "__main__":
    reload(sys)
    sys.setdefaultencoding("utf-8")
    main()
