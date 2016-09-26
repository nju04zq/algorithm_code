#include <stdio.h>
#include "tree.h"

static void
test_tree (char *str)
{
    tree_node_t *root;

    printf("Test tree: %s\n", str);

    root = create_tree(str);
    dump_tree(root);
    destory_tree(root);
    return;
}

int main (void)
{
    test_tree("#");
    test_tree("1");
    test_tree("1, 2, 3");
    test_tree("1, 2, #, #, 3, 4, 5");
    test_tree("12, 23, #, #, 34, 45, 567");
    return 0;
}

