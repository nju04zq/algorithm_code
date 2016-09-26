#include <stdio.h>
#include "stack.h"

static void
print_stack_entry (void *data)
{
    printf("%d", *(int *)data);
    return;
}

static void
test_case_1 (void)
{
    int a[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12};
    int i, size, rc, *top, *bottom;
    void *entry;
    ustack_t stack;

    size = sizeof(a)/sizeof(a[0]);

    rc = init_stack(&stack);
    if (rc != 0) {
        return;
    }

    printf("After init### ");
    dump_stack(&stack, print_stack_entry);

    for (i = 0; i < size/2; i++) {
        push(&stack, &a[i]);
        printf("After push %d### ", a[i]);
        dump_stack(&stack, print_stack_entry);
    }

    for (i = 0; i < 3; i++) {
        entry = pop(&stack);
        if (!entry) {
            break;
        }
        printf("After pop %d### ", *(int *)entry);
        dump_stack(&stack, print_stack_entry);
    }

    for (i = size/2; i < size; i++) {
        push(&stack, &a[i]);
        printf("After enqueue %d### ", a[i]);
        dump_stack(&stack, print_stack_entry);
    }

    top = peek_stack_top(&stack);
    if (top) {
        printf("Stack TOP %d\n", *top);
    } else {
        printf("Stack TOP NULL\n");
    }
    bottom = peek_stack_bottom(&stack);
    if (bottom) {
        printf("Stack BOTTOM %d\n", *bottom);
    } else {
        printf("Stack BOTTOM NULL\n");
    }

    for (;;) {
        entry = pop(&stack);
        if (!entry) {
            break;
        }
        printf("After pop %d### ", *(int *)entry);
        dump_stack(&stack, print_stack_entry);
    }

    printf("Before clean### ");
    dump_stack(&stack, print_stack_entry);
    clean_stack(&stack);

    return;
}

int main (void)
{
    test_case_1();
    return 0;
}

