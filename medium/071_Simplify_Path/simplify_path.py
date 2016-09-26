class Solution(object):
    def get_next_dir_name(self, path, i, max_len):
        while i < max_len and path[i] == "/":
            i += 1
        if i == max_len:
            return max_len, 0

        start = i
        while i < max_len and path[i] != "/":
            i += 1
        name_len = (i - start)
        return start, name_len

    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        max_len = len(path)
        if max_len == 0 or path[0] != "/":
            return ""
        stack = []
        i = 0
        while i < max_len:
            i, name_len = self.get_next_dir_name(path, i, max_len)
            if i == max_len:
                break

            dir_name = path[i:i+name_len]
            if dir_name == "..":
                if len(stack) > 0:
                    stack.pop()
            elif dir_name != ".":
                stack.append(dir_name)
            i += name_len

        result = ""
        for dir_name in stack:
            result += "/{}".format(dir_name)
        if result == "":
            result = "/"
        return result

def test_simplify(path, answer):
    result = solution.simplifyPath(path)
    if result != answer:
        print "{}, get {}, should be {}".format(path, result, answer)

solution = Solution()
test_simplify("/", "/")
test_simplify("/abc", "/abc")
test_simplify("/abc/", "/abc")
test_simplify("/abc//", "/abc")
test_simplify("/abc/123", "/abc/123")
test_simplify("/abc//123", "/abc/123")
test_simplify("/a/./b/../../c/", "/c")
test_simplify("/a/./b/../../../c/", "/c")
test_simplify("/a/./b/../../../", "/")
