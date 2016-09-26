#include <string.h>

typedef unsigned char bool;

#define TRUE 1
#define FALSE 0

#define SUDOKU_SIZE 9
#define SUKOKU_BLOCK_SIZE 3
#define SUDOKU_NULL_VAL '.'

typedef struct cord_s {
    int x;
    int y;
} cord_t;

static bool
check_val (char val, int *nums)
{
    int offset;

    if (val == SUDOKU_NULL_VAL) {
        return TRUE;
    }
    if (val < '1' || val > '9') {
        return FALSE;
    }
    offset = val - '1';
    if (nums[offset] > 0) {
        return FALSE;
    } else {
        nums[offset] = 1;
        return TRUE;
    }
}

static bool
validate_cords (char **boards, cord_t *cord_p)
{
    int i, nums[SUDOKU_SIZE];
    bool rc;
    char val;

    memset(nums, 0, sizeof(int)*SUDOKU_SIZE);

    for (i = 0; i < SUDOKU_SIZE; i++) {
        val = boards[cord_p[i].y][cord_p[i].x];
        rc = check_val(val, nums);
        if (rc != TRUE) {
            return FALSE;
        }
    }
    return TRUE;
}

static void
fill_cords_for_one_block (cord_t *cord_p, int block)
{
    int i, base_x, base_y;

    memset(cord_p, SUDOKU_SIZE*sizeof(cord_t), 0);

    base_x = block % SUKOKU_BLOCK_SIZE * SUKOKU_BLOCK_SIZE;
    base_y = block / SUKOKU_BLOCK_SIZE * SUKOKU_BLOCK_SIZE;

    for (i = 0; i < SUDOKU_SIZE; i++) {
        cord_p[i].x = base_x + i % SUKOKU_BLOCK_SIZE;
        cord_p[i].y = base_y + i / SUKOKU_BLOCK_SIZE;
    }
    return;
}

static bool
is_all_blocks_valid (char **board)
{
    int i;
    bool rc;
    cord_t cords[SUDOKU_SIZE];

    for (i = 0; i < SUDOKU_SIZE; i++) {
        fill_cords_for_one_block(cords, i);
        rc = validate_cords(board, cords);
        if (rc != TRUE) {
            return FALSE;
        }
    }
    return TRUE;
}

static void
fill_cords_for_one_col (cord_t *cord_p, int col)
{
    int i;

    memset(cord_p, SUDOKU_SIZE*sizeof(cord_t), 0);

    for (i = 0; i < SUDOKU_SIZE; i++) {
        cord_p[i].x = col;
        cord_p[i].y = i;
    }
    return;
}

static bool
is_all_cols_valid (char **board)
{
    int i;
    bool rc;
    cord_t cords[SUDOKU_SIZE];

    for (i = 0; i < SUDOKU_SIZE; i++) {
        fill_cords_for_one_col(cords, i);
        rc = validate_cords(board, cords);
        if (rc != TRUE) {
            return FALSE;
        }
    }
    return TRUE;
}

static void
fill_cords_for_one_row (cord_t *cord_p, int row)
{
    int i;

    memset(cord_p, SUDOKU_SIZE*sizeof(cord_t), 0);

    for (i = 0; i < SUDOKU_SIZE; i++) {
        cord_p[i].x = i;
        cord_p[i].y = row;
    }
    return;
}

static bool
is_all_rows_valid (char **board)
{
    int i;
    bool rc;
    cord_t cords[SUDOKU_SIZE];

    for (i = 0; i < SUDOKU_SIZE; i++) {
        fill_cords_for_one_row(cords, i);
        rc = validate_cords(board, cords);
        if (rc != TRUE) {
            return FALSE;
        }
    }
    return TRUE;
}

static bool
check_row_col (char **board, int row, int col)
{
    int i;

    if (row != SUDOKU_SIZE || col != SUDOKU_SIZE) {
        return FALSE;
    }

    for (i = 0; i < SUDOKU_SIZE; i++) {
        if (strlen(board[i]) != SUDOKU_SIZE) {
            return FALSE;
        }
    }
    return TRUE;
}

bool
isValidSudoku (char **board, int row, int col)
{
    bool rc;

    rc = check_row_col(board, row, col);
    if (rc != TRUE) {
        return FALSE;
    }

    rc = is_all_rows_valid(board);
    if (rc != TRUE) {
        return FALSE;
    }
    
    rc = is_all_cols_valid(board);
    if (rc != TRUE) {
        return FALSE;
    }
    
    rc = is_all_blocks_valid(board);
    if (rc != TRUE) {
        return FALSE;
    }

    return TRUE;
}

#include <stdio.h>

static void
dump_sudoku (char **board, int row, int col)
{
    int i, j;

    for (i = 0; i < row; i++) {
        for (j = 0; j < col; j++) {
            printf("%c ", board[i][j]);
        }
        printf("\n");
    }
    return;
}

static void
test_is_valid_sudoku(char **board, int row, int col)
{
    dump_sudoku(board, row, col);
    printf("****Is valid %d\n", isValidSudoku(board, row, col));
    return;
}

int main (void)
{
    char *a0[9] = {".87654321", "2........", "3........",
                   "4........", "5........", "6........",
                   "7........", "8........", "9........"};
    char *a1[9] = {".87654321", "2........", "3.8......",
                   "4........", "5........", "6........",
                   "7........", "8........", "9........"};
    char *a2[9] = {".87654321", "28.......", "3.8......",
                   "4........", "5........", "6........",
                   "7........", "8........", "9........"};
    char *a3[9] = {".87654321", "2........", "3......3.",
                   "4........", "5........", "6........",
                   "7........", "8........", "9........"};
    char *a4[9] = {".........", "......3..", "...18....",
                   "...7.....", "....1.97.", ".........",
                   "...36.1..", ".........", ".......2."};
    test_is_valid_sudoku(a4, 9, 9);
    test_is_valid_sudoku(a0, 9, 9);
    test_is_valid_sudoku(a1, 9, 9);
    test_is_valid_sudoku(a2, 9, 9);
    test_is_valid_sudoku(a3, 9, 9);
    test_is_valid_sudoku(a4, 9, 9);
    return 0;
}

