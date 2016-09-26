#include <stddef.h>
#include "../../common/list.h"

list_node_t *
partition (list_node_t *head, int x)
{
    list_node_t **pprev_below, *head_above, **pprev_above, *cur, *next;

    head_above = NULL;
    pprev_below = &head;
    pprev_above = &head_above;

    cur = head;
    while (cur) {
        if (cur->val < x) {
            *pprev_below = cur;
            pprev_below = &cur->next;
        } else {
            *pprev_above = cur;
            pprev_above = &cur->next;
        }
        next= cur->next;
        cur->next = NULL;
        cur = next;
    }

    *pprev_below = head_above;
    return head;
}

#include <stdio.h>
#include <stdlib.h>

static void
test_partition (int *a, int n, int x)
{
    list_node_t *head, *buf;

    buf = make_list(a, n);
    if (buf == NULL) {
        return;
    }
    head = buf;

    printf("List before partition %d\n", x);
    dump_list(head);
    head = partition(head, x);
    printf("List after partition %d\n", x);
    dump_list(head);

    free(buf);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof(a[0]))

int main (void)
{
    int a0[] = {1, 4, 3, 2, 5, 2};

    test_partition(a0, ARRAY_LEN(a0), 0);
    test_partition(a0, ARRAY_LEN(a0), 1);
    test_partition(a0, ARRAY_LEN(a0), 2);
    test_partition(a0, ARRAY_LEN(a0), 3);
    test_partition(a0, ARRAY_LEN(a0), 4);
    test_partition(a0, ARRAY_LEN(a0), 5);
    test_partition(a0, ARRAY_LEN(a0), 6);
    return 0;
}

