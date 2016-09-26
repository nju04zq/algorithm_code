#include <stddef.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

list_node_t *
swapPairs (list_node_t *head)
{
    list_node_t *p1, *p2, **pprev, *next;

    pprev = &head;
    p1 = head;
    while (p1) {
        p2 = p1->next;
        if (p2 == NULL) {
            *pprev = p1;
            break;
        }

        next = p2->next;
        *pprev = p2;
        p2->next = p1;
        p1->next = NULL;

        pprev = &p1->next;
        p1 = next;
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
test_swap_pairs (int *a, int n)
{
    list_node_t *head;

    head = make_list(a, n);
    if (head == NULL) {
        return;
    }

    printf("Before SWAP: ");
    dump_list(head);
    head = swapPairs(head);
    printf("After SWAP: ");
    dump_list(head);

    clean_list(head);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3, 4, 5};
    int a1[] = {1, 2, 3, 4, 5, 6};

    test_swap_pairs(a0, ARRAY_LEN(a0));
    test_swap_pairs(a1, ARRAY_LEN(a1));
    return 0;
}
