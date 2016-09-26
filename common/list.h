#ifndef __LIST_H__
#define __LIST_H__

struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode list_node_t;

list_node_t *
make_list(int *a, int n);

void
free_list(list_node_t *p);

void
dump_list(list_node_t *head);

#endif //__LIST_H__
