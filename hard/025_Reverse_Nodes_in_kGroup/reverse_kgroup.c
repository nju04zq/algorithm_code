#include <stddef.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

static list_node_t *
get_k_nodes (list_node_t *p, int k, int *cnt)
{
    int i;
    list_node_t *prev;

    *cnt = 0;
    prev = NULL;
    for (i = 1;;i++) {
        if (i == k) {
            break;
        }
        prev = p;
        p = p->next;
        if (p == NULL) {
            break;
        }
    }

    if (p) {
        *cnt = k;
        return p;
    } else {
        *cnt = i;
        return prev;
    }
}

static void
reverse_list (list_node_t *head)
{
    list_node_t *cur, *prev, *next;

    prev = NULL;
    cur = head;
    while (cur) {
        next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }
    return;
}

list_node_t *
reverseKGroup (list_node_t *head, int k)
{
    int cnt;
    list_node_t *p, *pfirst, *plast, **pprev, *next;

    if (head == NULL || k <= 1) {
        return head;
    }

    p = head;
    pprev = &head;
    pfirst = head;
    while (pfirst) {
        plast = get_k_nodes(pfirst, k, &cnt);
        if (cnt < k) {
            *pprev = pfirst;
            break;
        }

        next = plast->next;
        plast->next = NULL;
        reverse_list(pfirst);
        *pprev = plast;

        pprev = &pfirst->next;
        pfirst = next;
    }

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
test_swap_pairs (int *a, int n, int k)
{
    list_node_t *head;

    head = make_list(a, n);
    if (head == NULL) {
        return;
    }

    printf("Before reverse: ");
    dump_list(head);
    head = reverseKGroup(head, k);
    printf("After reverse: ");
    dump_list(head);

    clean_list(head);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3, 4, 5};
    int a1[] = {1, 2, 3, 4, 5, 6};

    test_swap_pairs(a0, ARRAY_LEN(a0), 2);
    test_swap_pairs(a0, ARRAY_LEN(a0), 3);
    test_swap_pairs(a1, ARRAY_LEN(a1), 6);
    test_swap_pairs(a1, ARRAY_LEN(a1), 7);
    return 0;
}

