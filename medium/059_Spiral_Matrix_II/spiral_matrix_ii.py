class Solution(object):
    def get_next_num(self, num):
        next_num = num[0]
        num[0] += 1
        return next_num

    def generate_internal(self, a, i, n, num):
        if n == 0:
            return
        if n == 1:
            a[i][i] = self.get_next_num(num)
            num[0] += 1
            return

        start = i
        for j in xrange(start, start+n):
            a[i][j] = self.get_next_num(num)

        for i in xrange(start+1, start+n):
            a[i][j] = self.get_next_num(num)

        for j in xrange(j-1, start-1, -1):
            a[i][j] = self.get_next_num(num)

        for i in xrange(i-1, start, -1):
            a[i][j] = self.get_next_num(num)

        self.generate_internal(a, start+1, n-2, num)

    def generateMatrix(self, n):
        """
        :type n: int
        :rtype: List[List[int]]
        """
        if n == 0:
            return []

        num = [1]
        a = [[0 for i in xrange(n)] for i in xrange(n)] 
        self.generate_internal(a, 0, n, num)
        return a

def dump_array(a):
    print "-"*15
    for line in a:
        matrix_line = ""
        for num in line:
            matrix_line += "{:3}".format(num)
        else:
            print matrix_line
    print "-"*15

solution = Solution()

a = solution.generateMatrix(1)
dump_array(a)

a = solution.generateMatrix(2)
dump_array(a)

a = solution.generateMatrix(3)
dump_array(a)

a = solution.generateMatrix(4)
dump_array(a)

a = solution.generateMatrix(5)
dump_array(a)
