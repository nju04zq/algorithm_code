class Solution(object):
    digits_map = {"2":"abc",  "3":"def", "4":"ghi", "5":"jkl", "6":"mno",
                  "7":"pqrs", "8":"tuv", "9":"wxyz"}

    def letter_combination_internal(self, digits, path, result):
        depth = len(path)
        if depth == len(digits):
            result.append("".join(path))
            return

        for ch in self.digits_map[digits[depth]]:
            path.append(ch)
            self.letter_combination_internal(digits, path, result)
            path.pop()

    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        if len(digits) == 0:
            return []
        path, result = [], []
        self.letter_combination_internal(digits, path, result)
        return result

def test_combination(digits):
    result = solution.letterCombinations(digits)
    print "combination for {}".format(digits)
    for s in result:
        print s
    print

solution = Solution()
test_combination("23")
test_combination("234")
