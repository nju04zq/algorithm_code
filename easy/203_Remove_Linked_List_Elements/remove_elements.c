struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

list_node_t *
removeElements (list_node_t *head, int val)
{
    list_node_t *p, **pprev;

    pprev = &head;
    for (p = head; p; p = p->next) {
        if (p->val != val) {
            pprev = &p->next;
        } else {
            *pprev = p->next;
        }
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
test_remove (list_node_t *head, int val)
{
    printf("Before remove %d, ", val);
    dump_list(head);

    head = removeElements(head, val);

    printf("After remove %d, ", val);
    dump_list(head);
    return;
}

#define DEF_NODE(x, val) list_node_t n##x = {val, NULL};

static void
test_case_1 (void)
{
    DEF_NODE(0, 2);
    DEF_NODE(1, 1);
    DEF_NODE(2, 2);
    DEF_NODE(3, 3);
    DEF_NODE(4, 2);

    n0.next = &n1;
    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;

    test_remove(&n0, 2);
    return;
}

static void
test_case_2 (void)
{
    DEF_NODE(0, 2);
    DEF_NODE(1, 1);
    DEF_NODE(2, 2);
    DEF_NODE(3, 3);
    DEF_NODE(4, 2);

    n0.next = &n1;
    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;

    test_remove(&n0, 5);
    return;
}

int main (void)
{
    test_case_1();
    test_case_2();
    return 0;
}

