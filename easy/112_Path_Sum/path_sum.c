#include <stddef.h>
#include "../../common/bool.h"
#include "../../common/tree.h"

static bool
hasPathSum_internal (tree_node_t *root, int sum)
{
    bool rc;
    int sum_left;

    sum_left = sum - root->val;
    if (root->left == NULL && root->right == NULL) {
        if (sum_left == 0) {
            return TRUE;
        } else {
            return FALSE;
        }
    }

    if (root->left) {
        rc = hasPathSum_internal(root->left, sum_left);
        if (rc == TRUE) {
            return TRUE;
        }
    }

    if (root->right) {
        rc = hasPathSum_internal(root->right, sum_left);
        if (rc == TRUE) {
            return TRUE;
        }
    }

    return FALSE;
}

bool
hasPathSum (tree_node_t *root, int sum)
{
    bool rc;

    if (root == NULL) {
        return FALSE;
    }

    rc = hasPathSum_internal(root, sum);
    return rc;
}

#include <stdio.h>

static void
test_has_path_sum (char *str, int sum)
{
    bool rc;
    tree_node_t *root;

    root = create_tree(str);
    rc = hasPathSum(root, sum);
    printf("Tree \"%s\" has path sum %d, %d\n", str, sum, rc);
    destory_tree(root);
    return;
}

int main (void)
{
    test_has_path_sum("", 1);
    test_has_path_sum("1", 1);
    test_has_path_sum("1", 2);
    test_has_path_sum("1, 2", 1);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 2);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 3);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 5);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 10);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 11);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 13);
    test_has_path_sum("1, 2, 4, #, #, 6, 8", 14);
    return 0;
}

