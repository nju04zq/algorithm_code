class Solution(object):
    def numMatchingSubseq(self, s, words):
        """
        :type S: str
        :type words: List[str]
        :rtype: int
        """
        tbl = {}
        idxs = [0 for i in xrange(len(words))]
        for i, word in enumerate(words):
            ch = word[0]
            if ch in tbl:
                tbl[ch].append(i)
            else:
                tbl[ch] = [i]
        cnt = 0
        for ch in s:
            if ch not in tbl:
                continue
            vals = tbl[ch]
            tbl[ch] = []
            for i in vals:
                idxs[i] += 1
                if idxs[i] == len(words[i]):
                    cnt += 1
                    continue
                ch0 = words[i][idxs[i]]
                if ch0 in tbl:
                    tbl[ch0].append(i)
                else:
                    tbl[ch0] = [i]
        return cnt
