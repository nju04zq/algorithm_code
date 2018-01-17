#!/usr/bin/env python
# encoding: utf-8

import os
import sys
import json
import time
import base64
import datetime
import optparse
import requests
import collections

CACHE_FPATH = "/tmp/803684751aa4b9dbd9c8853a6d683496"
STALE_AFTER = datetime.timedelta(7)

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

d = base64.b64decode
class OnelineProgress(object):
    def __init__(self):
        self.prev_line = None
        if sys.stdout.isatty():
            self.dump_progress = True
        else:
            self.dump_progress = False

    def clear_prev_line(self):
        space_cnt = len(self.prev_line)
        self.switch_to_line_head()
        sys.stdout.write(" " * space_cnt)

    def switch_to_line_head(self):
        sys.stdout.write("\r")

    def dump(self, msg):
        if not self.dump_progress:
            return
        if self.prev_line and len(msg) < len(self.prev_line):
            self.clear_prev_line()
        self.switch_to_line_head()
        sys.stdout.write(msg)
        sys.stdout.flush()
        self.prev_line = msg

def level_to_str(level):
    a = ["EASY", "MEDIUM", "HARD"]
    if level >= 1 and level <= len(a):
        return a[level-1]
    else:
        return "UNKNOWN({0})".format(level)

def acp_to_str(acp):
    return "{0:.1f}%".format(acp)

class Item(object):
    KEY_QID = "qid"
    KEY_NAME = "name"
    KEY_TITLE = "title"
    KEY_LEVEL = "level"
    KEY_AC = "ac"
    KEY_SUB = "sub"
    KEY_ACP = "acp"
    KEY_LIKE = "like"
    KEY_DISLIKE = "dislike"
    SORT_KEYS = [KEY_LIKE, KEY_DISLIKE, KEY_QID, KEY_SUB, KEY_ACP, KEY_LEVEL]
    SORT_REVERSE = [KEY_LIKE, KEY_SUB, KEY_ACP]
    ALL_KEYS = [KEY_NAME, KEY_TITLE, KEY_AC] + SORT_KEYS
    HEADERS = collections.OrderedDict([
        (KEY_QID, "ID"),
        (KEY_NAME, "Name"),
        (KEY_LEVEL, "Level"),
        (KEY_SUB, "SUB"),
        (KEY_ACP, "AC"),
        (KEY_LIKE, "Like"),
        (KEY_DISLIKE, "Dislike")
    ])
    COL_LIMIT = {
        KEY_NAME: 40,
    }
    COL_MAPPER = {
        KEY_LEVEL: level_to_str,
        KEY_ACP: acp_to_str
    } 

    def __init__(self, raw=None, saved=None):
        self.attribs = {}
        if raw is not None:
            self.init_from_raw(raw)
        else:
            self.init_from_saved(saved)

    def init_from_raw(self, data):
        a = self.attribs
        a[self.KEY_QID] = int(data["stat"]["frontend_question_id"])
        a[self.KEY_NAME] = data["stat"]["question__title"]
        a[self.KEY_TITLE] = data["stat"]["question__title_slug"]
        a[self.KEY_LEVEL] = int(data["difficulty"]["level"])
        a[self.KEY_AC]= int(data["stat"]["total_acs"])
        a[self.KEY_SUB] = int(data["stat"]["total_submitted"])
        a[self.KEY_ACP] = float(a[self.KEY_AC]) / float(a[self.KEY_SUB]) * 100.
        a[self.KEY_LIKE] = 0
        a[self.KEY_DISLIKE] = 0
        self.paidonly = data["paid_only"]

    def init_from_saved(self, data):
        self.attribs = data
        self.paidonly = False

    def get_sort_key(self, key):
        if key in self.SORT_REVERSE:
            return -key
        else:
            return key

def create_session():
    headers = {
            "Accept": "text/html",
            "Accept-Encoding": "gzip",
            "Accept-Language": "en",
            "Host": d("bGVldGNvZGUuY29t"),
            "User-Agent": "Mozilla/5.0",
            "Referer": d("aHR0cHM6Ly9sZWV0Y29kZS5jb20vYWNjb3VudHMvbG9naW4v")
    }
    s = requests.session()
    url = d("aHR0cHM6Ly9sZWV0Y29kZS5jb20vYWNjb3VudHMvbG9naW4v")
    res = s.get(url=url, headers=headers)
    getin_data = {}
    getin_data["csrfmiddlewaretoken"]=res.cookies["csrftoken"]
    getin_data[d("bG9naW4=")] = d("bmp1MDR6cUAxNjMuY29t")
    getin_data[d("cGFzc3dvcmQ=")] = d("RCNCd2ReM0U=")
    res = s.post(url, headers=headers, data=getin_data)
    return s, res.cookies["csrftoken"]

def get_item_raw_deep(session, headers, query, title):
    rurl = d("aHR0cHM6Ly9sZWV0Y29kZS5jb20vcHJvYmxlbXMvezB9L2Rlc2NyaXB0aW9uLw==")
    headers["referer"] = rurl.format(title)
    query["variables"]["titleSlug"] = title
    url = d("aHR0cHM6Ly9sZWV0Y29kZS5jb20vZ3JhcGhxbA==")
    r = session.post(url, headers=headers, data=json.dumps(query))
    data = json.loads(r.text)
    val = data["data"]["question"]["likesDislikes"]
    return int(val["likes"]), int(val["dislikes"])

def get_items_raw_deep(items):
    session, token = create_session()
    headers = {
        "authority":d("bGVldGNvZGUuY29t"),
        "method":"POST",
        "path":"/graphql",
        "scheme":"https",
        "accept":"*/*",
        "accept-encoding":"gzip",
        "accept-language":"en",
        "content-length":"280",
        "content-type":"application/json",
        "origin":d("aHR0cHM6Ly9sZWV0Y29kZS5jb20="),
        "user-agent":"Mozilla/5.0",
        "x-csrftoken":token
    }
    q = "{\"query\":\"query getLikesAndFavorites($titleSlug: String!) "\
        "{\\n  question(titleSlug: $titleSlug) {\\n    likesDislikes "\
        "{\\n      likes\\n      dislikes\\n    }\\n    isLiked\\n  }\\n  "\
        "favoritesLists\\n}\\n\",\"variables\":{\"titleSlug\":\"\"},"\
        "\"operationName\":\"getLikesAndFavorites\"}"
    query = json.loads(q)
    i = 0
    progress = OnelineProgress()
    for item in items:
        i += 1
        title = item.attribs[item.KEY_TITLE]
        like, dislike = get_item_raw_deep(session, headers, query, title)
        item.attribs[item.KEY_LIKE] = like
        item.attribs[item.KEY_DISLIKE] = dislike
        progress.dump("{0}/{1}".format(i, len(items)))
        time.sleep(1)
    progress.dump("")

def get_items_raw():
    url = d("aHR0cHM6Ly9sZWV0Y29kZS5jb20vYXBpL3Byb2JsZW1zL2FsZ29yaXRobXMv")
    headers = {
        "authority":d("bGVldGNvZGUuY29t"),
        "method":"GET",
        "path":"/api/problems/algorithms/",
        "scheme":"https",
        "accept":"application/json",
        "accept-encoding":"gzip",
        "accept-language":"en",
        "content-type":"application/json",
        "referer":d("aHR0cHM6Ly9sZWV0Y29kZS5jb20vcHJvYmxlbXNldC9hbGdvcml0aG1zLw=="),
        "user-agent":"Mozilla/5.0",
        "x-requested-with":"XMLHttpRequest"
    }
    try:
        r = requests.get(url, headers=headers)
    except:
        print "Fail to get items list, error:"
        raise
    else:
        return r.text

def get_items_from_raw():
    items = []
    s = get_items_raw()
    data = json.loads(s)
    for item_data in data["stat_status_pairs"]:
        item = Item(raw=item_data)
        if not item.paidonly:
            items.append(item)
    get_items_raw_deep(items)
    return items

def save_items(items):
    a = []
    for item in items:
        a.append(item.attribs)
    with open(CACHE_FPATH, "w") as fp:
        fp.write(base64.b64encode(json.dumps(a))[::-1])

def get_items_from_saved():
    with open(CACHE_FPATH, "r") as fp:
        s = fp.read()
    data = json.loads(base64.b64decode(s[::-1]))
    items = []
    for item_data in data:
        items.append(Item(saved=item_data))
    return items

def is_cache_fresh():
    try:
        stat_info = os.stat(CACHE_FPATH)
    except:
        return False
    mtime = datetime.datetime.fromtimestamp(stat_info.st_mtime)
    now = datetime.datetime.now()
    if mtime + STALE_AFTER > now:
        return True
    else:
        return False

def build_items_list():
    if env.force or not is_cache_fresh():
        items = get_items_from_raw()
        save_items(items)
    else:
        items = get_items_from_saved()
    return items

def filter_items(items):
    res = []
    for item in items:
        qid = item.attribs[item.KEY_QID]
        level = item.attribs[item.KEY_LEVEL]
        if env.fromqid > 0 and qid < env.fromqid:
            continue
        if env.toqid > 0 and qid > env.toqid:
            continue
        if len(env.levels) > 0 and level not in env.levels:
            continue
        res.append(item)
    return res

def make_sortkeys():
    keys, reverses, non_reverses = [], [], []
    for key in env.sortkeys:
        if key.startswith("-"):
            key = key[1:]
            reverses.append(key)
        else:
            non_reverses.append(key)
        keys.append(key)
    for key in Item.SORT_KEYS:
        if key not in keys:
            keys.append(key)
    for key in Item.SORT_REVERSE:
        if key not in non_reverses and key not in reverses:
            reverses.append(key)
    return keys, reverses

def sort_items(items, sort_keys, reverses):
    for key in sort_keys[::-1]:
        reverse = False
        if key in reverses:
            reverse = True
        items.sort(key=lambda item:item.attribs[key], reverse=reverse)

def show_items(items):
    keys = Item.HEADERS.keys()
    header = Item.HEADERS.values()
    lines = []
    for item in items:
        line = []
        for key in keys:
            if key in Item.COL_MAPPER:
                val = Item.COL_MAPPER[key](item.attribs[key])
            else:
                val = str(item.attribs[key])
            if key in Item.COL_LIMIT:
                val = val[:Item.COL_LIMIT[key]]
            line.append(val)
        lines.append(line)
    tbl = PrettyTable(header, lines)
    tbl.show()

def get():
    items = build_items_list()
    items = filter_items(items)
    sortkeys, reverses = make_sortkeys()
    sort_items(items, sortkeys, reverses)
    show_items(items)

class Env(object):
    def __init__(self):
        self.fromqid = -1
        self.toqid = -1
        self.levels = []
        self.sortkeys = []
        self.force = False

def add_parser_option(parser):
    parser.add_option("", "--from", dest="fromqid",
                      help="From qid")
    parser.add_option("", "--to", dest="toqid",
                      help="To qid")
    parser.add_option("", "--levels", dest="levels",
                      help="Only show levels, valid value 0, 1, 2")
    parser.add_option("", "--sortkeys", dest="sortkeys",
                      help="Sort according to the order of keys: {0}".format(\
                            Item.SORT_KEYS))
    parser.add_option("", "--force", action="store_true", dest="force",
                      help="Force update")

def set_parser_option_default(parser):
    parser.set_defaults(force=False)

def parse_levels(levels):
    res = []
    a = levels.replace(" ", "").split(",")
    for s in a:
       res.append(int(s))
    return res

def parse_sortkeys(sortkeys):
    keys = sortkeys.replace(" ", "").split(",")
    for key in keys:
        if key.startswith("-"):
            key = key[1:]
        if key not in Item.ALL_KEYS:
            raise Exception("Item does not have key {0}".format(key))
        elif key not in Item.SORT_KEYS:
            raise Exception("Does not support key {0} for sorting".format(key))
    return keys

def parse_options(options):
    if options.fromqid is not None:
        env.fromqid = int(options.fromqid)
    if options.toqid is not None:
        env.toqid = int(options.toqid)
    if options.levels is not None:
        env.levels = parse_levels(options.levels)
    if options.sortkeys is not None:
        env.sortkeys = parse_sortkeys(options.sortkeys)
    if options.force:
        env.force = True

def main():
    parser = optparse.OptionParser()
    add_parser_option(parser)
    set_parser_option_default(parser)
    options, args = parser.parse_args()
    parse_options(options)
    get()

env = Env()
main()
