/*
 * the first step taking one stair, or two stairs
 * f1 = 1
 * f2 = 2
 * f(n) = f(n-1) + f(n-2), n>=3
 */
int
climbStairs (int n)
{
    int i, f, f1, f2;

    f1 = 1;
    f2 = 2;

    if (n <= 0) {
        return 0;
    } else if (n == 1) {
        return f1;
    } else if (n ==2 ) {
        return f2;
    }

    for (i = 3; i <= n; i++) {
        f = f1 + f2;
        f1 = f2;
        f2 = f;
    }
    return f;
}

#include <stdio.h>

static void
test_climb (int n)
{
    printf("n = %d, %d\n", n, climbStairs(n));
    return;
}

int main (void)
{
    test_climb(1);
    test_climb(2);
    test_climb(3);
    test_climb(4);
    test_climb(5);
    return 0;
}
