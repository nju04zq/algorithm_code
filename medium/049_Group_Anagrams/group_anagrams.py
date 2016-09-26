class Solution(object):
    def calc_ascii_sum(self, word):
        ascii_sum = 0
        for ch in word:
            ascii_sum += ord(ch)
        return ascii_sum

    def calc_word_key(self, word):
        ascii_sum = self.calc_ascii_sum(word)
        return (len(word), ascii_sum)

    def is_word_anagram(self, a, b):
        ch_map = {}

        for ch in a:
            if ch_map.has_key(ch):
                ch_map[ch] += 1
            else:
                ch_map[ch] = 1

        for ch in b:
            if not ch_map.has_key(ch):
                return False
            if ch_map[ch] <= 0:
                return False
            ch_map[ch] -= 1
        return True

    def insert_word(self, groups, word):
        for group in groups:
            if self.is_word_anagram(group[0], word):
                group.append(word)
                break
        else:
            groups.append([word])

    def convert_anagram_map(self, anagram_map):
        result = []

        for groups in anagram_map.values():
            for group in groups:
                group.sort()
                result.append(group)
        return result

    def groupAnagrams(self, a):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        anagram_map = {}
        for word in a:
            key = self.calc_word_key(word)
            if not anagram_map.has_key(key):
                anagram_map[key] = [[word]]
            else:
                self.insert_word(anagram_map[key], word)

        result = self.convert_anagram_map(anagram_map)
        return result

solution = Solution()

a = ["eat", "tea", "tan", "ate", "nat", "bat"]
result = solution.groupAnagrams(a)
print result

a = ["cab","tin","pew","duh","may","ill","buy","bar","max","doc"]
result = solution.groupAnagrams(a)
print result
