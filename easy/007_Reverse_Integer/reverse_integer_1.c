#include <stdint.h>
#include <limits.h>

typedef unsigned char bool;

#define TRUE 1
#define FALSE 0

int
reverse (int x)
{
    bool is_negative = FALSE;
    int tmp;
    uint64_t y = 0;

    if (x < 0) {
        is_negative = TRUE;
        x = -x;
    }

    while (x > 0) {
        tmp = x % 10;
        y = y * 10 + tmp;
        x /= 10;
    }

    if (y>>31) {
        return 0;
    }

    if (is_negative) {
        y = -y;
    }

    return (int)y;
}

#include <stdio.h>

static void
reverse_test (int x)
{
    int y;

    y = reverse(x);
    printf("%d, %d\n", x, y);
    return;
}

int main (void)
{
//    reverse_test(123);
//    reverse_test(100);
//    reverse_test(-100);
//    reverse_test(-1);
//    reverse_test(0);
//    reverse_test(-1481356204);
//    reverse_test(1481356204);
    reverse_test(1534236469);
    return 0;
}

