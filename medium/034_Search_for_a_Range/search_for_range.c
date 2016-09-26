#include <stddef.h>
#include <stdlib.h>

static void
search_range_bf (int *a, int start, int end, int target, int *result)
{
    int i, left, right;

    left = -1;
    right = -1;
    for (i = start; i <= end; i++) {
        if (a[i] == target) {
            if (left == -1) {
                left = i;
            }
            right = i;
        }
    }

    result[0] = left;
    result[1] = right;
    return;
}

static void
search_range_internal (int *a, int start, int end, int target, int *result)
{
    int ltemp[2], rtemp[2], mid, n;

    n = end - start + 1;
    if (n <= 2) {
        search_range_bf(a, start, end, target, result);
        return;
    }

    mid = (start+end)/2;
    if (a[mid] < target) {
        search_range_internal(a, mid+1, end, target, result);
        return;
    } else if (a[mid] > target) {
        search_range_internal(a, start, mid-1, target, result);
        return;
    }

    search_range_internal(a, start, mid-1, target, ltemp);
    search_range_internal(a, mid+1, end, target, rtemp);
    result[0] = mid;
    result[1] = mid;
    if (ltemp[0] != -1) {
        result[0] = ltemp[0];
    }
    if (rtemp[1] != -1) {
        result[1] = rtemp[1];
    }
    return;
}


int *
searchRange (int *a, int n, int target, int *returnSize)
{
    int *result;

    *returnSize = 0;
    if (a == NULL || n <= 0) {
        return NULL;
    }

    result = calloc(2, sizeof(int));
    if (result == NULL) {
        return NULL;
    }
    *returnSize = 2;

    search_range_internal(a, 0, n-1, target, result);

    return result;
}

#include <stdio.h>
#include <stdlib.h>

#define ARRAY_LEN_BASE 5
#define ARRAY_LEN_RANGE 10
#define ARRAY_MAX_ELEMENT 20

static int
generate_random_num (int max_val)
{
    int x;

    x = random() % max_val;
    return x;
}

static int *
generate_random_array (int n)
{
    int *a, i;

    a = calloc(n, sizeof(int));
    if (a == NULL) {
        return NULL;
    }

    for (i = 0; i < n; i++) {
        a[i] = random() % ARRAY_MAX_ELEMENT;
    }
    return a;
}

static int
quicksort_split (int *a, int n)
{
    int i, j, pilot;

    pilot = a[0];
    for (i = 1, j = 0; i < n; i++) {
        if (a[i] <= pilot) {
            a[j++] = a[i];
            a[i] = a[j];
        }
    }
    a[j] = pilot;
    return j;
}

static void
quicksort (int *a, int n)
{
    int i;

    if (n <= 1) {
        return;
    }

    i = quicksort_split(a, n);
    quicksort(&a[0], i);
    quicksort(&a[i+1], n-i-1);
    return;
}

static int *
generate_test_array (int *len)
{
    int n, *a;

    n = generate_random_num(ARRAY_LEN_RANGE);
    n += ARRAY_LEN_BASE;
    a = generate_random_array(n);
    if (a == NULL) {
        return NULL;
    }

    quicksort(a, n);

    *len = n;
    return a;
}

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

static int
test_search (int i)
{
    int *result, answer[2], size, *a, n, target, rc;

    a = generate_test_array(&n);
    if (a == NULL) {
        return -1;
    }

    target = generate_random_num(ARRAY_MAX_ELEMENT);

    rc = 0;
    result = searchRange(a, n, target, &size);
    if (result == NULL) {
        free(a);
        return -1;
    }

    search_range_bf(a, 0, n-1, target, answer);
    if (result[0] != answer[0] || result[1] != answer[1]) {
        printf("Search %d in array: ", target);
        dump_array(a, n);
        printf("Get (%d, %d), should be (%d, %d)\n",
               result[0], result[1], answer[0], answer[1]);
        rc = -1;
    }

    free(a);
    return rc;
}

#define TEST_CASE_CNT 1000

int main (void)
{
    int i, rc;

    for (i = 0; i < TEST_CASE_CNT; i++) {
        rc = test_search(i);
        if (rc != 0) {
            printf("Fail on case ##%d##\n", i);
            break;
        }
    }
    return 0;
}
