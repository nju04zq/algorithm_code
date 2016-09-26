#include <stddef.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};
typedef struct ListNode list_node_t;

static list_node_t *
alloc_node (int val)
{
    list_node_t *p;

    p = calloc(1, sizeof(list_node_t));
    if (p == NULL) {
        return NULL;
    }

    p->val = val;
    return p;
}

static list_node_t *
add_two (int a, int b, int *carry)
{
    list_node_t *p3;
    int c;

    c = a + b + *carry;
    if (c >= 10) {
        c -= 10;
        *carry = 1;
    } else {
        *carry = 0;
    }
    
    p3 = alloc_node(c);
    return p3;
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

list_node_t *
addTwoNumbers (list_node_t *p1, list_node_t *p2)
{
    list_node_t *head = NULL, *p3, **pprev;
    int a, b, carry = 0;

    pprev = &head;
    while (p1 || p2) {
        a = 0;
        b = 0;
        if (p1 != NULL) {
            a = p1->val;
            p1 = p1->next;
        }
        if (p2 != NULL) {
            b = p2->val;
            p2 = p2->next;
        }

        p3 = add_two(a, b, &carry);
        if (p3 == NULL) {
            clean_list(head);
            return NULL;
        }
        *pprev = p3;
        pprev = &p3->next;
    }

    if (carry == 1) {
        p3 = alloc_node(1);
        *pprev = p3;
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
test_add_two (list_node_t *s1, list_node_t *s2)
{
    list_node_t *s3;

    s3 = addTwoNumbers(s1, s2);

    printf("Number1: ");
    dump_list(s1);
    printf("Number2: ");
    dump_list(s2);
    printf("Result: ");
    dump_list(s3);

    clean_list(s3);
    return;
}

#define DEF_NODE(x, val) list_node_t n##x = {val, NULL};

int main (void)
{
    DEF_NODE(0, 9);
    DEF_NODE(1, 9);
    DEF_NODE(2, 9);
    DEF_NODE(3, 1);

    n0.next = &n1;
    n1.next = &n2;
    test_add_two(&n0, &n3);
    return 0;
}

