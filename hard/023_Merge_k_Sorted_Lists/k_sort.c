#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

typedef struct winner_tree_node_s {
    bool real;
    int val;
    int idx;
    list_node_t *p;
} winner_tree_node_t;

static int
calc_parent (int i)
{
    return ((i+1)/2 -1);
}

static int
calc_sibling (int i)
{
    return (4 * ((i+1)/2) - i -1);
}

static void
compare_winner_tree_node (winner_tree_node_t *winner_tree,
                          int i, int sibling, int parent)
{
    winner_tree_node_t *winner;

    if (winner_tree[i].real == FALSE) {
        winner = &winner_tree[sibling];
    } else if (winner_tree[sibling].real == FALSE) {
        winner = &winner_tree[i];
    } else if (winner_tree[i].val > winner_tree[sibling].val) {
        winner = &winner_tree[sibling];
    } else {
        winner = &winner_tree[i];
    }

    memcpy(&winner_tree[parent], winner, sizeof(winner_tree_node_t));
    return;
}

static winner_tree_node_t *
compete_on_winner_tree (winner_tree_node_t *winner_tree,
                        list_node_t **s, int k, int i)
{
    int parent, sibling;

    i += (k-1);
    while (i > 0) {
        parent = calc_parent(i);
        sibling = calc_sibling(i);
        compare_winner_tree_node(winner_tree, i, sibling, parent);
        i = parent;
    }
    return &winner_tree[0];
}

static void
add_to_winner_tree (winner_tree_node_t *winner_tree,
                    list_node_t **s, int k, int i)
{
    int offset;
    winner_tree_node_t *node;

    offset = k-1;
    node = &winner_tree[offset+i];

    if (s[i] == NULL) {
        node->real = FALSE;
        node->idx = i;
    } else {
        node->real = TRUE;
        node->idx = i;
        node->val = s[i]->val;
        node->p = s[i];
        s[i] = s[i]->next;
    }
    return;
}

static void
init_winner_tree (winner_tree_node_t *winner_tree, list_node_t **s, int k)
{
    int i, offset;

    offset = k-1;
    winner_tree[offset].real = FALSE;
    for (i = 1; i < k; i++) {
        add_to_winner_tree(winner_tree, s, k, i);
    }

    for (i = 1; i < k; i++) {
        (void)compete_on_winner_tree(winner_tree, s, k, i);
    }

    winner_tree[0].idx = 0;
    winner_tree[0].real = TRUE;
    return;
}

list_node_t *
mergeKLists (list_node_t **lists, int k)
{
    winner_tree_node_t *winner_tree, *winner;
    list_node_t **s, *head = NULL, **pprev, *entry;

    if (lists == NULL || k <= 0) {
        return NULL;
    }

    s = calloc(k, sizeof(list_node_t *));
    if (s == NULL) {
        return NULL;
    }
    memcpy(s, lists, k*sizeof(list_node_t *));

    winner_tree = calloc(2*k, sizeof(winner_tree_node_t));
    if (winner_tree == NULL) {
        free(s);
        return NULL;
    }
    init_winner_tree(winner_tree, s, k);

    pprev = &head;
    winner = &winner_tree[0];
    for (;;) {
        add_to_winner_tree(winner_tree, s, k, winner->idx);
        winner = compete_on_winner_tree(winner_tree, s, k, winner->idx);
        if (winner->real == FALSE) {
            break;
        }

        entry = winner->p;
        *pprev = entry;
        pprev = &entry->next;
    }

    free(winner_tree);
    free(s);
    return head;
}

#include <stdio.h>

static void
test_merge_k_lists (list_node_t **s, int k)
{
    list_node_t *head, *p;

    head = mergeKLists(s, k);
    if (head == NULL) {
        return;
    }

    for (p = head; p; p = p->next) {
        printf("%d ", p->val);
    }
    printf("\n");
    return;
}

#define DEF_NODE(x, y) list_node_t n##x##y = {x*5+y, NULL};

static void
test_case_0 (void)
{
    DEF_NODE(0, 0);
    DEF_NODE(0, 1);
    DEF_NODE(1, 0);
    DEF_NODE(1, 1);
    DEF_NODE(1, 2);
    DEF_NODE(2, 0);
    DEF_NODE(2, 1);
    DEF_NODE(2, 2);
    DEF_NODE(3, 0);
    DEF_NODE(3, 1);
    list_node_t *nodes[4];

    n00.next = &n01;
    n10.next = &n11;
    n11.next = &n12;
    n20.next = &n21;
    n21.next = &n22;
    n30.next = &n31;

    nodes[0] = &n00;
    nodes[1] = &n10;
    nodes[2] = &n20;
    nodes[3] = &n30;

    test_merge_k_lists(nodes, 4);
    return;
}

static void
test_case_1 (void)
{
    DEF_NODE(0, 0);
    DEF_NODE(0, 1);
    DEF_NODE(1, 0);
    DEF_NODE(1, 1);
    DEF_NODE(1, 2);
    DEF_NODE(2, 0);
    DEF_NODE(2, 1);
    DEF_NODE(2, 2);
    DEF_NODE(3, 0);
    DEF_NODE(3, 1);
    DEF_NODE(4, 0);
    DEF_NODE(4, 1);
    DEF_NODE(4, 2);
    DEF_NODE(4, 3);
    list_node_t *nodes[5];

    n00.next = &n01;
    n10.next = &n11;
    n11.next = &n12;
    n20.next = &n21;
    n21.next = &n22;
    n30.next = &n31;
    n40.next = &n41;
    n41.next = &n42;
    n42.next = &n43;

    nodes[0] = &n00;
    nodes[1] = &n10;
    nodes[2] = &n20;
    nodes[3] = &n30;
    nodes[4] = &n40;

    test_merge_k_lists(nodes, 5);
    return;
}

static void
test_case_2 (void)
{
    DEF_NODE(0, 0);
    DEF_NODE(0, 1);
    DEF_NODE(1, 0);
    DEF_NODE(1, 1);
    DEF_NODE(1, 2);
    DEF_NODE(2, 0);
    DEF_NODE(2, 1);
    DEF_NODE(2, 2);
    DEF_NODE(3, 0);
    DEF_NODE(3, 1);
    list_node_t *nodes[5];

    n00.next = &n01;
    n10.next = &n11;
    n11.next = &n12;
    n20.next = &n21;
    n21.next = &n22;
    n30.next = &n31;

    nodes[0] = NULL;
    nodes[1] = &n00;
    nodes[2] = &n10;
    nodes[3] = &n20;
    nodes[4] = &n30;

    test_merge_k_lists(nodes, 5);
    return;
}

static void
test_case_3 (void)
{
    list_node_t *nodes[3];

    nodes[0] = NULL;
    nodes[1] = NULL;
    nodes[2] = NULL;
    test_merge_k_lists(nodes, 3);
    return;
}

int main (void)
{
    test_case_0();
    test_case_1();
    test_case_2();
    test_case_3();
    return 0;
}

