#include <stddef.h>

#define MIN(a, b) ((a) < (b) ? (a) : (b))

static int 
find_kth (int *n1, int size1, int *n2, int size2, int k)
{
    int pa, pb, a, la, lb;

    if (size1 > size2) {
        return find_kth(n2, size2, n1, size1, k);
    }

    if (size1 == 0) {
        return n2[k-1];
    }
    if (size2 == 0) {
        return n1[k-1];
    }
    if (k == 1) {
        return MIN(n1[0], n2[0]);
    }

    /*
     * d = size2-pb = size2+pa-k
     * if pa = k/2, size1 >= k/2, size1 <= size2, d = size2-k/2 >= 0
     * if pa = size1, size1 < k/2, d = size2+size1-k >= 0
     */
    pa = MIN(k/2, size1);
    pb = k - pa;
    if (n1[pa-1] < n2[pb-1]) {
        a = find_kth(&n1[pa], size1-pa, &n2[0], pb, k-pa);
    } else if (n1[pa-1] > n2[pb-1]) {
        a = find_kth(&n1[0], pa, &n2[pb], size2-pb, k-pb);
    } else {
        return n1[pa-1];
    }

    return a;
}

double
findMedianSortedArrays (int *n1, int size1, int *n2, int size2)
{
    int x, y;
    int total;

    if (n1 == NULL || n2 == NULL) {
        return -1;
    }

    if (size1 == 0 && size2 == 0) {
        return 0;
    }

    total = size1 + size2;
    if (total %2 == 0) {
        x = find_kth(n1, size1, n2, size2, total/2+1);
        y = find_kth(n1, size1, n2, size2, total/2);
        return ((double)(x+y))/2;
    } else {
        x = find_kth(n1, size1, n2, size2, total/2+1);
        return (double)(x);
    }
}

#include <stdio.h>
#include <stdlib.h>

#define ARRAY_SIZE_BASE 1
#define ARRAY_SIZE_RANGE 5

#define ARRAY_VAL_MAX 100

static int
decide_array_size (void)
{
    int x;

    x = rand();
    x = x % ARRAY_SIZE_RANGE + ARRAY_SIZE_BASE;
    return x;
}

#define SWAP(a, b) \
do {\
    int __tmp;\
    __tmp = (a);\
    (a) = (b);\
    (b) = __tmp;\
} while (0)

static int
quicksort_divide (int *s, int n)
{
    int i, j;

    for (i = 0, j = 0; i < n; i++) {
        if (s[i] <= s[0] && i != j) {
            j++;
            SWAP(s[i], s[j]);
        }
    }
    SWAP(s[0], s[j]);

    return j;
}

static void
quicksort (int *s, int n)
{
    int k;

    if (n <= 1) {
        return;
    }

    k = quicksort_divide(s, n);
    quicksort(&s[0], k);
    quicksort(&s[k+1], n-k-1);
    return;
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
    quicksort(*a, n);
    return 0;
}

static void
dump_array (int *s, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        printf("%d ", s[i]);
    }
    printf("\n");
    return;
}

static int
find_kth_bf (int *n1, int size1, int *n2, int size2, int k)
{
    int i, j, idx, a;

    for (i = 0, j = 0, idx = 0; i < size1 || j < size2;) {
        if (i >= size1) {
            a = n2[j++];
        } else if (j >= size2) {
            a = n1[i++];
        } else {
            a = (n1[i] < n2[j]) ? n1[i++] : n2[j++];
        }
        if (++idx == k) {
            return a;
        }
    }

    return -1;
}

static double
find_median_bf (int *n1, int size1, int *n2, int size2)
{
    int total, x, y;

    total = size1 + size2;
    if (total % 2 == 0) {
        x = find_kth_bf(n1, size1, n2, size2, total/2);
        y = find_kth_bf(n1, size1, n2, size2, total/2+1);
        return ((double)(x+y))/2;
    } else {
        x = find_kth_bf(n1, size1, n2, size2, total/2+1);
        return (double)x;
    }
}

static int
test_find_merged_median (void)
{
    double result, answer;
    int *n1, *n2, size1, size2, rc;

    rc = generate_array(&n1, &size1);
    if (rc != 0) {
        return -1;
    }
    rc = generate_array(&n2, &size2);
    if (rc != 0) {
        free(n1);
        return -1;
    }

    result = findMedianSortedArrays(n1, size1, n2, size2);
    answer = find_median_bf(n1, size1, n2, size2);

    printf("Array n1<%d>: ", size1);
    dump_array(n1, size1);
    printf("Array n2<%d>: ", size2);
    dump_array(n2, size2);
    printf("Get median %f\n", result);
    if ((int)(result) != (int)(answer)) {
        printf("answer should be %f\n", answer);
        rc = -1;
    } else {
        rc = 0;
    }

    free(n1);
    free(n2);
    return rc;
}

#define TEST_CASE_CNT 100

int main (void)
{
    int i, rc;

    for (i = 0; i < TEST_CASE_CNT; i++) {
        rc = test_find_merged_median();
        if (rc != 0) {
            break;
        }
        printf("##PASS test set %03d##\n", i);
    }

    return 0;
}

