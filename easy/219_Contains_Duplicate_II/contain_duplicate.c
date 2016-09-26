#include <stdlib.h>
#include "../../common/bool.h"

#define HASH_SET_SIZE 97

typedef struct hash_entry_s {
    struct hash_entry_s *h_next;
    int data;
    int index;
} hash_entry_t;

typedef struct hash_set_s {
    hash_entry_t **buf;
} hash_set_t;

static int
hash_set_init (hash_set_t *hash_set_p)
{
    hash_entry_t **buf;

    buf = calloc(HASH_SET_SIZE, sizeof(hash_entry_t));
    if (buf == NULL) {
        return -1;
    }

    hash_set_p->buf = buf;
    return 0;
}

static void
clean_one_hash_set_slot (hash_entry_t *head_p)
{
    hash_entry_t *entry_p, *next_p;

    entry_p = head_p;
    while (entry_p) {
        next_p = entry_p->h_next;
        free(entry_p);
        entry_p = next_p;
    }
    return;
}

static void
hash_set_clean (hash_set_t *hash_set_p)
{
    int i;

    for (i = 0; i < HASH_SET_SIZE; i++) {
        clean_one_hash_set_slot(hash_set_p->buf[i]);
    }

    free(hash_set_p->buf);
    return;
}

static int
hash_set_fn (int n)
{
    n = n >= 0 ? n : -n;
    return (n % HASH_SET_SIZE);
}

static int
hash_set_insert (hash_set_t *hash_set_p, int data, int index)
{
    int i;
    hash_entry_t *entry_p;

    entry_p = calloc(1, sizeof(hash_entry_t));
    if (entry_p == NULL) {
        return -1;
    }

    entry_p->data = data;
    entry_p->index = index;

    i = hash_set_fn(data);
    entry_p->h_next = hash_set_p->buf[i];
    hash_set_p->buf[i] = entry_p;
    return 0;
}

static hash_entry_t *
hash_set_get (hash_set_t *hash_set_p, int data)
{
    int i;
    hash_entry_t *entry_p;

    i = hash_set_fn(data);

    entry_p = hash_set_p->buf[i];
    while (entry_p) {
        if (entry_p->data == data) {
            return entry_p;
        }
        entry_p = entry_p->h_next;
    }

    return NULL;
}

static bool
contain_nearby_duplicate_internal (hash_set_t *hash_set_p,
                                   int *nums, int size, int k)
{
    int i, rc;
    hash_entry_t *entry_p;

    for (i = 0; i < size; i++) {
        entry_p = hash_set_get(hash_set_p, nums[i]);
        if (entry_p) {
            if ((i - entry_p->index) <= k) {
                return TRUE;
            } else {
                entry_p->index = i;
                continue;
            }
        }

        rc = hash_set_insert(hash_set_p, nums[i], i);
        if (rc != 0) {
            return FALSE;
        }
    }

    return FALSE;
}

bool
containsNearbyDuplicate (int *nums, int size, int k)
{
    int rc;
    bool result;
    hash_set_t hash_set;

    if (nums == NULL || size == 0 || k == 0) {
        return FALSE;
    } 

    rc = hash_set_init(&hash_set);
    if (rc != 0) {
        return FALSE;
    }

    result = contain_nearby_duplicate_internal(&hash_set, nums, size, k);

    hash_set_clean(&hash_set);
    return result;
}

#include <stdio.h>

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

static void
test_contain_neaby_dup (int *nums, int size, int k, bool answer)
{
    int i;
    bool result;

    result = containsNearbyDuplicate(nums, size, k);
    if (result == answer) {
        return;
    }

    for (i = 0; i < size; i++) {
        printf("%d ", nums[i]);
    }
    printf(", contain dups nearby <= %d, get %d, should be %d\n",
           k, result, answer);
    return;
}

int main (void)
{
    int a0[] = {1, 2, 3, 1, 4};
    int a1[] = {1, 2, 3, 4};
    int a2[] = {-1, -1};
    int a3[] = {2, 2};

    test_contain_neaby_dup(a0, ARRAY_LEN(a0), 3, TRUE);
    test_contain_neaby_dup(a0, ARRAY_LEN(a0), 2, FALSE);
    test_contain_neaby_dup(a1, ARRAY_LEN(a1), 1, FALSE);
    test_contain_neaby_dup(a2, ARRAY_LEN(a2), 1, TRUE);
    test_contain_neaby_dup(a3, ARRAY_LEN(a3), 3, TRUE);
    return 0;
}

