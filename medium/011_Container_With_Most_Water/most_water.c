#include <stddef.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))
#define MIN(a, b) ((a) < (b) ? (a) : (b))

static int
calc_area (int *a, int i, int j)
{
    int d, area;

    d = (i > j) ? (i - j) : (j - i);
    area = MIN(a[i], a[j]) * d;
    return area;
}

int
maxArea (int *a, int n)
{
    int i, j, area, max_area = 0;

    if (a == NULL) {
        return 0;
    }

    for (i = 0, j = n-1; i < j;) {
        area = calc_area(a, i, j);
        max_area = MAX(area, max_area);
        if (a[i] < a[j]) {
            i++;
        } else {
            j--;
        }
    }

    return max_area;
}

#include <stdio.h>
#include <stdlib.h>

#define ARRAY_SIZE_BASE 1
#define ARRAY_SIZE_RANGE 50

#define ARRAY_VAL_MAX 100

static int
max_area_bf (int *a, int n)
{
    int i, j, area, max_area = 0;

    for (i = 0; i < n; i++) {
        for (j = 0; j < i; j++) {
            area = calc_area(a, i, j);
            max_area = MAX(area, max_area);
        }
    }
    return max_area;
}

static int
decide_array_size (void)
{
    int x;

    x = rand();
    x = x % ARRAY_SIZE_RANGE + ARRAY_SIZE_BASE;
    return x;
}

static void
rand_fill_array (int *s, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        s[i] = rand() % ARRAY_VAL_MAX;
    }
    return;
}

static int
generate_array (int **a, int *size)
{
    int n;

    *a = NULL;
    *size = 0;

    n = decide_array_size();
    *a = calloc(n, sizeof(int));
    if (*a == NULL) {
        return -1;
    }

    *size = n;

    rand_fill_array(*a, n);
    return 0;
}

static int
test_max_area (void)
{
    int *a, n, rc, max_area, answer;

    rc = generate_array(&a, &n);
    if (rc != 0) {
        return -1;
    }

    max_area = maxArea(a, n);
    answer = max_area_bf(a, n);

    rc = 0;
    if (max_area != answer) {
        rc = -1;
        printf("Get %d, answer %d\n", max_area, answer);
    }

    free(a);
    return rc;
}

#define TEST_CASE_CNT 100

int main (void)
{
    int i, rc;

    for (i = 0; i < TEST_CASE_CNT; i++) {
        rc = test_max_area();
        if (rc != 0) {
            printf("##Fail on test set %03d##\n", i);
            break;
        }
        printf("##PASS test set %03d##\n", i);
    }

    return 0;
}

