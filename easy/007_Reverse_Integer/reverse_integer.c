#include <stdint.h>

typedef unsigned char bool;

#define TRUE 1
#define FALSE 0

// Not the accurate one, but enough to store int
#define INT_MAX_DIGIT_CNT 32

static uint32_t
decompose_digits (int x, int *digits)
{
    uint32_t i = 0;

    while (x > 0) {
        digits[i++] = x % 10;
        x /= 10;

        if (i >= INT_MAX_DIGIT_CNT) { //not possible, but in case
            return i;
        }
    }

    return i;
}

static int
compose_digits_reverse (int *digits, uint32_t digit_cnt)
{
    uint64_t x = 0;
    uint32_t i;

    for (i = 0; i < digit_cnt; i++) {
        x = x * 10 + digits[i];
        if (x & (0x1<<31)) { //overflow
            return -1;
        }
    }

    return (int)x;
}

int
reverse (int x)
{
    int digits[INT_MAX_DIGIT_CNT];
    uint32_t digit_cnt;
    bool is_negative = FALSE;

    if (x < 0) {
        is_negative = TRUE;
        x = -x;
    }

    digit_cnt = decompose_digits(x, digits);
    if (digit_cnt >= INT_MAX_DIGIT_CNT) {
        return 0;
    }

    x = compose_digits_reverse(digits, digit_cnt);
    if (x < 0) { //overflow
        return 0;
    }

    if (is_negative) {
        return -x;
    } else {
        return x;
    }
    
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
    reverse_test(123);
    reverse_test(100);
    reverse_test(-100);
    reverse_test(-1);
    reverse_test(0);
    reverse_test(-1481356204);
    reverse_test(1481356204);
    reverse_test(1534236469);
    return 0;
}

