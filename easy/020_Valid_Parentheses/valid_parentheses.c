#include <stdlib.h>
#include <stddef.h>
#include <string.h>

typedef unsigned char bool;

#define TRUE 1
#define FALSE 0

typedef struct ustack_s {
    char *buf;
    int top;
    int max_size;
} ustack_t; //stack_t conflicts with one on OS X...

#define STACK_INIT_SIZE 2

static bool
stack_is_full (ustack_t *stack_p)
{
    if (stack_p->top >= (stack_p->max_size-1)) {
        return TRUE;
    } else {
        return FALSE;
    }
}

static bool
stack_is_empty (ustack_t *stack_p)
{
    if (stack_p->top == -1) {
        return TRUE;
    } else {
        return FALSE;
    }
}

static int
stack_resize (ustack_t *stack_p, int new_max_size)
{
    char *new_buf;
    int cur_size;

    new_buf = calloc(new_max_size, sizeof(char));
    if (new_buf == NULL) {
        return -1;
    }

    cur_size = stack_p->top + 1;
    memcpy(new_buf, stack_p->buf, cur_size*sizeof(char));

    stack_p->buf = new_buf;
    stack_p->max_size = new_max_size;
    return 0;
}

static int
stack_expand (ustack_t *stack_p)
{
    int rc, new_max_size;

    new_max_size = stack_p->max_size * 2;
    rc = stack_resize(stack_p, new_max_size);
    return rc;
}

static void
stack_shrink (ustack_t *stack_p)
{
    int new_max_size, cur_size;


    if (stack_p->max_size <= STACK_INIT_SIZE) {
        return;
    }

    cur_size = stack_p->top + 1;
    if (cur_size >= stack_p->max_size/4) {
        return;
    }

    new_max_size = stack_p->max_size/2;
    (void)stack_resize(stack_p, new_max_size);
    return;
}

static int
stack_push (ustack_t *stack_p, char val)
{
    int rc = 0;

    if (stack_is_full(stack_p)) {
        rc = stack_expand(stack_p);
    }
    if (rc != 0) {
        return rc;
    }

    stack_p->buf[++stack_p->top] = val;
    return 0;
}

static int 
stack_pop (ustack_t *stack_p, char *val_p)
{
    *val_p = '\0';

    if (stack_is_empty(stack_p)) {
        return -1;
    }

    *val_p = stack_p->buf[stack_p->top--];

    stack_shrink(stack_p);
    return 0;
}

static void
stack_clean (ustack_t *stack_p)
{
    free(stack_p->buf);
    memset(stack_p, 0, sizeof(ustack_t));
    return;
}

static int
stack_init (ustack_t *stack_p)
{
    stack_p->buf = calloc(STACK_INIT_SIZE, sizeof(char));
    if (stack_p->buf == NULL) {
        return -1;
    }
    stack_p->top = -1;
    stack_p->max_size = STACK_INIT_SIZE;
    return 0;
}

bool
check_pair (char prev, char cur)
{
    if (prev == '(' && cur == ')') {
        return TRUE;
    }
    if (prev == '{' && cur == '}') {
        return TRUE;
    }
    if (prev == '[' && cur == ']') {
        return TRUE;
    }
    return FALSE;
}

static bool
is_pair (char left, char right)
{
    if (left == '(' && right == ')') {
        return TRUE;
    }
    if (left == '{' && right == '}') {
        return TRUE;
    }
    if (left == '[' && right == ']') {
        return TRUE;
    }
    return FALSE;
}

static bool
is_left (char c)
{
    if (c == '{' || c == '[' || c == '(') {
        return TRUE;
    } else {
        return FALSE;
    }
}

bool
isValid (char *s)
{
    char left;
    bool is_valid = TRUE;
    int rc;
    ustack_t stack;

    if (s == NULL || *s == '\0') {
        return TRUE;
    }

    stack_init(&stack);

    for (; *s; s++) {
        if (is_left(*s)) {
            rc = stack_push(&stack, *s);
            if (rc == -1) {
                is_valid = FALSE;
                break;
            }
            continue;
        }

        if (stack_is_empty(&stack)) {
            is_valid = FALSE;
            break;
        }

        (void)stack_pop(&stack, &left);
        if (is_pair(left, *s) == FALSE) {
            is_valid = FALSE;
            break;
        }
    }

    if (is_valid && stack_is_empty(&stack) == FALSE) {
        is_valid = FALSE;
    }

    stack_clean(&stack);
    return is_valid;
}

#include <stdio.h>

static void
test_check (char *s)
{
    bool result;

    result = isValid(s);
    printf("Check result for %s, %d\n", s, result);
    return;
}

int main (void)
{
    test_check("");
    test_check("[");
    test_check("]");
    test_check("[]");
    test_check("[[]");
    test_check("[{}]");
    test_check("[{]}");
    test_check("[{}][](){}{([({{{}}})])}");
    return 0;
}

