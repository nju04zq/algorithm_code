#include <stddef.h>
#include <stdlib.h>
#include "../../common/bool.h"

static int **
alloc_array_2d (int m, int n)
{
    int **a, *buf, size1, size2, i;

    size1 = m*sizeof(int *);
    size2 = m*n*sizeof(int);
    a = calloc(1, size1 + size2);
    if (a == NULL) {
        return NULL;
    }

    buf = (int *)((unsigned long)a + size1);
    for (i = 0; i < m; i++) {
        a[i] = buf;
        buf += n;
    }

    return a;
}

static bool
validate_cell (char **board, int **track, int m, int n, int i, int j, char ch)
{

    if (i < 0 || j < 0 || i >= m || j >= n) {
        return FALSE;
    }
    if (board[i][j] != ch) {
        return FALSE;
    }
    if (track[i][j] != 0) {
        return FALSE;
    }
    return TRUE;
}

static bool
exist_internal (char **board, int **track, int m, int n,
                int i, int j, char *word)
{
    if (*word == '\0') {
        return TRUE;
    }

    if (validate_cell(board, track, m, n, i, j, *word) == FALSE) {
        return FALSE;
    }

    word++;
    track[i][j] = 1;

    if (exist_internal(board, track, m, n, i-1, j, word)) {
        return TRUE;
    }
    if (exist_internal(board, track, m, n, i, j-1, word)) {
        return TRUE;
    }
    if (exist_internal(board, track, m, n, i+1, j, word)) {
        return TRUE;
    }
    if (exist_internal(board, track, m, n, i, j+1, word)) {
        return TRUE;
    }

    track[i][j] = 0;
    return FALSE;
}

bool
exist (char **board, int m, int n, char *word)
{
    int **track, i, j;
    bool result = FALSE;

    if (board == NULL || m <= 0 || n <= 0 || word == NULL || word[0] == '\0') {
        return FALSE;
    }

    track = alloc_array_2d(m, n);
    if (track == NULL) {
        return FALSE;
    }

    for (i = 0; i < m; i++) {
        for (j = 0; j < n; j++) {
            result = exist_internal(board, track, m, n, i, j, word);
            if (result) {
                break;
            }
        }
        if (result) {
            break;
        }
    }

    free(track);
    return result;
}

#include <stdio.h>

char board1[3][4] = 
{
    {'A', 'B', 'C', 'E'},
    {'S', 'F', 'C', 'S'},
    {'A', 'D', 'E', 'E'},
};

static void
test_exist_1 (char *word)
{
    bool result;
    char *board_in[3];

    board_in[0] = board1[0];
    board_in[1] = board1[1];
    board_in[2] = board1[2];

    result = exist((char **)board_in, 3, 4, word);
    printf("%s exist, %d\n", word, result);
    return;
}

char board2[1][2] = 
{
    {'a', 'a'},
};

static void
test_exist_2 (char *word)
{
    bool result;
    char *board_in[1];

    board_in[0] = board2[0];

    result = exist((char **)board_in, 1, 2, word);
    printf("%s exist, %d\n", word, result);
    return;
}

int main (void)
{
    test_exist_1("ABCCED");
    test_exist_1("SEE");
    test_exist_1("ABCB");
    test_exist_2("aaa");
    return 0;
}


