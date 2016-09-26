#include <stdint.h>

int
hammingWeight (uint32_t n)
{
    int cnt = 0;

    for (; n != 0; n = n >> 1) {
        cnt += (n & 0x1);
    }
    return cnt;
}

#include <stdio.h>

static void
test_hamming_weight (uint32_t x, int answer)
{
    int cnt;

    cnt = hammingWeight(x);
    if (cnt != answer) {
        printf("Hamming weight of %x, get %d, should be %d", x, cnt, answer);
    }
    return;
}

int main (void)
{
    test_hamming_weight(0x0, 0);
    test_hamming_weight(0x1, 1);
    test_hamming_weight(0xc, 2);
    test_hamming_weight(0xc0000000, 2);
    test_hamming_weight(0xffffffff, 32);
    return 0;
}

