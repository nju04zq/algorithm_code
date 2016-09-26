/*
 * a -> ... -> an -> c -> ... -> cn
 * b -> ... -> bn -> c -> ... -> cn
 * |a->an| = x1, |b->bn| = x2, |c->cn| = x3
 * x1 + x3 = m, x2 + x3 = n
 ************************************************************
 * If there is intersection, total running time 2(x1+x2+x3)
 * Total running time 2(x1+x2+x3)
 * x1+x3+x2 for travel from head_a
 * x2+x3+x1 for travel from head_a
 ************************************************************
 * If there is no intersection, total running time 2(m+n)
 * m+n for travel from head_a
 * n+m for travel from head_b
 */
#include <stddef.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

list_node_t *
getIntersectionNode (list_node_t *head_a, list_node_t *head_b)
{
    list_node_t *pa, *pa_prev, *pb, *pb_prev;

    pa = head_a;
    pb = head_b;
    pa_prev = NULL;
    pb_prev = NULL;

    if (head_a == NULL || head_b == NULL) {
        return NULL;
    }

    for (;;) {
        if (pa == NULL && pa_prev == NULL) {
            pa_prev = head_a;
            pa = head_b;
            continue;
        }
        if (pb == NULL && pb_prev == NULL) {
            pb_prev = head_b;
            pb = head_a;
            continue;
        }
        if (pa == NULL || pb == NULL) {
            break;
        }
        if (pa == pb) {
            return pa;
        }
        pa = pa->next;
        pb = pb->next;
    }

    return NULL;
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


