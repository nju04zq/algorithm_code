#include <stddef.h>
#include <stdlib.h>
#include "../../common/queue.h"

static int
quicksort_split (int *a, int size)
{
    int i, j, pilot;

    pilot = a[0];
    for (i = 1, j = 0; i < size; i++) {
        if (a[i] <= pilot) {
            a[j++] = a[i];
            a[i] = a[j];
        }
    }

    a[j] = pilot;
    return j;
}

static void
quicksort (int *a, int size)
{
    int m;

    if (size <= 1) {
        return;
    }

    m = quicksort_split(a, size);
    quicksort(&a[0], m);
    quicksort(&a[m+1], size-m-1);
    return;
}

static int
add_sum_to_queue (uqueue_t *queue, int a, int b, int c, int d)
{
    int rc, *nums;

    nums = calloc(4, sizeof(int));
    if (nums == NULL) {
        return -1;
    }

    nums[0] = a;
    nums[1] = b;
    nums[2] = c;
    nums[3] = d;

    rc = enqueue(queue, nums);
    if (rc != 0) {
        free(nums);
        return -1;
    }

    return 0;
}

static int
do_2sum_internal (uqueue_t *queue, int *a, int size, int target, int a0, int a1)
{
    int i, j, total, rc;

    i = 0;
    j = size-1;
    while (i < j) {
        if (i > 0 && a[i] == a[i-1]) {
            i++;
            continue;
        }
        if (j < (size-1) && a[j] == a[j+1]) {
            j--;
            continue;
        }
        total = a[i] + a[j];
        if (total == target) {
            rc = add_sum_to_queue(queue, a0, a1, a[i], a[j]);
            if (rc != 0) {
                return rc;
            }
            i++;
            j--;
        } else if (total < target) {
            i++;
        } else {
            j--;
        }
    }

    return 0;
}

static int
do_3sum_internal (uqueue_t *queue, int *a, int size, int target, int a0)
{
    int i, rc;

    for (i = 0; i < size-2; i++) {
        if (i > 0 && a[i] == a[i-1]) {
            continue;
        }
        rc = do_2sum_internal(queue, &a[i+1], size-i-1, target-a[i], a0, a[i]);
        if (rc != 0) {
            return rc;
        }
    }

    return 0;
}

static int
do_4sum_internal (uqueue_t *queue, int *a, int size, int target)
{
    int i, rc;

    quicksort(a, size);

    for (i = 0; i < size-3; i++) {
        if (i > 0 && a[i] == a[i-1]) {
            continue;
        }
        rc = do_3sum_internal(queue, &a[i+1], size-i-1, target-a[i], a[i]);
        if (rc != 0) {
            return rc;
        }
    }

    return 0;
}

static int **
copy_out_queue (uqueue_t *queue, int *returnSize)
{
    int **result, i, size;

    size = get_queue_size(queue);
    if (size == 0) {
        return NULL;
    }

    result = calloc(size, sizeof(int *));
    if (result == NULL) {
        return NULL;
    }

    *returnSize = size;
    for (i = 0; size > 0; size--, i++) {
        result[i] = dequeue(queue);
    }

    return result;
}

int **
fourSum (int *a, int size, int target, int *returnSize)
{
    uqueue_t queue;
    int rc, **result = NULL;

    *returnSize = 0;
    if (a == NULL) {
        return NULL;
    }

    rc = init_queue(&queue);
    if (rc != 0) {
        return NULL;
    }

    rc = do_4sum_internal(&queue, a, size, target);
    if (rc == 0) {
        result = copy_out_queue(&queue, returnSize);
    }

    clean_queue(&queue);
    return result;
}

#include <stdio.h>

static void
dump_array (int *a, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
    return;
}

static void
dump_result (int **result, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        dump_array(result[i], 4);
    }
    return;
}

static void
free_result (int **result, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        free(result[i]);
    }
    free(result);
    return;
}

static void
test_4sum (int *a, int size, int target)
{
    int **result, result_size;

    printf("4sum for ");
    dump_array(a, size);

    result = fourSum(a, size, target, &result_size);
    if (result) {
        dump_result(result, result_size);
        free_result(result, result_size);
    } else {
        printf("<NULL>\n");
    }
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {0, 0, 0, 0, 0, 0, 0};
    int a1[] = {-2, -1, 0, 0, 1, 2};

    test_4sum(a0, ARRAY_LEN(a0), 0);
    test_4sum(a1, ARRAY_LEN(a1), 0);
    return 0;
}

