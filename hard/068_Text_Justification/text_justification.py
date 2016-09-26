class Solution(object):
    def add_one_line(self, result, words, start, end, max_width):
        word_num = end - start
        blank_num = word_num - 1 #blank number between words
        if blank_num == 0:
            blank_num = 1
        width = sum([len(w) for w in words[start:end]])
        spaces = max_width - width
        space_each = spaces/blank_num
        space_extra = spaces%blank_num
        content = ""
        for i in xrange(start, end):
            content += words[i]
            if word_num > 1 and i == end-1:
                continue
            content += " "*space_each
            if space_extra > 0:
                content += " "
                space_extra -= 1
        result.append(content)

    def add_last_line(self, result, words, start, end, max_width):
        content = ""
        for i in xrange(start, end):
            content += words[i]
            if i < end-1:
                content += " "
        spaces_extra = max_width - len(content)
        content += " "*spaces_extra
        result.append(content)

    def fullJustify(self, words, max_width):
        """
        :type words: List[str]
        :type maxWidth: int
        :rtype: List[str]
        """
        if len(words) == 0 or max_width <= 0:
            return [""]
        result = []
        start = 0
        prev_width = 0
        for i in xrange(len(words)):
            word_width = len(words[i])
            cur_width = prev_width + word_width
            if cur_width == max_width:
                prev_width = cur_width
                continue
            elif (cur_width+1) <= max_width:
                prev_width = cur_width + 1
                continue
            self.add_one_line(result, words, start, i, max_width)
            start = i
            prev_width = word_width + 1
        self.add_last_line(result, words, start, i+1, max_width)
        return result

def test_justify(words, max_width):
    result = solution.fullJustify(words, max_width)
    print max_width, words
    for line in result:
        print "#{}#".format(line)

solution = Solution()

a = ["This", "is", "an", "example", "of", "text", "justification."]
test_justify(a, 16)

a = ["12345"]
test_justify(a, 6)

a = ["12345", "1234"]
test_justify(a, 6)

a = [""]
test_justify(a, 0)

a = ["a"]
test_justify(a, 2)

a = ["a"]
test_justify(a, 1)

a = ["a", "b", "c"]
test_justify(a, 1)

a = ["a", "b", "c"]
test_justify(a, 3)

a = ["This", "is", "an", "example", "of", "text", "justification."]
test_justify(a, 14)

