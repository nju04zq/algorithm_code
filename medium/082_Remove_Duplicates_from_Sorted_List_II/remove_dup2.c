#include <stddef.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

list_node_t *
deleteDuplicates (list_node_t *head)
{
    list_node_t *p, *start, **pprev;

    if (head == NULL) {
        return NULL;
    }

    pprev = &head;
    start = head;
    for (p = start->next;;p = p->next) {
        if (p == NULL) {
            if (start->next != NULL) {
                *pprev = NULL;
            }
            break;
        }
        if (p->val == start->val) {
            continue;
        }
        if (start->next == p) {
            pprev = &start->next;
            start = p;
        } else {
            *pprev = p;
            start = p;
        }
    }

    return head;
}

#include <stdio.h>
#include <stdlib.h>

static list_node_t *
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

static void
free_list (list_node_t *p)
{
    free(p);
    return;
}

static void
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

static void
test_remove_dup (int *a, int n)
{
    list_node_t *head, *p;

    p = make_list(a, n);
    if (p == NULL) {
        return;
    }
    
    head = p;
    printf("Before remove dup, ");
    dump_list(head);
    head = deleteDuplicates(head);
    printf("After remove dup, ");
    dump_list(head);

    free_list(p);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1};
    int a1[] = {1, 2};
    int a2[] = {1, 2, 3};
    int a3[] = {1, 2, 3, 4, 5};
    int a4[] = {1, 1, 1, 1, 1};
    int a5[] = {1, 2, 3, 3, 4};
    int a6[] = {1, 1, 3, 3, 4};
    int a7[] = {1, 1, 3, 3, 4, 4};
    int a8[] = {1, 2, 3, 3, 3};

    test_remove_dup(a0, ARRAY_LEN(a0));
    test_remove_dup(a1, ARRAY_LEN(a1));
    test_remove_dup(a2, ARRAY_LEN(a2));
    test_remove_dup(a3, ARRAY_LEN(a3));
    test_remove_dup(a4, ARRAY_LEN(a4));
    test_remove_dup(a5, ARRAY_LEN(a5));
    test_remove_dup(a6, ARRAY_LEN(a6));
    test_remove_dup(a7, ARRAY_LEN(a7));
    test_remove_dup(a8, ARRAY_LEN(a8));
    return 0;
}

