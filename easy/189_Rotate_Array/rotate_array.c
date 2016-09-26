#include <stddef.h>

#define SWAP(a, b) \
do {\
    int __tmp;\
    __tmp = a;\
    a = b;\
    b = __tmp;\
} while (0)

static void
reverse (int *a, int n)
{
    int i, j;

    for (i = 0, j = n - 1; i < j; i++, j--) {
        SWAP(a[i], a[j]);
    }
    return;
}

void
rotate (int *nums, int n, int k)
{
    if (nums == NULL || n <= 0 || k <= 0) {
        return;
    }

    k = k % n;
    if (k == 0) {
        return;
    }

    reverse(nums, n);
    reverse(&nums[0], k);
    reverse(&nums[k], n-k);
    return;
}

#include <stdio.h>

static void
dump_nums (int *nums, int n)
{
    int i;

    for (i = 0; i < n; i++) {
        printf("%d ", nums[i]);
    }
    return;
}

static void
test_rotate (int *nums, int n, int k)
{
    printf("Rotate %d\n", k);

    printf("Before:");
    dump_nums(nums, n);
    printf("\n");

    rotate(nums, n, k);

    printf("After:");
    dump_nums(nums, n);
    printf("\n");
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3, 4, 5};
    int a1[] = {1, 2, 3, 4, 5};
    int a2[] = {1, 2, 3, 4, 5};
    int a3[] = {1, 2, 3, 4, 5};
    int a4[] = {1, 2, 3, 4, 5};
    int a5[] = {1, 2, 3, 4, 5};
    int a6[] = {1, 2, 3, 4, 5, 6};

    test_rotate(a6, ARRAY_LEN(a6), 2);
    test_rotate(a0, ARRAY_LEN(a0), 0);
    test_rotate(a1, ARRAY_LEN(a1), 1);
    test_rotate(a2, ARRAY_LEN(a2), 2);
    test_rotate(a3, ARRAY_LEN(a3), 3);
    test_rotate(a4, ARRAY_LEN(a4), 4);
    test_rotate(a5, ARRAY_LEN(a5), 5);
    return 0;
}

