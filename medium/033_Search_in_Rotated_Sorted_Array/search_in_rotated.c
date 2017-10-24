#include <stddef.h>

static int
search_bf (int *a, int start, int end, int target)
{
    int i;

    for (i = start; i <= end; i++) {
        if (target == a[i]) {
            return i;
        }
    }
    return -1;
}

static int
binary_search (int *a, int start, int end, int target)
{
    int mid;

    while (start <= end) {
        mid = (start+end)/2;
        if (target == a[mid]) {
            return mid;
        } else if (target > a[mid]) {
            start = mid+1;
        } else {
            end = mid-1;
        }
    }
    return -1;
}

static int
search_internal (int *a, int start, int end, int target)
{
    int mid, n, result;

    n = end - start + 1;
    if (n <= 2) {
        result = search_bf(a, start, end, target);
        return result;
    }

    mid = (start + end)/2;
    if (target == a[mid]) {
        return mid;
    }
    if (a[start] < a[mid] && a[mid] < a[end]) {
        result = binary_search(a, start, end, target);
    }
    if (a[start] > a[mid] && a[mid] < a[end]) {
        if (target > a[mid] && target <= a[end]) {
            result = binary_search(a, mid+1, end, target);
        } else {
            result = search_internal(a, start, mid-1, target);
        }
    }
    if (a[start] < a[mid] && a[mid] > a[end]) {
        if (target < a[mid] && target >= a[start]) {
            result = binary_search(a, start, mid-1, target);
        } else {
            result = search_internal(a, mid+1, end, target);
        }
    }
    return result;
}

int
search (int *a, int n, int target)
{
    int result;

    if (a == NULL || n <= 0) {
        return -1;
    }
    
    result = search_internal(a, 0, n-1, target);
    return result;
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

#define SWAP(a, i, j) \
do {\
    int __tmp = a[i];\
    a[i] = a[j];\
    a[j] = __tmp;\
} while (0)

static void
reverse_array (int *a, int n)
{
    int i, j;

    for (i = 0, j = n-1; i < j; i++, j--) {
        SWAP(a, i, j);
    }
    return;
}

static void
shift_array (int *a, int n, int k)
{
    reverse_array(a, n);
    reverse_array(&a[0], k);
    reverse_array(&a[k], n-k);
    return;
}

static int *
generate_test_array (int *len)
{
    int n, k, *a;

    n = generate_random_num(ARRAY_LEN_RANGE);
    n += ARRAY_LEN_BASE;
    a = generate_no_dup_array(n);
    if (a == NULL) {
        return NULL;
    }

    k = generate_random_num(n);
    quicksort(a, n);
    shift_array(a, n, k);

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
    result = search(a, n, target);
    answer = search_bf(a, 0, n-1, target);
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

