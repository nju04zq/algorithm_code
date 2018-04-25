class Solution(object):
    def findLadders(self, beginWord, endWord, wordList):
        """
        :type beginWord: str
        :type endWord: str
        :type wordList: List[str]
        :rtype: List[List[str]]
        """
        words = set(wordList)
        if endWord not in words:
            return []
        dist = 1
        toVisit, visited = [beginWord], {beginWord: dist}
        backtrace, found = {beginWord: []}, False
        while len(toVisit) > 0 and not found:
            cnt = len(toVisit)
            dist += 1
            for i in xrange(cnt):
                word = toVisit[i]
                for j in xrange(len(word)):
                    for ch in xrange(ord("a"), ord("z")+1):
                        candidate = word[:j] + chr(ch) + word[j+1:]
                        if candidate not in words:
                            continue
                        if candidate == endWord:
                            found = True
                        if candidate not in visited or visited[candidate] == dist:
                            if candidate in backtrace:
                                backtrace[candidate].append(word)
                            else:
                                backtrace[candidate] = [word]
                                toVisit.append(candidate)
                            visited[candidate] = dist
            toVisit = toVisit[cnt:]
        if not found:
            return []
        else:
            path, res = [], []
            self.dfs(backtrace, endWord, path, res)
            return res
    
    def dfs(self, backtrace, word, path, res):
        path.append(word)
        if len(backtrace[word]) == 0:
            res.append(path[::-1])
        for neighbor in backtrace[word]:
            self.dfs(backtrace, neighbor, path, res)
        path.pop()
