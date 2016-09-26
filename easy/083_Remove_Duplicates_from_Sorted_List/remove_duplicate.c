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

static void
remove_node (list_node_t *p, list_node_t *prev)
{
    prev->next = p->next;
    p->next = NULL;
    return;
}

list_node_t *
deleteDuplicates (list_node_t *head)
{
    list_node_t *cur, *prev, *next;

    if (head == NULL) {
        return NULL;
    }

    prev = head;
    cur = head->next;
    while (cur) {
        next = cur->next;
        if (cur->val == prev->val) {
            remove_node(cur, prev);
        } else {
            prev = cur;
        }
        cur = next;
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
test_remove (list_node_t *head)
{
    printf("Before remove: ");
    dump_list(head);

    head = deleteDuplicates(head);

    printf("After remove: ");
    dump_list(head);
    return;
}

#define DEF_NODE(x) list_node_t n##x = {x, NULL};

int main (void)
{
    DEF_NODE(1);
    DEF_NODE(2);
    DEF_NODE(3);
    DEF_NODE(4);

    DEF_NODE(11);
    DEF_NODE(41);

    test_remove(&n1);

    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    test_remove(&n1);

    n11.val = 1;
    n41.val = 4;
    n1.next = &n11;
    n11.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    n4.next = &n41;
    test_remove(&n1);
    return 0;
}
