class Solution(object):
    def validate_diagonal_1(self, board, i, j, n):
        if i > j:
            i = i - j
            j = 0
        else:
            j = j - i
            i = 0
        got_queen = False
        while i < n and j < n:
            if board[i][j] == "Q":
                if got_queen:
                    return False
                else:
                    got_queen = True
            i += 1
            j += 1
        return True

    def validate_diagonal_2(self, board, i, j, n):
        if i + j < (n-1):
            j = i + j
            i = 0
        else:
            i = (i+j) - (n-1)
            j = n-1
        got_queen = False
        while i < n and j >= 0:
            if board[i][j] == "Q":
                if got_queen:
                    return False
                else:
                    got_queen = True
            i += 1
            j -= 1
        return True

    def validate_vertical(self, board, j, n):
        got_queen = False
        for i in xrange(n):
            if board[i][j] != "Q":
                continue
            if got_queen:
                return False
            else:
                got_queen = True
        return True

    def validate_board(self, board, i, j, n):
        rc = self.validate_vertical(board, j, n)
        if rc == False:
            return False
        rc = self.validate_diagonal_1(board, i, j, n)
        if rc == False:
            return False
        rc = self.validate_diagonal_2(board, i, j, n)
        if rc == False:
            return False
        return True

    def n_queens_internal(self, board, i, result):
        n = len(board)
        if i == n:
            result[0] += 1
            return

        for j in xrange(n):
            board[i][j] = "Q"
            if self.validate_board(board, i, j, n) == True:
                self.n_queens_internal(board, i+1, result)
            board[i][j] = "."

    def totalNQueens(self, n):
        """
        :type n: int
        :rtype: List[List[str]]
        """
        result = [0]
        board = [["." for i in xrange(n)] for i in xrange(n)]
        self.n_queens_internal(board, 0, result)
        return result[0]

solution = Solution()

result = solution.totalNQueens(1)
print result

result = solution.totalNQueens(3)
print result

result = solution.totalNQueens(4)
print result

result = solution.totalNQueens(8)
print result

