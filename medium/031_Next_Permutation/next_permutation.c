#include <stddef.h>

#define SWAP(a, b) \
do {\
    int __tmp = a;\
    a = b;\
    b = __tmp;\
} while (0)

static void
reverse (int *a, int cnt)
{
    int i, j;

    for (i = 0, j = cnt-1; i < j; i++, j--) {
        SWAP(a[i], a[j]);
    }
    return;
}

static int
find_last_greater (int *a, int n, int j)
{
    int i;

    for (i = n-1; i > j; i--) {
        if (a[i] > a[j]) {
            break;
        }
    }
    return i;
}

void
nextPermutation (int *a, int n)
{
    int i, j;

    if (a == NULL || n <= 0) {
        return;
    }

    for (i = n-1, j = -1; i >= 1; i--) {
        if (a[i] > a[i-1]) {
            j = i-1;
            break;
        }
    }

    if (j == -1) {
        reverse(a, n);
    } else {
        i = find_last_greater(a, n, j);
        SWAP(a[i], a[j]);
        reverse(&a[j+1], n-j-1);
    }

    return;
}

#include <stdio.h>

static void
dump_array (int *a, int n)
{
    int i;

    for (i = 0; i < n; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
    return;
}

static void
test_next_permutation (int *a, int n)
{
    printf("Before next permutation: ");
    dump_array(a, n);
    nextPermutation(a, n);
    printf("After next permutation: ");
    dump_array(a, n);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3};
    int a1[] = {3, 2, 1};
    int a2[] = {1, 1, 5};
    int a3[] = {1, 3, 2};
    int a4[] = {1, 4, 3, 2};
    int a5[] = {2, 4, 3, 1};

    test_next_permutation(a0, ARRAY_LEN(a0));
    test_next_permutation(a1, ARRAY_LEN(a1));
    test_next_permutation(a2, ARRAY_LEN(a2));
    test_next_permutation(a3, ARRAY_LEN(a3));
    test_next_permutation(a4, ARRAY_LEN(a4));
    test_next_permutation(a5, ARRAY_LEN(a5));
    return 0;
}

