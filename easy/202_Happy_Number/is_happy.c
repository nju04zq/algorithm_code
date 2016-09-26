#include <stdlib.h>
#include "../../common/bool.h"

#define HASH_SET_SIZE 97

typedef struct hash_entry_s {
    struct hash_entry_s *h_next;
    int data;
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
    return (n % HASH_SET_SIZE);
}

static int
hash_set_insert (hash_set_t *hash_set_p, int data)
{
    int i;
    hash_entry_t *entry_p;

    entry_p = calloc(1, sizeof(hash_entry_t));
    if (entry_p == NULL) {
        return -1;
    }

    entry_p->data = data;

    i = hash_set_fn(data);
    entry_p->h_next = hash_set_p->buf[i];
    hash_set_p->buf[i] = entry_p;
    return 0;
}

static bool
hash_set_has (hash_set_t *hash_set_p, int data)
{
    int i;
    hash_entry_t *entry_p;

    i = hash_set_fn(data);

    entry_p = hash_set_p->buf[i];
    while (entry_p) {
        if (entry_p->data == data) {
            return TRUE;
        }
        entry_p = entry_p->h_next;
    }

    return FALSE;
}

static int
perform_happy (int n)
{
    int sum = 0, x;

    while (n > 0) {
        x = n % 10;
        sum += (x * x);
        n /= 10;
    }
    return sum;
}

static bool
is_happy_internal (hash_set_t *hash_set_p, int n)
{
    int rc;

    for (;;) {
        if (hash_set_has(hash_set_p, n)) {
            break;
        }
        rc = hash_set_insert(hash_set_p, n);
        if (rc != 0) {
            break;
        }

        n = perform_happy(n);
        if (n == 1) {
            return TRUE;
        }
    }

    return FALSE;
}

bool
isHappy (int n)
{
    int rc;
    bool result;
    hash_set_t hash_set;

    if (n <= 0) {
        return FALSE;
    }

    rc = hash_set_init(&hash_set);
    if (rc != 0) {
        return FALSE;
    }

    result = is_happy_internal(&hash_set, n);

    hash_set_clean(&hash_set);
    return result;
}

#include <stdio.h>

static void
test_is_happy (int n, bool answer)
{
    int is_happy;

    is_happy = isHappy(n);
    if (is_happy != answer) {
        printf("%d is happy, get %d, should be %d\n", n, is_happy, answer);
    }
    return;
}

int main (void)
{
    test_is_happy(48, FALSE);
    test_is_happy(23, TRUE);
    test_is_happy(31, TRUE);
    test_is_happy(94, TRUE);
    test_is_happy(97, TRUE);
    test_is_happy(81, FALSE);
    test_is_happy(21, FALSE);
    return 0;
}

