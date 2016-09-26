#include <stddef.h>
#include <stdlib.h>
#include "../../common/bool.h"

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

typedef struct min_heap_s {
    list_node_t **buf;
    int size;
    int max_size;
} min_heap_t;

#define MIN_HEAP_VAL(i) (min_heap->buf[i]->val)

#define MIN_HEAP_SWAP(i, j) \
do {\
    list_node_t *__tmp = min_heap->buf[i];\
    min_heap->buf[i] = min_heap->buf[j];\
    min_heap->buf[j] = __tmp;\
} while (0)

#define MIN_HEAP_LAST() (min_heap->size-1)

static int
calc_parent (int i)
{
    return ((i+1)/2 -1);
}

static int
calc_lchild (int i)
{
    return (2*(i+1)-1);
}

static int
calc_rchild (int i)
{
    return (2*(i+1));
}

static void
min_heap_destory (min_heap_t *min_heap)
{
    free(min_heap->buf);
    return;
}

static int
min_heap_create (min_heap_t *min_heap, int max_size)
{
    list_node_t **buf;

    buf = calloc(max_size, sizeof(list_node_t *));
    if (buf == NULL) {
        return -1;
    }

    min_heap->buf = buf;
    min_heap->size = 0;
    min_heap->max_size = max_size;
    return 0;
}

static bool
min_heap_is_empty (min_heap_t *min_heap)
{
    if (min_heap->size == 0) {
        return TRUE;
    } else {
        return FALSE;
    }
}

static void
min_heap_slide_down (min_heap_t *min_heap, int i)
{
    int lchild, rchild, j;

    for (;;) {
        lchild = calc_lchild(i);
        rchild = calc_rchild(i);
        if (lchild > MIN_HEAP_LAST()) {
            break;
        }
        if (rchild > MIN_HEAP_LAST()) {
            if (MIN_HEAP_VAL(lchild) < MIN_HEAP_VAL(i)) {
                MIN_HEAP_SWAP(lchild, i);
            }
            break;
        }
        if (MIN_HEAP_VAL(lchild) < MIN_HEAP_VAL(rchild)) {
            j = lchild;
        } else {
            j = rchild;
        }
        if (MIN_HEAP_VAL(i) <= MIN_HEAP_VAL(j)) {
            break;
        }
        MIN_HEAP_SWAP(i, j);
        i = j;
    }
    return;
}

static void
min_heap_slide_up (min_heap_t *min_heap, int i)
{
    int parent;

    while (i > 0) {
        parent = calc_parent(i);
        if (MIN_HEAP_VAL(i) < MIN_HEAP_VAL(parent)) {
            MIN_HEAP_SWAP(i, parent);
        }
        i = parent;
    }
    return;
}

static void
min_heap_init (min_heap_t *min_heap, list_node_t **s, int k)
{
    int i, j;

    for (i = 0, j = 0; i < k; i++) {
        if (s[i] == NULL) {
            continue;
        }
        min_heap->buf[j++] = s[i];
        min_heap->size++;
    }

    for (i = 0; i < min_heap->size; i++) {
        min_heap_slide_up(min_heap, i);
    }
    return;
}

static list_node_t *
min_heap_get_min (min_heap_t *min_heap)
{
    if (min_heap_is_empty(min_heap)) {
        return NULL;
    }
    return min_heap->buf[0];
}

static void
min_heap_reload (min_heap_t *min_heap, list_node_t *p)
{
    list_node_t *last;
    int last_idx;

    last_idx = MIN_HEAP_LAST();
    last = min_heap->buf[last_idx];

    if (p == NULL) {
        p = last;
        min_heap->size--;
    }

    if (min_heap_is_empty(min_heap)) {
        return;
    }

    min_heap->buf[0] = p;
    min_heap_slide_down(min_heap, 0);
    return;
}

static list_node_t *
merge_k_lists_internal (min_heap_t *min_heap, list_node_t **s, int k)
{
    list_node_t *head, **pprev, *p;

    min_heap_init(min_heap, s, k);

    head = NULL;
    pprev = &head;
    for (;;) {
        p = min_heap_get_min(min_heap);
        if (p == NULL) {
            break;
        }
        min_heap_reload(min_heap, p->next);

        p->next = NULL;
        *pprev = p;
        pprev = &p->next;
    }

    return head;
}

list_node_t *
mergeKLists (list_node_t **lists, int k)
{
    min_heap_t min_heap;
    list_node_t *head;
    int rc;

    if (lists == NULL || k <= 0) {
        return NULL;
    }

    rc = min_heap_create(&min_heap, k);
    if (rc != 0) {
        return NULL;
    }

    head = merge_k_lists_internal(&min_heap, lists, k);

    min_heap_destory(&min_heap);
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

