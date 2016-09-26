class Solution(object):
    def init_words_src(self, words):
        words_src = {}
        for word in words:
            if word in words_src:
                words_src[word] += 1
            else:
                words_src[word] = 1
        return words_src

    def findSubstring(self, s, words):
        """
        :type s: str
        :type words: List[str]
        :rtype: List[int]
        """
        locations = []
        words_src = self.init_words_src(words)
        len_w = len(words[0])
        cnt_w = len(words)
        words_total_len = len_w * cnt_w
        for i in xrange(0, len(s)-len_w*cnt_w+1):
            is_valid = True
            words_summary = {}
            window_end = i + words_total_len
            for j in xrange(i, window_end, len_w):
                word = s[j:j+len_w]
                if word not in words:
                    is_valid = False
                    break
                if word not in words_summary:
                    words_summary[word] = 1
                else:
                    words_summary[word] += 1
                if words_summary[word] > words_src[word]:
                    is_valid = False
                    break
            if is_valid:
                locations.append(i)

        return locations

solution = Solution()
result = solution.findSubstring("barfoothefoobarman", ["foo", "bar"])
print result
result = solution.findSubstring("012aaaaabaabxyaaab", ["aaa", "aab", "aab"])
print result
result = solution.findSubstring("aaa", ["a", "b"])
print result
