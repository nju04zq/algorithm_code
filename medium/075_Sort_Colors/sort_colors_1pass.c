#include <stddef.h>

void
sortColors (int *a, int n)
{
    int i, j, w = 0, b = 0, color;

    if (a == NULL || n <= 0) {
        return;
    }

    for (i = 0; i < n; i++) {
        color = a[i];
        if (color == 2) {
            continue;
        }
        a[i] = 2;
        j = b;
        b++;
        a[j] = 1;
        if (color == 1) {
            continue;
        }
        j = w;
        w++;
        a[j] = 0;
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
test_sort_color (int *a, int n)
{
    printf("before set: ");
    dump_array(a, n);
    sortColors(a, n);
    printf("after set: ");
    dump_array(a, n);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof(a[0]))

int main (void)
{
    int a0[] = {1, 0, 2};
    int a1[] = {0, 2};
    int a2[] = {2};
    int a3[] = {1, 0, 2, 1, 0, 2};
    int a4[] = {2, 2, 2};
    int a5[] = {1, 1, 1};
    int a6[] = {0, 0, 0};

    test_sort_color(a0, ARRAY_LEN(a0));
    test_sort_color(a1, ARRAY_LEN(a1));
    test_sort_color(a2, ARRAY_LEN(a2));
    test_sort_color(a3, ARRAY_LEN(a3));
    test_sort_color(a4, ARRAY_LEN(a4));
    test_sort_color(a5, ARRAY_LEN(a5));
    test_sort_color(a6, ARRAY_LEN(a6));
    return 0;
}


