#include <stddef.h>

int
firstMissingPositive (int *a, int n)
{
    int i, x;

    if (a == NULL || n == 0) {
        return 1;
    }

    for (i = 0; i < n;) {
        x = a[i];
        if (x <= 0 || (x-1) == i) {
            i++;
            continue;
        }
        if ((x-1) < i) {
            a[x-1] = x;
            i++;
            continue;
        }
        if ((x-1) >= n || a[x-1] == x) {
            i++;
            continue;
        }
        a[i] = a[x-1];
        a[x-1] = x;
    }

    for (i = 0; i < n; i++) {
        if ((a[i]-1) != i) {
            break;
        }
    }
    return i+1;
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
test_case (int *a, int n, int answer)
{
    int result;

    dump_array(a, n);
    result = firstMissingPositive(a, n);
    if (result != answer) {
        printf("Get result %d, should be %d\n", result, answer);
    } else {
        printf("Get result %d\n", result);
    }
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof(a[0]))

int main (void)
{
    int a0[] = {2, 3, 4, 5, -1, 1};
    int a1[] = {2, 3, 4, 5, -1, 3};
    int a2[] = {2, 3, 0, 5, -1, 1};
    int a3[] = {2, 2};
    int a4[] = {4};

    test_case(a0, ARRAY_LEN(a0), 6);
    test_case(a1, ARRAY_LEN(a1), 1);
    test_case(a2, ARRAY_LEN(a2), 4);
    test_case(a3, ARRAY_LEN(a3), 1);
    test_case(a4, ARRAY_LEN(a4), 1);
    return 0;
}

