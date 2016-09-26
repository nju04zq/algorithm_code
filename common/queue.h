#ifndef __QUEUE_H__
#define __QUEUE_H__

#include "bool.h"

typedef struct uqueue_s {
    void **buf;
    int head;
    int tail;
    int iter_index;
    int size;
    int max_size;
} uqueue_t;

bool
is_queue_full(uqueue_t *queue);

bool
is_queue_empty(uqueue_t *queue);

int
get_queue_size(uqueue_t *queue);

void *
dequeue(uqueue_t *queue);

void *
peek_queue_head(uqueue_t *queue);

void *
get_queue_next(uqueue_t *queue);

int
enqueue(uqueue_t *queue, void *udata);

void
clean_queue(uqueue_t *queue);

int
init_queue(uqueue_t *queue);

void
dump_queue(uqueue_t *queue, void (*print_udata)(void *));
#endif //__QUEUE_H__ 
