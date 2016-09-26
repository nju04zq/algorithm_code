#ifndef __STACK_H__
#define __STACK_H__

#include "bool.h"

typedef struct ustack_s {
    void **buf;
    int top;
    int iter_index;
    int size;
    int max_size;
} ustack_t;

bool
is_stack_full(ustack_t *stack);

bool
is_stack_empty(ustack_t *stack);

int
get_stack_size(ustack_t *stack);

void *
pop(ustack_t *stack);

int
push(ustack_t *stack, void *udata);

void *
peek_stack_top(ustack_t *stack);

void *
peek_stack_bottom(ustack_t *stack);

void *
get_stack_next(ustack_t *stack);

void
clean_stack(ustack_t *stack);

int
init_stack(ustack_t *stack);

void
dump_stack(ustack_t *stack, void (*print_udata)(void *));
#endif //__STACK_H__
