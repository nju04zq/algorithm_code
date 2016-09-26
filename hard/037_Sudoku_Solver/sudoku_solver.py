# Two stage optimization
# #1 Raw backtracing(179ms)
#    ncalls API
#    4209/1 solve_internal
#    37652  validate
#
# #2 Pre-analyze & make board1(109ms)
#    ncalls API
#    4209/1 solve_internal
#    13544  validate
#
# #3 Set board if result in #2 has only one solution(18ms)
#    ncalls API
#    559/1  solve_internal
#    2259   validate

class Solution(object):
    MIN_NUM = 1
    MAX_NUM = 9

    def validate_one_group(self, group):
        mask = {}
        for x in group:
            if x == ".":
                continue
            if x in mask:
                return False
            mask[x] = 1
        return True

    def validate_one_row(self, board, i, j):
        row = board[i]
        valid = self.validate_one_group(row)
        return valid

    def get_col(self, board, j):
        col = []
        for line in board:
            col.append(line[j])
        return col

    def validate_one_col(self, board, i, j):
        col = self.get_col(board, j)
        valid = self.validate_one_group(col)
        return valid

# 0 1 2 3 4 5 6 7 8
# 0 0 0 3 3 3 6 6 6
# 3 3 3 6 6 6 9 9 9
    def get_block(self, board, i, j):
        i1 = i/3*3
        i2 = i1 + 3
        j1 = j/3*3
        j2 = j1 + 3
        block = []
        for i in xrange(i1, i2):
            for j in xrange(j1, j2):
                block.append(board[i][j])
        return block

    def validate_one_block(self, board, i, j):
        block = self.get_block(board, i, j)
        valid = self.validate_one_group(block)
        return valid

    def validate(self, board, i, j):
        valid = self.validate_one_row(board, i, j)
        if not valid:
            return False
        valid = self.validate_one_col(board, i, j)
        if not valid:
            return False
        valid = self.validate_one_block(board, i, j)
        if not valid:
            return False
        return True

    def get_pending_cell(self, board):
        for i in xrange(len(board)):
            line = board[i]
            for j in xrange(len(line)):
                if line[j] == ".":
                    return i, j
        return None

    def solve_internal(self, board, board1):
        cell = self.get_pending_cell(board)
        if cell is None:
            return True
        i, j = cell
        line = board[i]
        for k in board1[i][j]:
            line[j] = k
            valid = self.validate(board, i, j)
            if not valid:
                line[j] = "."
                continue
            success = self.solve_internal(board, board1)
            if success:
                break
            line[j] = "."
        else:
            return False
        return True

    def solve_pending_cell(self, board, i, j):
        result = []
        for k in xrange(self.MIN_NUM, self.MAX_NUM+1):
            board[i][j] = str(k)
            valid = self.validate(board, i, j)
            if valid:
                result.append(str(k))
        board[i][j] = "."
        return result

    def analyze_board(self, board):
        board1 = [[] for i in xrange(len(board))]
        for i in xrange(len(board)):
            for j in xrange(len(board[i])):
                if board[i][j] != ".":
                    board1[i].append([])
                    continue
                result = self.solve_pending_cell(board, i, j)
                board1[i].append(result)
                if len(result) == 1:
                    board[i][j] = result[0]
        return board1

    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        board1 = self.analyze_board(board)
        self.solve_internal(board, board1)

def transform_board(board):
    for line in board:
        for i in xrange(len(line)):
            if line[i] == 0:
                line[i] = "."
            else:
                line[i] = str(line[i])

def print_board(board):
    result = ""
    for line in board:
        for i in xrange(len(line)):
            result += "{} ".format(line[i])
        result += "\n"
    print result

def test_case(board):
    transform_board(board)
    solution.solveSudoku(board)
    print_board(board)

solution = Solution()
board = [[5, 3, 0, 0, 7, 0, 0, 0, 0,],
         [6, 0, 0, 1, 9, 5, 0, 0, 0,],
         [0, 9, 8, 0, 0, 0, 0, 6, 0,],
         [8, 0, 0, 0, 6, 0, 0, 0, 3,],
         [4, 0, 0, 8, 0, 3, 0, 0, 1,],
         [7, 0, 0, 0, 2, 0, 0, 0, 6,],
         [0, 6, 0, 0, 0, 0, 2, 8, 0,],
         [0, 0, 0, 4, 1, 9, 0, 0, 5,],
         [0, 0, 0, 0, 8, 0, 0, 7, 9,],]
test_case(board)
