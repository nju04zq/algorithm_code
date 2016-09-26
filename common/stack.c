#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "bool.h"
#include "stack.h"

#ifdef __USTACK_TEST_DEBUG__
#define USTACK_INIT_SIZE 4 
#else
#define USTACK_INIT_SIZE 64
#endif

bool
is_stack_full (ustack_t *stack)
{
    if (stack->size == stack->max_size) {
        return TRUE;
    } else {
        return FALSE;
    }
}

bool
is_stack_empty (ustack_t *stack)
{
    if (stack->size == 0) {
        return TRUE;
    } else {
        return FALSE;
    }
}

int
get_stack_size (ustack_t *stack)
{
    return stack->size;
}

static int
resize_stack (ustack_t *stack, int new_max_size)
{
    void **buf;

    if (stack->size > new_max_size) {
        return -1;
    }

    buf = calloc(new_max_size, sizeof(void *));
    if (buf == NULL) {
        return -1;
    }

    memcpy(buf, stack->buf, stack->size*sizeof(void *));
    free(stack->buf);

    stack->buf = buf;
    stack->max_size = new_max_size;
    return 0;
}

static void
shrink_stack (ustack_t *stack)
{
    int new_max_size;

    if (stack->max_size <= USTACK_INIT_SIZE) {
        return;
    }
    if (stack->size > (stack->max_size/4)) {
        return;
    }

    new_max_size = stack->max_size/2;
    (void)resize_stack(stack, new_max_size);
    return;
}

static int
expand_stack (ustack_t *stack)
{
    int new_max_size, rc;

    new_max_size = stack->max_size * 2;
    rc = resize_stack(stack, new_max_size);
    return rc;
}

void *
pop (ustack_t *stack)
{
    void *udata;

    if (is_stack_empty(stack)) {
        return NULL;
    }

    udata = stack->buf[stack->top];
    stack->top--;
    stack->size--;

    shrink_stack(stack);
    return udata;
}

int
push (ustack_t *stack, void *udata)
{
    int rc;

    if (is_stack_full(stack)) {
        rc = expand_stack(stack);
        if (rc != 0) {
            return rc;
        }
    }

    stack->size++;
    stack->top++;
    stack->buf[stack->top] = udata;
    return 0;
}

static void *
get_stack_entry (ustack_t *stack)
{
    return stack->buf[stack->iter_index];
}

void *
peek_stack_top (ustack_t *stack)
{
    void *udata;

    stack->iter_index = stack->size - 1;

    if (is_stack_empty(stack)) {
        return NULL;
    }

    udata = get_stack_entry(stack);
    return udata;
}

void *
peek_stack_bottom (ustack_t *stack)
{
    void *udata;

    stack->iter_index = 0;

    if (is_stack_empty(stack)) {
        return NULL;
    }

    udata = get_stack_entry(stack);
    return udata;
}

void *
get_stack_next (ustack_t *stack)
{
    void *udata;

    if (stack->iter_index >= (stack->size - 1)) {
        return NULL;
    }

    stack->iter_index++;
    udata = get_stack_entry(stack);
    return udata;
}

void
clean_stack (ustack_t *stack)
{
    free(stack->buf);
    return;
}

int
init_stack (ustack_t *stack)
{
    memset(stack, 0, sizeof(ustack_t));
    stack->top = -1;
    stack->iter_index = 0;
    stack->size = 0;
    stack->max_size = USTACK_INIT_SIZE;

    stack->buf = calloc(stack->max_size, sizeof(void *));
    if (stack->buf == NULL) {
        return -1;
    }
    return 0;
}

void
dump_stack (ustack_t *stack, void (*print_udata)(void *))
{
    int i, size;
    void *udata;
    bool first_entry = TRUE;

    if (is_stack_empty(stack)) {
        printf("<Empty stack>\n");
        return;
    }

    size = get_stack_size(stack);
    for (i = 0; i < size; i++) {
        if (i == 0) {
            udata = peek_stack_bottom(stack);
        } else {
            udata = get_stack_next(stack);
        }
        if (first_entry) {
            first_entry = FALSE;
        } else {
            printf(", ");
        }
        print_udata(udata);
    }

    printf("\n");
    return;
}

