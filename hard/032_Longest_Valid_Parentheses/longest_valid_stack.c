#include <stddef.h>
#include <string.h>
#include "../../common/stack.h"

#define MAX(a, b) ((a) > (b) ? (a) : (b))

static int
longest_pairs (ustack_t *stack, char *s)
{
    int rc, cur_start, i, j, len, max_pairs, size;

    len = strlen(s);
    max_pairs = 0;
    cur_start = 0;
    for (i = 0; i < len; i++) {
        if (s[i] == '(') {
            rc = push(stack, (void *)(unsigned long)i);
            if (rc != 0) {
                return -1;
            }
            continue;
        }
        if (is_stack_empty(stack)) {
            max_pairs = MAX(max_pairs, (i-cur_start)/2);
            cur_start = i+1;
        } else {
            (void)pop(stack);
        }
    }

    size = get_stack_size(stack);
    for (; size > 0; size--) {
        j = (int)(unsigned long)pop(stack);
        max_pairs = MAX(max_pairs, (i-j)/2);
        i = j;
    }
    max_pairs = MAX(max_pairs, (i-cur_start)/2);
    return max_pairs;
}

int
longestValidParentheses (char *s)
{
    ustack_t stack;
    int rc, max_pairs;

    rc = init_stack(&stack);
    if (rc != 0) {
        return rc;
    }

    max_pairs = longest_pairs(&stack, s);

    clean_stack(&stack);
    return 2*max_pairs;
}

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

static bool
are_pairs_valid (char *s, int n)
{
    int i, balance = 0;

    for (i = 0; i < n; i++) {
        if (s[i] == '(') {
            balance++;
            continue;
        }
        balance--;
        if (balance < 0) {
            return FALSE;
        }
    }
    if (balance != 0) {
        return FALSE;
    } else {
        return TRUE;
    }
}

static int
longest_pairs_bf (char *s)
{
    int max_pairs, i, size, n;

    n = strlen(s);
    max_pairs = 0;
    for (size = 2; size <= n; size+=2) {
        for (i = 0; i <= n-size; i++) {
            if (are_pairs_valid(&s[i], size)) {
                max_pairs = size/2;
                break;
            }
        }
    }

    return 2*max_pairs;
}

#define TEST_STR_LEN_RANGE 6
#define TEST_STR_LEN_BASE 10

static int
generate_random_len (void)
{
    int len;

    len = random() % TEST_STR_LEN_RANGE;
    len += TEST_STR_LEN_BASE;
    return len;
}

static char *
generate_random_str (void)
{
    char *s;
    int i, len, flag;

    len = generate_random_len();
    s = calloc(len+1, sizeof(char));
    if (s == NULL) {
        return NULL;
    }

    for (i = 0; i < len; i++) {
        flag = random() % 2;
        s[i] = flag == 0 ? '(' : ')';
    }

    return s;
}

static int
test_longest_pairs (char *s)
{
    int max_pairs, answer;

    max_pairs = longestValidParentheses(s);
    answer = longest_pairs_bf(s);
    if (max_pairs != answer) {
        printf("%s, get %d, should be %d\n", s, max_pairs, answer);
        return -1;
    }
    return 0;
}

static int
run_random_test (void)
{
    char *s;
    int rc;

    s = generate_random_str();
    if (s == NULL) {
        return -1;
    }

    rc = test_longest_pairs(s);
    free(s);
    return rc;
}

#define TEST_CASE_CNT 1000

int main (void)
{
    int i, rc;

    test_longest_pairs("(()");
    test_longest_pairs(")()())");
    for (i = 0; i < TEST_CASE_CNT; i++) {
        rc = run_random_test();
        if (rc != 0) {
            printf("Fail on case ##%d##\n", i);
            break;
        }
    }
    return 0;
}

