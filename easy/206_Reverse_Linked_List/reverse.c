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
reverseList (list_node_t *head)
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

    return prev;
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
test_reverse (list_node_t *head)
{
    printf("Before reverse: ");
    dump_list(head);

    head = reverseList(head);

    printf("After reverse: ");
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

    test_reverse(&n1);

    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    test_reverse(&n1);
    return 0;
}

