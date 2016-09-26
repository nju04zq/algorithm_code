/*
 * a -> ... -> an -> c -> ... -> cn
 * b -> ... -> bn -> c -> ... -> cn
 * |a->an| = x1, |b->bn| = x2, |c->cn| = x3
 * x1 + x3 = m, x2 + x3 = n
 ************************************************************
 * If there is intersection, total running time 3(x1+x2+x3)
 * x1+x3 for reverse a and count
 * x2+x3 for get b tail and count
 * x1+x2 for count new b
 * x2    for get intersction
 * x1+x3 for reverse new a
 ************************************************************
 * If there is no intersection, total running time 2m+n
 * m for reverse a and count
 * n for get b tail and count
 * m for reverse new a
 */
#include <stddef.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

static list_node_t *
reverse_and_count_list (list_node_t *head, int *size)
{
    list_node_t *prev, *cur, *next;

    prev = NULL;
    cur = head;
    *size = 0;
    while (cur) {
        next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
        (*size)++;
    }

    head = prev;
    return head;
}

static list_node_t *
get_list_tail_and_count (list_node_t *head, int *size)
{
    list_node_t *prev, *cur, *tail;

    *size = 0;
    prev = NULL;
    cur = head;
    while (cur) {
        prev = cur;
        cur = cur->next;
        (*size)++;
    }

    tail = prev;
    return tail;
}

static int
get_list_size (list_node_t *head)
{
    list_node_t *p;
    int size = 0;

    for (p = head; p; p = p->next) {
        size++;
    }
    return size;
}

static list_node_t *
get_list_entry_by_index (list_node_t *head, int index)
{
    int i;
    list_node_t *p;

    for (p = head, i = 0; p; p = p->next, i++) {
        if (i == index) {
            return p;
        }
    }
    return NULL;
}

list_node_t *
getIntersectionNode (list_node_t *head_a, list_node_t *head_b)
{
    list_node_t *tail_a, *tail_b, *p;
    int size_a, size_b, size_b_new, index;

    if (head_a == NULL || head_b == NULL) {
        return NULL;
    }

    if (head_a == head_b) {
        return head_a;
    }

    tail_b = get_list_tail_and_count(head_b, &size_b);
    tail_a = reverse_and_count_list(head_a, &size_a);

    if (tail_a != tail_b) {
        (void)reverse_and_count_list(tail_a, &size_a);
        return NULL;
    }

    size_b_new = get_list_size(head_b) - 1; //do not count in the intersection
    index = (size_a + size_b + size_b_new)/2 - size_a;

    p = get_list_entry_by_index(head_b, index);

    (void)reverse_and_count_list(tail_a, &size_a);
    return p;
}

#include <stdio.h>

#define DEF_NODE(x) list_node_t n##x = {x, NULL};

static void
dump_list (int index, list_node_t *head)
{
    int i;
    list_node_t *p;

    printf("List #%d: ", index);
    for (p = head, i = 0; p; p = p->next, i++) {
        if (i == 0) {
            printf("%d", p->val);
        } else {
            printf(" -> %d", p->val);
        }
    }
    printf("\n");
    return;
}

static void
test_get_intersection(list_node_t *head_a, list_node_t *head_b)
{
    list_node_t *p;

    p = getIntersectionNode(head_a, head_b);
    dump_list(0, head_a);
    dump_list(1, head_b);
    printf("Intersection %d\n", p ? p->val : -1);
    return;
}

static void
test_case_1 (void)
{
    DEF_NODE(0);
    DEF_NODE(1);
    DEF_NODE(2);
    DEF_NODE(3);
    DEF_NODE(4);
    DEF_NODE(5);
    DEF_NODE(6);
    DEF_NODE(7);

    n0.next = &n1;
    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    n4.next = &n5;
    n5.next = NULL;
    n6.next = &n7;
    n7.next = &n3;
    test_get_intersection(&n0, &n6);
    return;
}

static void
test_case_2 (void)
{
    DEF_NODE(0);
    DEF_NODE(1);
    DEF_NODE(2);
    DEF_NODE(3);
    DEF_NODE(4);
    DEF_NODE(5);
    DEF_NODE(6);
    DEF_NODE(7);

    n0.next = &n1;
    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    n4.next = &n5;
    n5.next = NULL;
    n6.next = &n7;
    n7.next = NULL;
    test_get_intersection(&n0, &n6);
    return;
}

static void
test_case_3 (void)
{
    DEF_NODE(0);
    DEF_NODE(1);
    DEF_NODE(2);
    DEF_NODE(3);
    DEF_NODE(4);
    DEF_NODE(5);
    DEF_NODE(6);
    DEF_NODE(7);

    n0.next = &n1;
    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    n4.next = &n5;
    n5.next = &n6;
    n6.next = &n7;
    test_get_intersection(&n0, &n2);
    return;
}

int main (void)
{
    test_case_1();
    test_case_2();
    test_case_3();
    return 0;
}

