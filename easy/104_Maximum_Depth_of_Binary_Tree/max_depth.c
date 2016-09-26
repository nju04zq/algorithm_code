#include <stddef.h>
#include "../../common/tree.h"

#define MAX(a, b) ((a) > (b) ? (a) : (b))

int
maxDepth (tree_node_t *root)
{
    int lchild_depth, rchild_depth, max_depth;

    if (root == NULL) {
        return 0;
    }

    lchild_depth = maxDepth(root->left);
    rchild_depth = maxDepth(root->right);

    max_depth = MAX(lchild_depth, rchild_depth) + 1;
    return max_depth;
}

#include <stdio.h>

static void
test_maxdepth (char *str)
{
    tree_node_t *root;
    int max_depth;

    root = create_tree(str);
    max_depth = maxDepth(root);
    printf("Tree \"%s\" max depth %d\n", str, max_depth);
    return;
}

int main (void)
{
    test_maxdepth("");
    test_maxdepth("1");
    test_maxdepth("1, 2, #, #, 3, 4");
    return 0;
}

