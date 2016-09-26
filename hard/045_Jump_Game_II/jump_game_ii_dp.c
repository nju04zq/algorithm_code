#include <stddef.h>
#include <stdlib.h>

#define MIN(a, b) ((a) < (b) ? (a) : (b))

int
jump (int *a, int n)
{
    int *f, result, i, j, min_jump;
    
    if (a == NULL || n == 0) {
        return 0;
    }

    f = calloc(n, sizeof(int));
    if (f == NULL) {
        return -1;
    }
    
    f[n-1] = 0;
    for (i = n-2; i >= 0; i--) {
        min_jump = -1;
        for (j = i+1; j <= i+a[i] && j < n; j++) {
            if (min_jump == -1) {
                min_jump = 1+f[j];
            } else {
                min_jump = MIN(min_jump, 1+f[j]);
            }
        }
        f[i] = min_jump;
    }

    result = f[0];
    free(f);
    return result;
}

#include <stdio.h>

static void
dump_array (int *a, int n)
{
    int i;

    for (i = 0; i < n; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
    return;
}

static void
test_jump (int *a, int n, int answer)
{
    int result;

    result = jump(a, n);
    if (result != answer) {
        dump_array(a, n);
        printf("Get jump %d, should jump %d\n", result, answer);
    }
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {2, 3, 1, 1, 4};

    test_jump(a0, ARRAY_LEN(a0), 2);
    return 0;
}

