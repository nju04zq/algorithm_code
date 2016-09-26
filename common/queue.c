#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "bool.h"
#include "queue.h"

#ifdef __UQUEUE_TEST_DEBUG__
#define UQUEUE_INIT_SIZE 4 
#else
#define UQUEUE_INIT_SIZE 64
#endif

bool
is_queue_full (uqueue_t *queue)
{
    if (queue->size == queue->max_size) {
        return TRUE;
    } else {
        return FALSE;
    }
}

bool
is_queue_empty (uqueue_t *queue)
{
    if (queue->size == 0) {
        return TRUE;
    } else {
        return FALSE;
    }
}

int
get_queue_size (uqueue_t *queue)
{
    return queue->size;
}

static void
reassign_queue_buf (void **buf, uqueue_t *queue)
{
    int len;

    if (queue->tail > queue->head) {
        memcpy(buf, &queue->buf[queue->head], queue->size * sizeof(void*));
    } else {
        len = queue->max_size - queue->head;
        memcpy(buf, &queue->buf[queue->head], len * sizeof(void*));
        memcpy(&buf[len], &queue->buf[0], queue->tail * sizeof(void*));
    }

    queue->head = 0;
    queue->tail = queue->size;
    free(queue->buf);
    queue->buf = buf;
    return;
}

static int
resize_queue (uqueue_t *queue, int new_max_size)
{
    void **buf;

    if (queue->size > new_max_size) {
        return -1;
    }

    buf = calloc(new_max_size, sizeof(void *));
    if (buf == NULL) {
        return -1;
    }

    reassign_queue_buf(buf, queue);
    queue->max_size = new_max_size;

    return 0;
}

static void
shrink_queue (uqueue_t *queue)
{
    int new_max_size;

    if (queue->max_size <= UQUEUE_INIT_SIZE) {
        return;
    }
    if (queue->size > (queue->max_size/4)) {
        return;
    }

    new_max_size = queue->max_size/2;
    (void)resize_queue(queue, new_max_size);
    return;
}

static int
expand_queue (uqueue_t *queue)
{
    int new_max_size, rc;

    new_max_size = queue->max_size * 2;
    rc = resize_queue(queue, new_max_size);
    return rc;
}

void *
dequeue (uqueue_t *queue)
{
    void *udata;

    if (is_queue_empty(queue)) {
        return NULL;
    }

    udata = queue->buf[queue->head];

    queue->head++;
    if (queue->head == queue->max_size) {
        queue->head = 0;
    }
    queue->size--;

    shrink_queue(queue);
    return udata;
}

static void *
get_queue_entry (uqueue_t *queue)
{
    void *udata;
    int i;

    i = queue->iter_index + queue->head;
    if (i >= queue->max_size) {
        i -= queue->max_size;
    }

    udata = queue->buf[i];
    return udata;
}

void *
peek_queue_head (uqueue_t *queue)
{
    void *udata;

    queue->iter_index = 0;

    if (is_queue_empty(queue)) {
        return NULL;
    }

    udata = get_queue_entry(queue);
    return udata;
}

void *
get_queue_next (uqueue_t *queue)
{
    void *udata;

    if (queue->iter_index >= (queue->size - 1)) {
        return NULL;
    }

    queue->iter_index++;
    udata = get_queue_entry(queue);
    return udata;
}

int
enqueue (uqueue_t *queue, void *udata)
{
    int rc;

    if (is_queue_full(queue)) {
        rc = expand_queue(queue);
        if (rc != 0) {
            return -1;
        }
    }

    queue->buf[queue->tail] = udata;

    queue->tail++;
    if (queue->tail == queue->max_size) {
        queue->tail = 0;
    }

    queue->size++;
    return 0;
}

void
clean_queue (uqueue_t *queue)
{
    free(queue->buf);
    return;
}

int
init_queue (uqueue_t *queue)
{
    memset(queue, 0, sizeof(uqueue_t));

    queue->buf = calloc(UQUEUE_INIT_SIZE, sizeof(void *));
    if (queue->buf == NULL) {
        return -1;
    }

    queue->head = 0;
    queue->tail = 0;
    queue->size = 0;
    queue->max_size = UQUEUE_INIT_SIZE;
    return 0;
}

void
dump_queue (uqueue_t *queue, void (*print_udata)(void *))
{
    int i, size;
    void *udata;
    bool first_entry = TRUE;

    if (is_queue_empty(queue)) {
        printf("<Empty queue>\n");
        return;
    }

    size = get_queue_size(queue);
    for (i = 0; i < size; i++) {
        if (i == 0) {
            udata = peek_queue_head(queue);
        } else {
            udata = get_queue_next(queue);
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

