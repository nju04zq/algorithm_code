#include <stddef.h>
#include <stdlib.h>

typedef struct array_2d_s {
    int **f;
    int *buf;
} array_2d_t;

static int
calloc_array_2d (array_2d_t *array_2d_p, int m, int n)
{
    int **f, *buf, i;

    f = calloc(m, sizeof(int *));
    if (f == NULL) {
        return -1;
    }
    buf = calloc(m*n, sizeof(int));
    if (f == NULL) {
        return -1;
    }

    array_2d_p->f = f;
    array_2d_p->buf = buf;

    for (i = 0; i < m; i++) {
        f[i] = buf;
        buf += n;
    }
    return 0;
}

static void
free_array_2d (array_2d_t *array_2d_p)
{
    free(array_2d_p->f);
    free(array_2d_p->buf);
    return;
}

// f[m, n] = f[m-1, n] + f[m, n-1]
static int
unique_path_internal (int **f, int m, int n)
{
    int i, j;

    for (i = 0; i < m; i++) {
        f[i][0] = 1;
    }
    for (j = 0; j < n; j++) {
        f[0][j] = 1;
    }

    for (i = 1; i < m; i++) {
        for (j = 1; j < n; j++) {
            f[i][j] = f[i-1][j] + f[i][j-1];
        }
    }

    return f[m-1][n-1];
}

int
uniquePaths (int m, int n)
{
    array_2d_t array_2d;
    int rc, result;

    rc = calloc_array_2d(&array_2d, m, n);
    if (rc != 0) {
        return -1;
    }

    result = unique_path_internal(array_2d.f, m, n);

    free_array_2d(&array_2d);
    return result;
}

#include <stdio.h>

static void
test_unique_path (int m, int n)
{
    int result;

    result = uniquePaths(m, n);
    printf("Unique path for {%d, %d}, %d\n", m, n, result);
    return;
}

int main (void)
{
    test_unique_path(1, 1);
    test_unique_path(1, 10);
    test_unique_path(3, 4);
    test_unique_path(4, 4);
    test_unique_path(5, 4);
    return 0;
}

