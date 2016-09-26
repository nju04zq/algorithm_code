#ifndef __TREE_H__
#define __TREE_H__

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

typedef struct TreeNode tree_node_t;

tree_node_t *
create_tree(char *str);

void
destory_tree(tree_node_t *root);

void
dump_tree(tree_node_t *root);
#endif //__TREE_H__
