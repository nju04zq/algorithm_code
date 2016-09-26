#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

typedef struct ustack_s {
    int *buf;
    int top;
    int size;
    int max_size;
} ustack_t;

typedef struct {
    ustack_t org_stack;
    ustack_t min_stack;
} MinStack;

#define MIN(a, b) ((a) < (b) ? (a) : (b))

static bool
is_stack_empty (ustack_t *stack)
{
    if (stack->size == 0) {
        return TRUE;
    } else {
        return FALSE;
    }
}

static bool
is_stack_full (ustack_t *stack)
{
    if (stack->size == stack->max_size) {
        return TRUE;
    } else {
        return FALSE;
    }
}

static void
pop (ustack_t *stack)
{
    if (is_stack_empty(stack)) {
        return;
    }

    stack->top--;
    stack->size--;
    return;
}

static void
push (ustack_t *stack, int x)
{
    if (is_stack_full(stack)) {
        return;
    }

    stack->top++;
    stack->size++;
    stack->buf[stack->top] = x;
    return;
}

static int
peek_stack_top (ustack_t *stack)
{
    if (is_stack_empty(stack)) {
        return -1;
    } else {
        return stack->buf[stack->top];
    }
}

static void
clean_stack (ustack_t *stack)
{
    if (stack->buf) {
        free(stack->buf);
    }
    return;
}

static int
init_stack (ustack_t *stack, int max_size)
{
    memset(stack, 0, sizeof(ustack_t)); 

    stack->buf = calloc(max_size, sizeof(int));
    if (stack->buf == NULL) {
        return -1;
    }

    stack->top = -1;
    stack->size = 0;
    stack->max_size = max_size;
    return 0;
}

void
minStackCreate (MinStack *stack, int maxSize)
{
    int rc;

    rc = init_stack(&stack->org_stack, maxSize);
    if (rc != 0) {
        return;
    }

    rc = init_stack(&stack->min_stack, maxSize);
    if (rc != 0) {
        clean_stack(&stack->org_stack);
        return;
    }

    return;
}

void
minStackPush (MinStack *stack, int element)
{
    int x;

    push(&stack->org_stack, element);

    if (is_stack_empty(&stack->min_stack)) {
        x = element;
    } else {
        x = peek_stack_top(&stack->min_stack);
        x = MIN(x, element);
    }
    push(&stack->min_stack, x);
    return;
}

void
minStackPop (MinStack *stack)
{
    pop(&stack->org_stack);
    pop(&stack->min_stack);
    return;
}

int
minStackTop (MinStack *stack)
{
    int x;

    x = peek_stack_top(&stack->org_stack);
    return x;
}

int
minStackGetMin (MinStack *stack)
{
    int x;

    x = peek_stack_top(&stack->min_stack);
    return x;
}

void
minStackDestroy (MinStack *stack)
{
    clean_stack(&stack->org_stack);
    clean_stack(&stack->min_stack);
    return;
}

#include <stdio.h>

static void
dump_stack (MinStack *stack)
{
    int top, min, i;

    top = minStackTop(stack);
    min = minStackGetMin(stack);
    printf("Stack top %d, min %d, content ", top, min);

    for (i = 0; i < stack->org_stack.size; i++) {
        printf("%d ", stack->org_stack.buf[i]);
    }
    printf("\n");
    return;
}

static void
test_pop (MinStack *stack)
{
    //printf("Before pop, ");
    //dump_stack(stack);
    minStackPop(stack);
    printf("After pop, ");
    dump_stack(stack);
    return;
}

static void
test_push (MinStack *stack, int x)
{
    //printf("Before push, ");
    //dump_stack(stack);
    minStackPush(stack, x);
    printf("After push, ");
    dump_stack(stack);
    return;
}

static void
test_min_stack (void)
{
    int a[] = {5, 2, 3, 1, 2 ,3, 4, 1}, size, i;
    MinStack stack;

    minStackCreate(&stack, 100);

    size = sizeof(a)/sizeof(a[0]);
    for (i = 0; i < size; i++) {
        test_push(&stack, a[i]);
    }
    for (i = 0; i < size; i++) {
        test_pop(&stack);
    }

    minStackDestroy(&stack);
}

int main (void)
{
    test_min_stack();
    return 0;
}

