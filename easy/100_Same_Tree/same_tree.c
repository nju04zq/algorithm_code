#include <stddef.h>

typedef unsigned char bool;

#define TRUE 1
#define FALSE 0

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

typedef struct TreeNode tree_node_t;

bool
isSameTree (tree_node_t *p, tree_node_t *q)
{
    bool rc;

    if (p == NULL && q == NULL) {
        return TRUE;
    } else if (p == NULL || q == NULL) {
        return FALSE;
    }

    if (p->val != q->val) {
        return FALSE;
    }

    rc = isSameTree(p->left, q->left);
    if (rc == FALSE) {
        return FALSE;
    }

    rc = isSameTree(p->right, q->right);
    if (rc == FALSE) {
        return FALSE;
    }
    
    return TRUE;
}

#include <stdio.h>

static void
dump_tree (tree_node_t *p)
{
    if (p == NULL) {
        return;
    }

    printf("%d, ", p->val);
    if (p->left) {
        printf("%d, ", p->left->val);
    } else {
        printf("-, ");
    }
    if (p->right) {
        printf("%d, ", p->right->val);
    } else {
        printf("-, ");
    }
    printf("\n");

    dump_tree(p->left);
    dump_tree(p->right);
    return;
}

static void
test_is_same_tree (tree_node_t *p, tree_node_t *q)
{
    bool rc;

    printf("Tree #1:\n");
    dump_tree(p);
    printf("Tree #2:\n");
    dump_tree(q);

    rc = isSameTree(p, q);
    printf("Is the same %d\n\n", rc);
    return;
}

#define DEF_NODE(x, y) tree_node_t n##x##y = {y, NULL, NULL};

int main (void)
{
    DEF_NODE(1, 1);
    DEF_NODE(1, 2);
    DEF_NODE(1, 3);
    DEF_NODE(2, 1);
    DEF_NODE(2, 2);
    DEF_NODE(2, 3);

    n11.left = &n12;
    n11.right = &n13;
    n21.left = &n23;
    n21.right = &n22;

    test_is_same_tree(&n11, &n21);

    n11.left = &n12;
    n11.right = &n13;
    n21.left = &n22;
    n21.right = &n23;

    test_is_same_tree(&n11, &n21);
    return 0;
}

