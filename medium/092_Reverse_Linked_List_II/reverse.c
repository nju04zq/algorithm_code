#include <stddef.h>
#include "../../common/list.h"

static void
reverse_util (list_node_t *start, list_node_t *end)
{
    list_node_t *prev, *cur, *next;

    prev = NULL;
    cur = start;
    for (;;) {
        next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
        if (prev == end) {
            break;
        }
    }
    return;
}

list_node_t *
reverseBetween (list_node_t *head, int m, int n)
{
    int i = 1;
    list_node_t *cur, *next, *start = NULL, **pprev;

    if (head == NULL) {
        return NULL;
    }
    if (m == n) {
        return head;
    }

    pprev = &head;
    for (cur = head; cur; cur = cur->next, i++) {
        if (i == m) {
            start = cur;
        } else if (i < m) {
            pprev = &cur->next;
        } else if (i == n) {
            next = cur->next;
            reverse_util(start, cur);
            *pprev = cur;
            start->next = next;
            break;
        }
    }

    return head;
}

#include <stdio.h>

static void
test_reverse_list (int *a, int len, int m, int n)
{
    list_node_t *head, *p;

    head = make_list(a, len);
    if (head == NULL) {
        printf("Fail to make list.\n");
        return;
    }

    p = head;

    printf("Test reverse %d, %d\n", m, n);
    printf("Before reverse: ");
    dump_list(head);
    head = reverseBetween(head, m, n);
    printf("After reverse: ");
    dump_list(head);
    printf("\n");

    free_list(p);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1, 2, 3, 4, 5, 6, 7};
        
    test_reverse_list(a0, ARRAY_LEN(a0), 1, 7);
    test_reverse_list(a0, ARRAY_LEN(a0), 1, 6);
    test_reverse_list(a0, ARRAY_LEN(a0), 1, 2);
    test_reverse_list(a0, ARRAY_LEN(a0), 2, 7);
    test_reverse_list(a0, ARRAY_LEN(a0), 6, 7);
    test_reverse_list(a0, ARRAY_LEN(a0), 2, 3);
    test_reverse_list(a0, ARRAY_LEN(a0), 5, 5);
    return 0;
}

