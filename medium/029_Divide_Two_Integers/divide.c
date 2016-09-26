#include <stdint.h>
#include <limits.h>
#include "../../common/bool.h"

static int64_t
uabs (int64_t x)
{
    return x < 0 ? -x : x;
}

static void
shift_till (int64_t b, int64_t a, int64_t *times, int64_t *remain)
{
    int64_t x, y, z, x1;

    x = (int64_t)b;
    y = (int64_t)a;

    x1 = x;
    z = 1;
    for (;;) {
        x <<= 1;
        if (x > y) {
            break;
        }
        x1 = x;
        z <<= 1;
    }

    *times = z;
    *remain = y - x1;
    return;
}

int64_t
divide_internal (int64_t a, int64_t b)
{
    int64_t result, times, remain;

    result = 0;
    remain = a;
    while (remain > b) {
        shift_till(b, remain, &times, &remain);
        result += times;
    }
    if (remain == b) {
        result += 1;
    }
    return result;
}

int
divide (int a, int b)
{
    bool is_neg = FALSE;
    int64_t x, y, result;

    x = (int64_t)a;
    y = (int64_t)b;

    if (y == 0) {
        return INT_MAX;
    }
    if ((x>0 && y<0) || (x<0 && y>0)) {
        is_neg = TRUE;
    }

    x = uabs(x);
    y = uabs(y);

    result = divide_internal(x, y);

    result = is_neg ? -result : result;
    if (result > INT_MAX) {
        result = INT_MAX;
    }
    return result;
}

#include <stdio.h>
#include <stdlib.h>

#define RANDOM_RANGE 1000

static int
random_num (void)
{
    int x, y;

    x = random() % RANDOM_RANGE + 1;
    y = random() % 2;
    x = y == 0 ? x : -x;
    return x;
}

static int
test_divide (int x, int y)
{
    int z;
    int64_t x1, y1, z1;

    z = divide(x, y);

    x1 = (int64_t)x;
    y1 = (int64_t)y;
    z1 = x1/y1;
    if (z1 > INT_MAX) {
        z1 = INT_MAX;
    }

    if (z != z1) {
        printf("%d/%d, get %d, shoulde be %d\n", x, y, z, x/y);
        return -1;
    }
    return 0;
}

static int
test_divide_random (void)
{
    int x, y, rc;

    x = random_num();
    y = random_num();

    rc = test_divide(x, y);
    return rc;
}

#define TEST_CASE_CNT 1000

int main (void)
{
    int i, rc;

    for (i = 0; i < TEST_CASE_CNT; i++) {
        rc = test_divide_random();
        if (rc != 0) {
            printf("Fail on %d\n", i);
            break;
        }
    }

    test_divide(-1010369383, -2147483648);
    test_divide(-2147483648, -1);
    return 0;
}

