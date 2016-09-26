#include <stddef.h>
#include <stdlib.h>
#include <string.h>

/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *
plusOne (int *digits, int digitsSize, int *returnSize)
{
    int *digits_ret, i, flag, sum;

    *returnSize = 0;

    if (digits == NULL || digitsSize <= 0 || returnSize == NULL) {
        return NULL;
    }

    //Check each number in digits, TODO

    digits_ret = calloc(digitsSize+1, sizeof(int));
    if (digits_ret == NULL) {
        return NULL;
    }

    flag = 1;
    for (i = digitsSize-1; i >= 0; i--) {
        sum = digits[i] + flag;
        if (sum >= 10) {
            sum -= 10;
            flag = 1;
        } else {
            flag = 0;
        }
        digits_ret[i] = sum;
    }

    if (flag == 0) {
        *returnSize = digitsSize;
        return digits_ret;
    }

    memmove(&digits_ret[1], &digits_ret[0], digitsSize*sizeof(int));
    digits_ret[0] = 1;
    *returnSize = digitsSize + 1;
    return digits_ret;
}

#include <stdio.h>

static void
dump_digits (int *a, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        printf("%d", a[i]);
    }
    return;
}

static void
test_plusone (int *a, int size)
{
    int *ret, ret_size;

    ret = plusOne(a, size, &ret_size);
    dump_digits(a, size);
    printf(" + 1 = ");
    dump_digits(ret, ret_size);
    printf("\n");
    return;
}

int main (void)
{
    int a0[] = {0, 0, 0};
    int a1[] = {0, 0, 9};
    int a2[] = {0, 9, 9};
    int a3[] = {9, 9, 9};
    int a4[] = {0};
    int a5[] = {9};

    test_plusone(a0, 3);
    test_plusone(a1, 3);
    test_plusone(a2, 3);
    test_plusone(a3, 3);
    test_plusone(a4, 1);
    test_plusone(a5, 1);
    return 0;
}

