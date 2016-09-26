#include <limits.h>
#include <stddef.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

int
maxSubArray (int *a, int n)
{
    int i, prev, cur, max_sum;

    if (a == NULL || n <= 0) {
        return 0;
    }

    prev = 0;
    max_sum = INT_MIN;
    for (i = 0; i < n; i++) {
        if (prev < 0) {
            cur = a[i];
        } else {
            cur = prev + a[i];
        }
        max_sum = MAX(max_sum, cur);
        prev = cur;
    }

    return max_sum;
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
test_case (int *a, int n)
{
    int result;

    result = maxSubArray(a, n);
    dump_array(a, n);
    printf("Max subarray is %d\n", result);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1};
    int a1[] = {1, 2, 3};
    int a2[] = {-1, -2, -3};
    int a3[] = {-2, 1, -3, 4, -1, 2, 1, -5, 4};

    test_case(a0, ARRAY_LEN(a0));
    test_case(a1, ARRAY_LEN(a1));
    test_case(a2, ARRAY_LEN(a2));
    test_case(a3, ARRAY_LEN(a3));
    return 0;
}

