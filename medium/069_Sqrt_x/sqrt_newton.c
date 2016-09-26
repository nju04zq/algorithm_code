// x[n+1] = x[n] - f(x[n])/f'(x[n])
// x[n+1] = (x[n] + a/x[n])/2

#define MIN(a, b) ((a) < (b) ? (a) : (b))
#define INT_MAX_SQRT 46340

int
mySqrt (int y)
{
    int x;

    x = MIN(y, INT_MAX_SQRT);
    while (x * x > y) {
        x = (x + y/x)/2;
    }

    return x;
}

#include <math.h>
#include <stdio.h>
#include <limits.h>

int
test_sqrt (int x)
{
    int result, answer, rc = 0;

    result = mySqrt(x);
    answer = (int)sqrt((double)x);
    if (result != answer) {
        printf("Fail on %d, get %d, should be %d\n", x, result, answer);
        rc = -1;
    }
    return rc;
}

#define TEST_MAX_NUM 10000

int main (void)
{
    int i, rc;

    for (i = 0; i <= TEST_MAX_NUM; i++) {
        rc = test_sqrt(i);
        if (rc != 0) {
            break;
        }
    }
    test_sqrt(INT_MAX);
    return 0;
}
