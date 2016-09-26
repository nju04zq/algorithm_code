/*
 * https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
 */

#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

int
countPrimes (int n)
{
    int i, j, cnt;
    uint8_t *flags;

    flags = calloc(n, sizeof(uint8_t));
    if (flags == NULL) {
        return -1;
    }
    memset(flags, 1, (n+1)*sizeof(uint8_t));

    for (i = 2; i*i < n; i++) {
        if (flags[i] != 1) {
            continue;
        }
        for (j = 2; i*j < n; j++) {
            flags[i*j] = 0;
        }
    }

    for (i = 2, cnt = 0; i < n; i++) {
        cnt += flags[i];
    }

    free(flags);
    return cnt;
}

#include <stdio.h>

static void
test_count_prime (int n, int answer)
{
    int cnt;

    cnt = countPrimes(n);
    if (cnt != answer) {
        printf("Prime numbers in [1, %d], get %d, should be %d\n",
               n, cnt, answer);
    }
    return;
}

int main (void)
{
    test_count_prime(2, 0);
    test_count_prime(10, 4);
    test_count_prime(100, 25);
    test_count_prime(1000, 168);
    return 0;
}

