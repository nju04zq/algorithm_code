#include <stddef.h>
#include <stdlib.h>

static void
free_all_output (int **output, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        if (output[i]) {
            free(output[i]);
        }
    }
    return;
}

static int **
alloc_output (int rows)
{
    int **output, i;

    output = calloc(rows, sizeof(int *));
    if (output == NULL) {
        return NULL;
    }

    for (i = 0; i < rows; i++) {
        output[i] = calloc(i+1, sizeof(int));
        if (output[i] == NULL) {
            free_all_output(output, rows);
            free(output);
            return NULL;
        }
    }

    return output;
}

static int *
alloc_column_sizes (int rows)
{
    int *columns, i;

    columns = calloc(rows, sizeof(int));
    if (columns == NULL) {
        return NULL;
    }

    for (i = 0; i < rows; i++) {
        columns[i] = i+1;
    }

    return columns;
}

// C(n, k) = C(n-1, k) + C(n-1, k-1)
int **
generate (int numRows, int **columnSizes, int *returnSize)
{
    int **output, i, j;

    *columnSizes = NULL;
    *returnSize = 0;

    if (numRows == 0) {
        return NULL;
    }

    output = alloc_output(numRows);
    if (output == NULL) {
        return NULL;
    }
    *columnSizes = alloc_column_sizes(numRows);
    if (*columnSizes == NULL) {
        free(output);
        return NULL;
    }
    *returnSize = numRows;

    for (i = 0; i < numRows; i++) {
        for (j = 0; j < i+1; j++) {
            if (j == 0 || j == i) {
                output[i][j] = 1;
            } else {
                output[i][j] = output[i-1][j-1] + output[i-1][j];
            }
        }
    }

    return output;
}

#include <stdio.h>

static void
print_spaces (int cnt)
{
    int i;

    for (i = 0; i < cnt; i++) {
        printf(" ");
    }
    return;
}

static void
dump_pascal_triangle (int **output, int *columns, int output_size)
{
    int i, j;

    for (i = 0; i < output_size; i++) {
        print_spaces(output_size - i - 1);
        for (j = 0; j < columns[i]; j++) {
            printf("%d ", output[i][j]);
        }
        printf("\n");
    }
    return;
}

static void
test_generate (int rows)
{
    int **output, *columns, output_size;

    printf("Pascal triangle for rows %d:\n", rows);

    output = generate(rows, &columns, &output_size);
    if (output == NULL) {
        printf("<NULL>\n");
        return;
    }

    dump_pascal_triangle(output, columns, output_size);

    free(output);
    free(columns);
    return;
}

int main (void)
{
    test_generate(0);
    test_generate(1);
    test_generate(2);
    test_generate(3);
    test_generate(4);
    test_generate(5);
    test_generate(6);
    return 0;
}

