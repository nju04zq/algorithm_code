#include <stddef.h>
#include <string.h>

void
merge (int *n1, int cnt1, int *n2, int cnt2)
{
    int i, j, k, x, y;

    if (n1 == NULL || n2 == NULL) {
        return;
    }

    memmove(&n1[cnt2], &n1[0], cnt1 * sizeof(int));

    for (i = 0, j = 0, k = 0; i < cnt1 && j < cnt2;) {
        x = n1[cnt2+i];
        y = n2[j];
        if (x <= y) {
            n1[k++] = x;
            i++;
        } else {
            n1[k++] = y;
            j++;
        }
    }

    if (i == cnt1) {
        memcpy(&n1[k], &n2[j], (cnt2-j)*sizeof(int));
    } else if (j == cnt2) {
        memmove(&n1[k], &n1[cnt2+i], (cnt1-i)*sizeof(int));
    }

    return;
}

#include <stdio.h>

static void
dump_nums (int *p, int cnt)
{
    int i;

    for (i = 0; i < cnt; i++) {
        printf("%d ", p[i]);
    }
    printf("\n");
    return;
}

static void
test_merge (int *n1, int cnt1, int *n2, int cnt2)
{
    printf("Before merge #1: ");
    dump_nums(n1, cnt1);
    printf("Before merge #2: ");
    dump_nums(n2, cnt2);
    merge(n1, cnt1, n2, cnt2);
    printf("After  merge #1: ");
    dump_nums(n1, cnt1+cnt2);
    printf("\n");
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof(a[0]))

int main (void)
{
    int n1[100] = {};
    int n2[] = {1, 2, 3};
    int n3[100] = {1, 2, 3};
    int n4[] = {};
    int n5[100] = {1, 2, 3};
    int n6[] = {4, 5, 6};
    int n7[100] = {4, 5, 6};
    int n8[] = {1, 2, 3};
    int n9[100] = {1, 3, 5};
    int n10[] = {2, 4, 6};

    test_merge(n1, 0, n2, ARRAY_LEN(n2));
    test_merge(n3, 3, n4, ARRAY_LEN(n4));
    test_merge(n5, 3, n6, ARRAY_LEN(n6));
    test_merge(n7, 3, n8, ARRAY_LEN(n8));
    test_merge(n9, 3, n10, ARRAY_LEN(n8));
    return 0;
}

