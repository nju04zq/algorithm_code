class Solution(object):
    def groupAnagrams(self, a):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        anagram_map = {}
        for word in a:
            key = "".join(sorted(list(word)))
            if not anagram_map.has_key(key):
                anagram_map[key] = [word]
            else:
                anagram_map[key].append(word)

        result = anagram_map.values()
        for group in result:
            group.sort()
        return result

solution = Solution()

a = ["eat", "tea", "tan", "ate", "nat", "bat"]
result = solution.groupAnagrams(a)
print result

a = ["cab","tin","pew","duh","may","ill","buy","bar","max","doc"]
result = solution.groupAnagrams(a)
print result
