#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include "bool.h"
#include "tree.h"
#include "queue.h"

void
destory_tree (tree_node_t *root)
{
    if (root == NULL) {
        return;
    }
    destory_tree(root->left);
    destory_tree(root->right);
    free(root);
    return;
}

static int
get_token (char *str)
{
    int len = 0;

    while (*str != '\0' && *str != ',') {
        len++;
        str++;
    }
    return len;
}

static char *
goto_next_token (char *str, int len)
{
    str += len;
    if (*str == ',') {
        str++;
    }
    return str;
}

static int
token_to_int (char *token, int len)
{
    int i, val = 0;
    char ch;

    for (i = 0; i < len; i++) {
        ch = token[i];
        if (ch == '#') {
            val = -1;
            break;
        }
        if (ch < '0' || ch > '9') {
            continue;
        }
        val = val * 10 + (ch - '0');
    }
    return val;
}

static tree_node_t *
create_tree_node (char *token, int len)
{
    tree_node_t *p;
    int val;

    val = token_to_int(token, len);
    if (val == -1) { //if token is #, no need to create node
        return NULL;
    }

    p = calloc(1, sizeof(tree_node_t));
    if (p == NULL) {
        return NULL;
    }

    p->val = val;
    return p;
}

static tree_node_t *
create_tree_internal (char *str, uqueue_t *queue)
{
    tree_node_t *root = NULL, *p, *parent;
    bool hook_as_left_child = TRUE;
    int len;

    for (;;) {
        len = get_token(str);
        if (len == 0) {
            break;
        }

        p = create_tree_node(str, len);
        if (p) {
            enqueue(queue, (void *)p);
        }

        if (root == NULL) {
            root = p;
            str = goto_next_token(str, len);
            continue;
        }
        
        parent = peek_queue_head(queue);
        if (!parent) {
            break;
        }
        if (hook_as_left_child) {
            parent->left = p;
            hook_as_left_child = FALSE;
        } else {
            parent->right = p;
            (void)dequeue(queue);
            hook_as_left_child = TRUE;
        }

        str = goto_next_token(str, len);
    }

    return root;
}

tree_node_t *
create_tree (char *str)
{
    uqueue_t queue;
    tree_node_t *root = NULL;
    int rc;

    if (str == NULL) {
        return NULL;
    }

    rc = init_queue(&queue);
    if (rc != 0) {
        return NULL;
    }

    root = create_tree_internal(str, &queue);
    clean_queue(&queue);
    return root;
}

void
dump_tree (tree_node_t *root)
{
    bool print_comma = FALSE;
    uqueue_t queue;
    tree_node_t *p;
    int rc;

    rc = init_queue(&queue);
    if (rc != 0) {
        return;
    }

    enqueue(&queue, root);
    while (is_queue_empty(&queue) == FALSE) {
        if (print_comma == TRUE) {
            printf(", ");
        } else {
            print_comma = TRUE;
        }

        p = dequeue(&queue);
        if (p == NULL) {
            printf("#");
            continue;
        }

        printf("%d", p->val);
        if (p->left == NULL && p->right == NULL) {
            continue;
        }
        enqueue(&queue, p->left);
        enqueue(&queue, p->right);
    }

    printf("\n");
    clean_queue(&queue);
    return;
}

