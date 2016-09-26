#include <stddef.h>

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode {
    int val;
    struct ListNode *next;
} list_node_t;

list_node_t *
mergeTwoLists (list_node_t *p1, list_node_t *p2)
{
    list_node_t *head = NULL, *p, **pprev;

    pprev = &head;
    while (p1 && p2) {
        if (p1->val <= p2->val) {
            p = p1;
            p1 = p1->next;
        } else {
            p = p2;
            p2 = p2->next;
        }

        *pprev = p;
        pprev = &p->next;
        p->next = NULL;
    }

    if (p1) {
        *pprev = p1;
    } else if (p2) {
        *pprev = p2;
    }

    return head;
}

#include <stdio.h>

static void
dump_list (list_node_t *p)
{
    if (p == NULL) {
        printf("<Empty list>\n");
        return;
    }

    while (p) {
        printf("%d ", p->val);
        p = p->next;
    }

    printf("\n");
    return;
}

static void
test_merge (list_node_t *p1, list_node_t *p2)
{
    list_node_t *p;

    printf("Merge two lists:\n");
    printf("p1: ");
    dump_list(p1);
    printf("p2: ");
    dump_list(p2);

    p = mergeTwoLists(p1, p2);
    printf("p: ");
    dump_list(p);
    return;
}

#define DEF_NODE(x) list_node_t n##x = {x, NULL};

int main (void)
{
    DEF_NODE(1);
    DEF_NODE(2);
    DEF_NODE(3);
    DEF_NODE(4);
    DEF_NODE(5);
    DEF_NODE(6);

    test_merge(NULL, NULL);
    test_merge(NULL, &n1);
    test_merge(&n2, &n1);

    n2.next = &n4;
    n4.next = &n6;
    n1.next = &n3;
    test_merge(&n1, &n2);

    n2.next = &n4;
    n4.next = &n6;
    n1.next = &n3;
    n3.next = &n5;
    test_merge(&n1, &n2);
    return 0;
}
