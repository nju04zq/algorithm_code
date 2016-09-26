#include <stdint.h>

typedef unsigned char bool;
#define TRUE 1
#define FALSE 0

static uint64_t
reverse_uint32 (uint32_t x)
{
    uint64_t tmp, y = 0;

    while (x > 0) {
        tmp = (uint64_t)(x % 10);
        y = y * 10 + tmp;
        x /= 10;
    }

    return y;
}

bool
isPalindrome (int x)
{
    uint64_t x_reversed;

    if (x < 0) {
        return FALSE;
    }

    x_reversed = reverse_uint32((uint32_t)x);
    if (x_reversed == (uint64_t)x) {
        return TRUE;
    } else {
        return FALSE;
    }
}

#include <stdio.h>

static void
test_check (int x)
{
    printf("%d, %d\n", x, isPalindrome(x));
    return;
}

int main (void)
{
    test_check(0);
    test_check(10);
    test_check(-1);
    test_check(123);
    test_check(121);
    test_check(1221);
    test_check(1000021);
    test_check(100001);
    return 0;
}

