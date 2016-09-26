#include <stddef.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

static list_node_t *
reverse_and_count (list_node_t *head, int *n)
{
    list_node_t *prev, *cur, *next;

    *n = 0;
    prev = NULL;
    cur = head;

    while (cur) {
        next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
        (*n)++;
    }
    head = prev;
    return head;
}

static list_node_t *
reverse_first_kth (list_node_t *head, int k)
{
    list_node_t *prev, *cur, *next;

    prev = NULL;
    cur = head;
    for (; cur != NULL && k > 0; k--) {
        next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }

    head->next = cur;

    head = prev;
    return head;
}

static list_node_t *
rotate_right_internal (list_node_t *head, int k)
{
    int n;
    list_node_t *last, *head1, *head2;

    head = reverse_and_count(head, &n);

    k %= n;
    if (k == 0) {
        head = reverse_and_count(head, &n);
        return head;
    }

    last = head;
    head1 = reverse_first_kth(head, k);

    head2 = reverse_first_kth(last->next, n-k);
    last->next = head2;
    return head1;
}

list_node_t *
rotateRight (list_node_t *head, int k)
{
    if (head == NULL || k <= 0) {
        return head;
    }

    head = rotate_right_internal(head, k);
    return head;
}

#include <stdio.h>
#include <stdlib.h>

static void
dump_list (list_node_t *p)
{
    while (p) {
        printf("%d ", p->val);
        p = p->next;
    }
    printf("\n");
    return;
}

static void
clean_list (list_node_t *head)
{
    list_node_t *p, *next;

    p = head;
    while (p) {
        next = p->next;
        free(p);
        p = next;
    }
    return;
}

static list_node_t *
make_list (int *a, int n)
{
    list_node_t *head, **pprev, *p;
    int i;

    head = NULL;
    pprev = &head;

    for (i = 0; i < n; i++) {
        p = calloc(1, sizeof(list_node_t));
        if (p == NULL) {
            clean_list(head);
            return NULL;
        }
        p->val = a[i];
        *pprev = p;
        pprev = &p->next;
    }

    return head;
}

static void
test_shift (int *a, int n, int k)
{
    list_node_t *p;

    p = make_list(a, n);
    if (p == NULL) {
        return;
    }

    p = rotateRight(p, k);

    printf("Shift for %d\n", k);
    dump_list(p);
    printf("----------------\n");

    clean_list(p);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3, 4, 5};

    test_shift(a0, ARRAY_LEN(a0), 0);
    test_shift(a0, ARRAY_LEN(a0), 1);
    test_shift(a0, ARRAY_LEN(a0), 2);
    test_shift(a0, ARRAY_LEN(a0), 3);
    test_shift(a0, ARRAY_LEN(a0), 4);
    test_shift(a0, ARRAY_LEN(a0), 5);
    test_shift(a0, ARRAY_LEN(a0), 6);
    return 0;
}

