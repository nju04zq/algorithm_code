#include <limits.h>
#include <stddef.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

static int
max_crossing_mid (int *a, int n)
{
    int i, mid, sum, max_left, max_right;

    mid = n/2;

    sum = 0;
    max_left = INT_MIN;
    for (i = mid-1; i >= 0; i--) {
        sum += a[i];
        max_left = MAX(max_left, sum);
    }

    sum = 0;
    max_right = INT_MIN;
    for (i = mid; i < n; i++) {
        sum += a[i];
        max_right = MAX(max_right, sum);
    }
    return (max_left + max_right);
}

static int
max_subarray_internal (int *a, int n)
{
    int mid, max_sum1, max_sum2, max_sum3;

    if (n == 1) {
        return a[0];
    }

    mid = n/2;
    max_sum1 = max_subarray_internal(&a[0], mid);
    max_sum2 = max_subarray_internal(&a[mid], n-mid);
    max_sum3 = max_crossing_mid(a, n);
    return MAX(MAX(max_sum1, max_sum2), max_sum3);
}

int
maxSubArray (int *a, int n)
{
    int max_sum;

    if (a == NULL || n <= 0) {
        return 0;
    }

    max_sum = max_subarray_internal(a, n);
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

