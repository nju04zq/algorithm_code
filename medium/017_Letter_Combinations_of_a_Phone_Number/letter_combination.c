#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/queue.h"

static char *maps[] = 
{
    "",   //0
    "",   //1
    "abc",//2
    "def",//3
    "ghi",//4
    "jkl",//5
    "mno",//6
    "pqrs",//7
    "tuv",//8
    "wxyz",//9
};

static char *
get_digit_map (int i)
{
    if (i >= 0 && i <= 9) {
        return maps[i];
    }
    return "";
}

static int
add_result (uqueue_t *queue, int len, char *s)
{
    char *p;
    int rc;

    p = calloc(len+1, sizeof(char));
    if (p == NULL) {
        return 0;
    }

    strncpy(p, s, len);

    rc = enqueue(queue, p);
    if (rc != 0) {
        free(p);
        return rc;
    }

    return 0;
}

static int
letter_combinations (uqueue_t *queue, char *s, int len,
                     char *mask, int depth)
{
    int i, k, rc;
    char *map;

    if (depth == len) {
        rc = add_result(queue, len, mask);
        return rc;
    }

    k = s[depth] - '0';
    map = get_digit_map(k);
    for (i = 0; map[i]; i++) {
        mask[depth] = map[i];
        rc = letter_combinations(queue, s, len, mask, depth+1);
        if (rc != 0) {
            return rc;
        }
    }
    mask[depth] = '\0';
    return 0;
}

static void
clean_out_queue (uqueue_t *queue)
{
    char *s;
    int i, size;

    size = get_queue_size(queue);
    for (i = 0; i < size; i++) {
        s = dequeue(queue);
        free(s);
    }
    return;
}

static char **
copy_out_queue (uqueue_t *queue, int *returnSize)
{
    char **result;
    int i, size;

    size = get_queue_size(queue);
    result = calloc(size, sizeof(char *));
    if (result == NULL) {
        return NULL;
    }

    *returnSize = size;
    for (i = 0; i < size; i++) {
        result[i] = dequeue(queue);
    }
    return result;
}

char **
letterCombinations (char *digits, int *returnSize)
{
    uqueue_t queue;
    int rc, len;
    char *mask, **result = NULL;

    *returnSize = 0;
    if (digits == NULL || digits[0] == '\0') {
        return NULL;
    }

    rc = init_queue(&queue);
    if (rc != 0) {
        return NULL;
    }

    len = strlen(digits);
    mask = calloc(len+1, sizeof(char));
    if (mask == NULL) {
        clean_queue(&queue);
        return NULL;
    }

    rc = letter_combinations(&queue, digits, len, mask, 0);
    if (rc == 0) {
        result = copy_out_queue(&queue, returnSize);
    }
    if (result == NULL) {
        clean_out_queue(&queue);
    }

    free(mask);
    clean_queue(&queue);
    return result;
}

#include <stdio.h>

static void
test_letter_combination (char *s)
{
    char **result;
    int i, size;

    printf("Combintations for %s\n", s);
    result = letterCombinations(s, &size);
    if (result == NULL) {
        printf("<NULL>\n");
        return;
    }

    printf("result total %d\n", size);
    for (i = 0; i < size; i++) {
        printf("%s\n", result[i]);
        free(result[i]);
        //printf("%s\n", result[i]);
    }
    free(result);
    return;
}

int main (void)
{
    test_letter_combination("23");
    test_letter_combination("29323");
    return 0;
}

