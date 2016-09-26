#include <stdint.h>

int
trailingZeroes (int n)
{
    uint64_t x = 5;
    int cnt = 0;

    while (n >= x) {
        cnt += (n/x);
        x *= 5;
    }

    return cnt;
}

#include <stdio.h>

static void
test_trailing_zero (int n, int answer)
{
    int cnt;

    cnt = trailingZeroes(n);
    if (cnt != answer) {
        printf("Trailing zero for %d, get %d, should be %d\n", n, cnt, answer);
    }
    return;
}

int main (void)
{
    test_trailing_zero(100, 24);
    test_trailing_zero(200, 49);
    test_trailing_zero(2147483647, 536870902);
    return 0;
}

