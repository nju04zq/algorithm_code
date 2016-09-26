#include <stddef.h>
#include "../../common/tree.h"

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

#define MIN(a, b) ((a) < (b) ? (a) : (b))
#define MAX(a, b) ((a) > (b) ? (a) : (b))

int
minDepth (tree_node_t *root)
{
    int l_min, r_min, min_depth;

    if (root == NULL) {
        return 0;
    }

    l_min = minDepth(root->left);
    r_min = minDepth(root->right);
    if (l_min == 0 || r_min == 0) {
        min_depth = MAX(l_min, r_min) + 1;
    } else {
        min_depth = MIN(l_min, r_min) + 1;
    }
    return min_depth;
}

#include <stdio.h>

static void
test_min_depth (char *str)
{
    int min_depth;
    tree_node_t *root;

    root = create_tree(str);
    min_depth = minDepth(root);
    printf("Tree \"%s\", min depth %d\n", str, min_depth);
    destory_tree(root);
    return;
}

int main (void)
{
    test_min_depth("");
    test_min_depth("1");
    test_min_depth("1, 2");
    test_min_depth("1, 2, #, 3");
    test_min_depth("1, 2, 3, 4, 5");
    return 0;
}

