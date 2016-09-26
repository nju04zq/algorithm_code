#include <stddef.h>
#include <stdlib.h>

#define SWAP_LINE(a, b) \
do {\
    int *__tmp;\
    __tmp = a;\
    a = b;\
    b = __tmp;\
} while (0)

int *
getRow (int row_index, int *returnSize)
{
    int max_col, *cur_line, *prev_line, i, j;

    *returnSize = 0;

    max_col = row_index + 1;
    cur_line = calloc(max_col, sizeof(int));
    if (cur_line == NULL) {
        return NULL;
    }
    prev_line = calloc(max_col, sizeof(int));
    if (prev_line == NULL) {
        free(cur_line);
        return NULL;
    }

    for (i = 0; i < row_index+1; i++) {
        SWAP_LINE(cur_line, prev_line);
        for (j = 0; j < i+1; j++) {
            if (j == 0 || j == i) {
                cur_line[j] = 1;
            } else {
                cur_line[j] = prev_line[j-1] + prev_line[j];
            }
        }
    }

    *returnSize = max_col;
    free(prev_line);
    return cur_line;
}

#include <stdio.h>

static void
test_get_row (int row)
{
    int *line, line_size, i;

    printf("The #%d line: ", row);

    line = getRow(row, &line_size);
    if (line == NULL) {
        printf("<NULL>\n");
        return;
    }

    for (i = 0; i < line_size; i++) {
        printf("%d ", line[i]);
    }
    printf("\n");

    free(line);
    return;
}

int main (void)
{
    test_get_row(0);
    test_get_row(1);
    test_get_row(2);
    test_get_row(3);
    test_get_row(4);
    test_get_row(5);
    test_get_row(6);
    return 0;
}

