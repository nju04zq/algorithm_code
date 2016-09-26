#define MAX(a, b) ((a) > (b) ? (a) : (b))

/*
 * f(0) = a[0]
 * f(1) = max(a[0], a[1])
 * f(n) = max(a[n]+f(n-2), f(n-1))
 */
int
rob (int *a, int n)
{
    int i, f_n, f_n1, f_n2;

    f_n2 = a[0];
    f_n1 = MAX(a[0], a[1]);

    if (n == 0) {
        return 0;
    } else if (n == 1) {
        return f_n2;
    } else if (n == 2) {
        return f_n1;
    }

    for (i = 3; i <= n; i++) {
        f_n = MAX(a[i-1]+f_n2, f_n1);
        f_n2 = f_n1;
        f_n1 = f_n;
    }

    return f_n;
}

#include <stdio.h>

static void
test_rob (int *a, int n)
{
    int i, result;

    result = rob(a, n);
    
    for (i = 0; i < n; i++) {
        printf("%d ", a[i]);
    }
    printf(", %d\n", result);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3, 4, 5};
    int a1[] = {1, 3, 5, 2, 4};

    test_rob(a0, ARRAY_LEN(a0));
    test_rob(a1, ARRAY_LEN(a1));
    return 0;
}

