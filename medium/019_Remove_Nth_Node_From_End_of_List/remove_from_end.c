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

static list_node_t *
remove_node (list_node_t *head, list_node_t *p, list_node_t *prev)
{
    if (prev == NULL) {
        head = p->next;
        p->next = NULL;
    } else {
        prev->next = p->next;
        p->next = NULL;
    }
    return head;
}

list_node_t *
removeNthFromEnd (list_node_t *head, int n)
{
    list_node_t *p, *nth_prev = NULL, *nth = NULL;
    int cnt = 0;

    if (head == NULL || n <= 0) {
        return head;
    }

    p = head;
    nth = head;
    while (p) {
        if (cnt < n) {
            cnt++;
        } else {
            nth_prev = nth;
            nth = nth->next;
        }
        p = p->next;
    }

    if (cnt < n) {
        return head;
    }

    head = remove_node(head, nth, nth_prev);
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
test_remove_from_end (list_node_t *p, int n)
{
    printf("Before remove %d, ", n);
    dump_list(p);

    p = removeNthFromEnd(p, n);

    printf("After remove %d, ", n);
    dump_list(p);
    return;
}

#define DEF_NODE(x) list_node_t n##x = {x, NULL};

#define CACSCADE_NODES() \
do {\
    n1.next = &n2;\
    n2.next = &n3;\
    n3.next = &n4;\
} while (0)

int main (void)
{
    DEF_NODE(1);
    DEF_NODE(2);
    DEF_NODE(3);
    DEF_NODE(4);

    test_remove_from_end(&n1, 1);

    CACSCADE_NODES();
    test_remove_from_end(&n1, 0);

    CACSCADE_NODES();
    test_remove_from_end(&n1, 1);

    CACSCADE_NODES();
    test_remove_from_end(&n1, 2);

    CACSCADE_NODES();
    test_remove_from_end(&n1, 3);

    CACSCADE_NODES();
    test_remove_from_end(&n1, 4);

    CACSCADE_NODES();
    test_remove_from_end(&n1, 5);
    return 0;
}

