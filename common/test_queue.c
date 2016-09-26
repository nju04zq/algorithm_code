#include <stdio.h>
#include "queue.h"

static void
print_queue_entry (void *data)
{
    printf("%d", *(int *)data);
    return;
}

static void
test_case_1 (void)
{
    int a[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12};
    int i, size, rc;
    void *entry;
    uqueue_t queue;

    size = sizeof(a)/sizeof(a[0]);

    rc = init_queue(&queue);
    if (rc != 0) {
        return;
    }

    printf("After init### ");
    dump_queue(&queue, print_queue_entry);

    for (i = 0; i < size; i++) {
        enqueue(&queue, &a[i]);
        printf("After enqueue %d### ", a[i]);
        dump_queue(&queue, print_queue_entry);
    }

    for (;;) {
        entry = dequeue(&queue);
        if (!entry) {
            break;
        }
        printf("After dequeue %d### ", *(int *)entry);
        dump_queue(&queue, print_queue_entry);
    }

    printf("Before clean### ");
    dump_queue(&queue, print_queue_entry);
    clean_queue(&queue);
    return;
}

static void
test_case_2 (void)
{
    int a[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12};
    int i, size, rc;
    void *entry;
    uqueue_t queue;

    size = sizeof(a)/sizeof(a[0]);

    rc = init_queue(&queue);
    if (rc != 0) {
        return;
    }

    printf("After init### ");
    dump_queue(&queue, print_queue_entry);

    for (i = 0; i < size/2; i++) {
        enqueue(&queue, &a[i]);
        printf("After enqueue %d### ", a[i]);
        dump_queue(&queue, print_queue_entry);
    }

    for (i = 0; i < 3; i++) {
        entry = dequeue(&queue);
        if (!entry) {
            break;
        }
        printf("After dequeue %d### ", *(int *)entry);
        dump_queue(&queue, print_queue_entry);
    }

    for (i = size/2; i < size; i++) {
        enqueue(&queue, &a[i]);
        printf("After enqueue %d### ", a[i]);
        dump_queue(&queue, print_queue_entry);
    }

    for (;;) {
        entry = dequeue(&queue);
        if (!entry) {
            break;
        }
        printf("After dequeue %d### ", *(int *)entry);
        dump_queue(&queue, print_queue_entry);
    }

    printf("Before clean### ");
    dump_queue(&queue, print_queue_entry);
    clean_queue(&queue);
    return;
}

int main (void)
{
    test_case_1();
    test_case_2();
    return 0;
}

