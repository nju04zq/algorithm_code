//******************NOTE*******************
//For this problem, a height-balanced binary tree is defined as
//a binary tree in which the depth of the two subtrees of every
//node never differ by more than 1.
//******************NOTE*******************
#include <stddef.h>
#include "../../common/bool.h"
#include "../../common/tree.h"

#define MAX(a, b) ((a) > (b) ? (a) : (b))
#define ABS(x) ((x) >= 0 ? (x) : -(x))

static bool
is_balanced_internal (tree_node_t *root, int *depth)
{
    int l_depth, r_depth;
    bool rc;

    *depth = 0;

    if (root == NULL) {
        return TRUE;
    }

    rc = is_balanced_internal(root->left, &l_depth);
    if (rc == FALSE) {
        return FALSE;
    }
    rc = is_balanced_internal(root->right, &r_depth);
    if (rc == FALSE) {
        return FALSE;
    }

    if (ABS(l_depth - r_depth) > 1) {
        return FALSE;
    }

    *depth = MAX(l_depth, r_depth) + 1;
    return TRUE;
}

bool
isBalanced (tree_node_t *root)
{
    int depth;
    bool rc;

    rc = is_balanced_internal(root, &depth);
    return rc;
}

#include <stdio.h>

static void
test_is_balanced (char *str)
{
    tree_node_t *root;
    bool rc;

    root = create_tree(str);
    rc = isBalanced(root);
    printf("Tree \"%s\" is balanced, %d\n", str, rc);

    destory_tree(root);
    return;
}

int main (void)
{
    test_is_balanced("");
    test_is_balanced("1");
    test_is_balanced("1, 2");
    test_is_balanced("1, 2, 3");
    test_is_balanced("1, 2, #, 3");
    test_is_balanced("1, 2, 3, 4, 5, #, 6, 7");
    test_is_balanced("1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, #, #, 5, 5");
    test_is_balanced("1, 2, 3, 4, 5, 6, #, 7");
    return 0;
}

