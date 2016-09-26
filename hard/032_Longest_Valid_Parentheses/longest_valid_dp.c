#include <stddef.h>
#include <stdlib.h>
#include <string.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

static int
longest_valid_len (int *dp, char *s, int n)
{
    int i, j, max_len;

    if (n <= 0) {
        return 0;
    }

    dp[n-1] = 0;
    for (i = n-2; i >= 0; i--) {
        if (s[i] == ')') {
            dp[i] = 0;
            continue;
        }
        j = i + dp[i+1] + 1;
        if (j >= n || s[j] != ')') {
            dp[i] = 0;
            continue;
        }
        dp[i] = dp[i+1] + 2;
        if ((j+1) < n) {
            dp[i] += dp[j+1];
        }
    }

    max_len = 0;
    for (i = 0; i < n; i++ ) {
        max_len = MAX(max_len, dp[i]);
    }
    return max_len;
}

int
longestValidParentheses (char *s)
{
    int *dp, max_len, len;

    if (s == NULL) {
        return 0;
    }

    len = strlen(s);
    dp = calloc(len, sizeof(int));
    if (dp == NULL) {
        return -1;
    }

    max_len = longest_valid_len(dp, s, len);

    free(dp);
    return max_len;
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

    test_longest_pairs("()(())");
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

