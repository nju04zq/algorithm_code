#include <stddef.h>
#include <stdlib.h>
#include "../../common/stack.h"
#include "../../common/queue.h"
#include "../../common/tree.h"

typedef struct tree_path_s {
    int *val_buf;
    int val_cnt;
} tree_path_t;

typedef struct node_rec_s {
    tree_node_t *node_p;
    tree_node_t *next_p;
} node_rec_t;

static bool
is_leaf_node (tree_node_t *node_p)
{
    if (node_p->left == NULL && node_p->right == NULL) {
        return TRUE;
    } else {
        return FALSE;
    }

}

static void
pop_node_rec (ustack_t *stack)
{
    node_rec_t *rec_p;

    rec_p = pop(stack);
    if (rec_p) {
        free(rec_p);
    }
    return;
}

static int
push_node_rec (ustack_t *stack, tree_node_t *node_p)
{
    node_rec_t *rec_p;
    int rc;

    rec_p = calloc(1, sizeof(node_rec_t));
    if (rec_p == NULL) {
        return -1;
    }

    rec_p->node_p = node_p;
    if (node_p->left) {
        rec_p->next_p = node_p->left;
    } else {
        rec_p->next_p = node_p->right;
    }

    rc = push(stack, rec_p);
    if (rc != 0) {
        free(rec_p);
        return -1;
    }

    return 0;
}

static void
copy_to_tree_path (tree_path_t *path_p, ustack_t *stack)
{
    int i, stack_size;
    node_rec_t *rec_p;

    stack_size = get_stack_size(stack);
    for (i = 0; i < stack_size; i++) {
        if (i == 0) {
            rec_p = peek_stack_bottom(stack);
        } else {
            rec_p = get_stack_next(stack);
        }
        path_p->val_buf[i] = rec_p->node_p->val;
    }
    return;
}

static tree_path_t *
make_tree_path (ustack_t *stack)
{
    int stack_size;
    tree_path_t *path_p;

    stack_size = get_stack_size(stack);

    path_p = calloc(1, sizeof(tree_path_t));
    if (path_p == NULL) {
        return NULL;
    }

    path_p->val_buf = calloc(stack_size, sizeof(int));
    if (path_p->val_buf == NULL) {
        free(path_p);
        return NULL;
    }

    path_p->val_cnt = stack_size;
    copy_to_tree_path(path_p, stack);
    return path_p;
}

static int
add_tree_path (uqueue_t *queue, ustack_t *stack)
{
    tree_path_t *path_p;
    int rc;

    path_p = make_tree_path(stack);
    if (path_p == NULL) {
        return -1;
    }

    rc = enqueue(queue, path_p);
    if (rc != 0) {
        free(path_p);
    }

    return 0;
}

static int
path_sum_internal (tree_node_t *root, int sum,
                   ustack_t *stack, uqueue_t *queue)
{
    int rc = 0, stack_sum = 0;
    node_rec_t *rec_p;

    rc = push_node_rec(stack, root);
    if (rc != 0) {
        return -1;
    }

    stack_sum = root->val;
    while (is_stack_empty(stack) == FALSE) {
        rec_p = peek_stack_top(stack);

        if (is_leaf_node(rec_p->node_p) && stack_sum == sum) {
            rc = add_tree_path(queue, stack);
            if (rc != 0) {
                break;
            }
        }

        if (rec_p->next_p == NULL) {
            stack_sum -= rec_p->node_p->val;
            pop_node_rec(stack);
            continue;
        }

        rc = push_node_rec(stack, rec_p->next_p);
        if (rc != 0) {
            break;
        }

        stack_sum += rec_p->next_p->val;

        if (rec_p->next_p == rec_p->node_p->left) {
            rec_p->next_p = rec_p->node_p->right;
        } else {
            rec_p->next_p = NULL;
        }
    }

    return rc;
}

static void
free_tree_path_output (int **output, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        if (output[i]) {
            free(output[i]);
        }
    }
    return;
}

static int *
copy_out_one_tree_path (tree_path_t *path_p, int *col_size)
{
    int i, *path_buf;

    path_buf = calloc(path_p->val_cnt, sizeof(int));
    if (path_buf == NULL) {
        return NULL;
    }

    for (i = 0; i < path_p->val_cnt; i++) {
        path_buf[i] = path_p->val_buf[i];
    }

    *col_size = path_p->val_cnt;
    return path_buf;
}

static int **
copy_out_tree_path_queue (uqueue_t *queue, int **columnSizes, int *returnSize)
{
    int **output, queue_size, *col_buf, i, col_size;
    tree_path_t *path_p;

    queue_size = get_queue_size(queue);

    output = calloc(queue_size, sizeof(int *));
    if (output == NULL) {
        return NULL;
    }
    col_buf = calloc(queue_size, sizeof(int));
    if (col_buf == NULL) {
        free(output);
        return NULL;
    }

    for (i = 0; i < queue_size; i++) {
        if (i == 0) {
            path_p = peek_queue_head(queue);
        } else {
            path_p = get_queue_next(queue);
        }
        output[i] = copy_out_one_tree_path(path_p, &col_size);
        if (output[i] == NULL) {
            free_tree_path_output(output, queue_size);
            free(col_buf);
            free(output);
            return NULL;
        }
        col_buf[i] = col_size;
    }

    *columnSizes = col_buf;
    *returnSize = queue_size;
    return output;
}

static void
clean_all_entry_on_stack (ustack_t *stack)
{
    int i, stack_size;
    node_rec_t *rec_p;

    stack_size = get_stack_size(stack);
    for (i = 0; i < stack_size; i++) {
        rec_p = pop(stack);
        if (rec_p) {
            free(rec_p);
        }
    }
    return;
}

static void
clean_all_entry_in_queue (uqueue_t *queue)
{
    int i, queue_size;
    tree_path_t *path_p;

    queue_size = get_queue_size(queue);
    for (i = 0; i < queue_size; i++) {
        path_p = dequeue(queue);
        if (path_p) {
            free(path_p);
        }
    }
    return;
}

int **
pathSum (tree_node_t *root, int sum, int **columnSizes, int *returnSize)
{
    int rc, **output = NULL;
    ustack_t stack;
    uqueue_t queue;

    *columnSizes = 0;
    *returnSize = 0;

    if (root == NULL) {
        return NULL;
    }

    rc = init_stack(&stack);
    if (rc != 0) {
        return NULL;
    }
    rc = init_queue(&queue);
    if (rc != 0) {
        clean_stack(&stack);
        return NULL;
    }

    rc = path_sum_internal(root, sum, &stack, &queue);
    if (rc == 0) {
        output = copy_out_tree_path_queue(&queue, columnSizes, returnSize);
    }

    clean_all_entry_on_stack(&stack);
    clean_all_entry_in_queue(&queue);
    return output;
}

#include <stdio.h>

static void
dump_one_path (int idx, int *path, int path_size)
{
    int i;

    printf("Path #%d: ", idx);
    for (i = 0; i < path_size; i++) {
        printf("%d ", path[i]);
    }
    printf("\n");
    return;
}

static void
dump_all_path (int **paths, int *col_sizes, int path_cnt)
{
    int i;

    if (paths == NULL || col_sizes == NULL || path_cnt == 0) {
        printf("<No Path>\n");
    }

    for (i = 0; i < path_cnt; i++) {
        dump_one_path(i, paths[i], col_sizes[i]);
    }
    return;
}

static void
test_path_sum (char *str, int sum)
{
    tree_node_t *root;
    int **output, *col_sizes, ret_size;

    printf("All paths for tree \"%s\" with sum %d:\n", str, sum);

    root = create_tree(str);
    output = pathSum(root, sum, &col_sizes, &ret_size);
    dump_all_path(output, col_sizes, ret_size);
    destory_tree(root);
    return;
}

int main (void)
{
    test_path_sum("", 0);
    test_path_sum("1", 1);
    test_path_sum("1", 2);
    test_path_sum("1, 2", 1);
    test_path_sum("1, 2, 4, #, #, 6, 8", 2);
    test_path_sum("1, 2, 4, #, #, 6, 8", 3);
    test_path_sum("1, 2, 4, #, #, 6, 8", 5);
    test_path_sum("1, 2, 4, #, #, 6, 8", 10);
    test_path_sum("1, 2, 4, #, #, 6, 8", 11);
    test_path_sum("1, 2, 4, #, #, 6, 8", 13);
    test_path_sum("1, 2, 3, 3, #, 2, 4, #, 2, 1", 7);
    test_path_sum("1, 2, 3, 3, #, 2, 4, #, 2, 1", 8);
    return 0;
}

