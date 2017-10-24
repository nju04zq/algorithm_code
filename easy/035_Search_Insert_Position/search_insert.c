#include <stddef.h>

static int
search_insert_bf (int *a, int start, int end, int target)
{
    int i;

    for (i = start; i <= end; i++) {
        if (a[i] >= target) {
            break;
        }
    }
    return i;
}

static int
search_insert_internal (int *a, int start, int end, int target)
{
    int pos, mid, n;

    n = end - start + 1;
    if (n <=2 ) {
        pos = search_insert_bf(a, start, end, target);
        return pos;
    }

    mid = (start+end)/2;
    if (a[mid] == target) {
        pos = mid;
    } else if (a[mid] < target) {
        pos = search_insert_internal(a, mid+1, end, target);
    } else {
        pos = search_insert_internal(a, start, mid-1, target);
    }
    return pos;
}

int
searchInsert (int *a, int n, int target)
{
    int pos;

    if (a == NULL || n <= 0) {
        return -1;
    }

    pos = search_insert_internal(a, 0, n-1, target);
    return pos;
}

#include <stdio.h>
#include <stdlib.h>

#define ARRAY_LEN_BASE 5
#define ARRAY_LEN_RANGE 10
#define ARRAY_MAX_ELEMENT 120

static int
generate_random_num (int max_val)
{
    int x;

    x = random() % max_val;
    return x;
}

static int *
generate_no_dup_array (int n)
{
    int *mask, *a, i, j;

    mask = calloc(ARRAY_MAX_ELEMENT, sizeof(int));
    if (mask == NULL) {
        return NULL;
    }
    a = calloc(n, sizeof(int));
    if (a == NULL) {
        free(mask);
        return NULL;
    }

    for (i = 0; i < n; i++) {
        for (;;) {
            j = random() % ARRAY_MAX_ELEMENT;
            if (mask[j] == 0) {
                a[i] = j;
                mask[j] = 1;
                break;
            }
        }
    }

    free(mask);
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
    a = generate_no_dup_array(n);
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
    int result, answer, *a, n, target, rc;

    a = generate_test_array(&n);
    if (a == NULL) {
        return -1;
    }

    target = generate_random_num(ARRAY_MAX_ELEMENT);

    rc = 0;
    result = searchInsert(a, n, target);
    answer = search_insert_bf(a, 0, n-1, target);
    if (result != answer) {
        printf("Search %d in array: ", target);
        dump_array(a, n);
        printf("Get %d, should be %d\n", result, answer);
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
