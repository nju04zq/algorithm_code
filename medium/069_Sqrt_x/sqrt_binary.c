#include <stdint.h>

int
mySqrt (int i)
{
    int64_t x, start, mid, end, temp;

    x = (int64_t)i;
    if (x < 0) {
        return -1;
    }
    if (x == 0 || x == 1) {
        return x;
    }

    start = 0;
    end = x;

    while (start < end) {
        mid = (start + end)/2;
        temp = mid * mid;
        if (temp == x) {
            return mid;
        } else if (temp < x) {
            start = mid;
        } else {
            end = mid;
        }
        if (end == (start+1)) {
            break;
        }
    }

    return start;
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

