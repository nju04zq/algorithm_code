#include <stddef.h>
#include <stdlib.h>
#include "../../common/tree.h"
#include "../../common/queue.h"

typedef struct level_output_s {
    int *val_buf;
    int col_size;
} level_output_t;

static void
free_level_output (level_output_t *output_p)
{
    free(output_p->val_buf);
    free(output_p);
    return;
}

static level_output_t *
create_level_output (int col_size)
{
    level_output_t *p;

    p = calloc(1, sizeof(level_output_t));
    if (p == NULL) {
        return NULL;
    }

    p->val_buf = calloc(col_size, sizeof(int));
    if (p->val_buf == NULL) {
        free(p);
        return NULL;
    }

    p->col_size = col_size;
    return p;
}

static void
make_level_output (uqueue_t *level_queue, uqueue_t *output_queue)
{
    int queue_size, i = 0, rc;
    level_output_t *output_p;
    tree_node_t *p;

    queue_size = get_queue_size(level_queue);
    output_p = create_level_output(queue_size);
    if (!output_p) {
        return;
    }

    p = peek_queue_head(level_queue);
    for (i = 0; i < queue_size; i++) {
        if (p != NULL) {
            output_p->val_buf[i] = p->val;
        }
        p = get_queue_next(level_queue);
    }

    rc = enqueue(output_queue, output_p);
    if (rc != 0) {
        free_level_output(output_p);
    }
    return;
}

static void
clean_level_queue (uqueue_t *level_queue)
{
    int i, size;

    size = get_queue_size(level_queue);
    for (i = 0; i < size; i++) {
        (void)dequeue(level_queue);
    }
    return;
}

static void
iter_level_queue (uqueue_t *level_queue)
{
    int i, queue_size;
    tree_node_t *p, **buf;

    queue_size = get_queue_size(level_queue);
    buf = calloc(queue_size, sizeof(tree_node_t *));
    if (buf == NULL) {
        clean_level_queue(level_queue);
        return;
    }

    p = peek_queue_head(level_queue);
    for (i = 0; i < queue_size; i++) {
        buf[i] = p;
        p = get_queue_next(level_queue);
    }

    clean_level_queue(level_queue);

    for (i = 0; i < queue_size; i++) {
        p = buf[i];
        if (p == NULL) {
            continue;
        }
        if (p->left != NULL) {
            enqueue(level_queue, p->left);
        }
        if (p->right != NULL) {
            enqueue(level_queue, p->right);
        }
    }

    free(buf);
    return;
}

static int **
make_level_order_output (uqueue_t *output_queue, int **col_sizes, int *ret_size)
{
    int i, queue_size, **output;
    level_output_t *level_output_p;

    queue_size = get_queue_size(output_queue);
    output = calloc(queue_size, sizeof(int *));
    if (output == NULL) {
        return NULL;
    }
    *col_sizes = calloc(queue_size, sizeof(int));
    if (*col_sizes == NULL) {
        free(output);
        return NULL;
    }

    *ret_size = queue_size;
    for (i = 0; i < queue_size; i++) {
        level_output_p = dequeue(output_queue);
        output[i] = level_output_p->val_buf;
        (*col_sizes)[i] = level_output_p->col_size;
        free(level_output_p);
    }

    return output;
}

int **
levelOrder (tree_node_t *root, int **col_sizes, int *ret_size)
{
    uqueue_t level_queue, output_queue;
    int **output;
    int rc;

    *col_sizes = NULL;
    *ret_size = 0;

    if (root == NULL || col_sizes == NULL || ret_size == NULL) {
        return NULL;
    }

    rc = init_queue(&level_queue);
    if (rc != 0) {
        return NULL;
    }
    rc = init_queue(&output_queue);
    if (rc != 0) {
        return NULL;
    }

    enqueue(&level_queue, root);
    for (;;) {
        if (is_queue_empty(&level_queue)) {
            break;
        }
        make_level_output(&level_queue, &output_queue);
        iter_level_queue(&level_queue);
    }

    output = make_level_order_output(&output_queue, col_sizes, ret_size);

    clean_queue(&level_queue);
    clean_queue(&output_queue);
    return output;
}

#include <stdio.h>

static void
dump_one_level (int *level, int size)
{
    int i;

    for (i = 0; i < size; i++) {
        if (i != 0) {
            printf(", ");
        }
        printf("%d", level[i]);
    }
    return;
}

static void
dump_level_order (int **output, int *col_sizes, int ret_size)
{
    int i;

    for (i = 0; i < ret_size; i++) {
        printf("[");
        dump_one_level(output[i], col_sizes[i]);
        printf("]");
    }
    printf("\n");
    return;
}

static void
test_level_order (char *str)
{
    int **output, *col_sizes, ret_size;
    tree_node_t *root;

    root = create_tree(str);
    output = levelOrder(root, &col_sizes, &ret_size);

    printf("Level order traversal on %s:\n", str);
    if (output == NULL) {
        printf("<NULL>\n");
    } else {
        dump_level_order(output, col_sizes, ret_size);
        free(col_sizes);
        free(output);
    }

    destory_tree(root);
    return;
}

int main (void)
{
    test_level_order("");
    test_level_order("1");
    test_level_order("1, 2");
    test_level_order("12, 23, #, #, 34, 45, 567");
    return 0;
}

