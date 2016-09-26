#include <stdio.h>
#include <stdlib.h>
#include "list.h"

list_node_t *
make_list (int *a, int n)
{
    int i;
    list_node_t *p;

    p = calloc(n, sizeof(list_node_t));
    if (p == NULL) {
        return NULL;
    }

    for (i = 0; i < n; i++) {
        p[i].val = a[i];
        if (i < n-1) {
            p[i].next = &p[i+1];
        }
    }

    return p;
}

void
free_list (list_node_t *p)
{
    free(p);
    return;
}

void
dump_list (list_node_t *head)
{
    list_node_t *p;

    if (head == NULL) {
        printf("<NULL>\n");
        return;
    }

    for (p = head; p; p = p->next) {
        if (p != head) {
            printf("->");
        }
        printf("%d", p->val);
    }
    printf("\n");
    return;
}

