#include <stddef.h>
#include "../../common/bool.h"
#include "../../common/tree.h"

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

static bool
is_trees_symmetric (tree_node_t *root1, tree_node_t *root2)
{
    bool rc;

    if (root1 == NULL && root2 == NULL) {
        return TRUE;
    }
    if (root1 == NULL || root2 == NULL) {
        return FALSE;
    }
    if (root1->val != root2->val) {
        return FALSE;
    }

    rc = is_trees_symmetric(root1->left, root2->right);
    if (rc == FALSE) {
        return FALSE;
    }
    rc = is_trees_symmetric(root1->right, root2->left);
    if (rc == FALSE) {
        return FALSE;
    }

    return TRUE;
}

bool
isSymmetric (tree_node_t *root)
{
    bool rc;

    if (root == NULL) {
        return TRUE;
    }

    rc = is_trees_symmetric(root->left, root->right);
    return rc;
}

#include <stdio.h>

static void
test_is_symmetric (char *str)
{
    tree_node_t *root;
    bool rc;

    root = create_tree(str);
    rc = isSymmetric(root);
    printf("Tree %s is symmetric, %d\n", str, rc);
    destory_tree(root);
    return;
}

int main (void)
{
    test_is_symmetric("1");
    test_is_symmetric("1, 2, 2, 3, 4, 4, 3");
    test_is_symmetric("1, 2, 2, #, 3, #, 3");
    return 0;
}

