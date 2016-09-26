#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

static bool
is_same_char (char s, char p)
{
    if (p == '?') {
        return TRUE;
    }
    return (s == p);
}

static void
free_2d_array (void **a, int m, int n)
{
    int i;

    if (a == NULL) {
        return;
    }

    for (i = 0; i < m; i++) {
        if (a[i]) {
            free(a[i]);
        }
    }
    return;
}

static void *
calloc_2d_array (int m, int n, int element_size)
{
    void **a;
    int i;

    a = calloc(m, sizeof(void *));
    if (a == NULL) {
        return NULL;
    }

    for (i = 0; i < m; i++) {
        a[i] = calloc(n, element_size);
        if (a[i] == NULL) {
            free_2d_array(a, m, n);
            return NULL;
        }
    }

    return a;
}

static bool
is_match_internal (char *s, char *p)
{
    int len_s, len_p, i, j;
    bool **result, is_match;

    len_s = strlen(s);
    len_p = strlen(p);

    result = calloc_2d_array(len_s+1, len_p+1, sizeof(bool));
    if (result == NULL) {
        return FALSE;
    }

    // pay attention to init
    result[0][0] = TRUE;
    for (i = 0; i < len_p; i++) {
        if (p[i] == '*' && result[0][i]) {
            result[0][i+1] = TRUE;
        }
    }

    // loop var i/j can be switched
    for (i = 0; i < len_s; i++) {
        for (j = 0; j < len_p; j++) {
            if (p[j] != '*') {
                if (is_same_char(s[i], p[j])) {
                    result[i+1][j+1] = result[i][j];
                }
                continue;
            }

            if (result[i][j] || result[i][j+1] || result[i+1][j]) {
                result[i+1][j+1] = TRUE;
            }
        }
    }

    is_match = result[len_s][len_p];
    free_2d_array((void **)result, len_s+1, len_p+1);
    return is_match;
}

bool
isMatch (char *s, char *p)
{
    bool result;

    if (s == NULL || p == NULL) {
        return FALSE;
    }

    result = is_match_internal(s, p);
    return result;
}

#include <stdio.h>

static void
test_is_match (char *p, char *s, bool answer)
{
    bool result;

    result = isMatch(s, p);
    if (result != answer) {
        printf("P:\"%s\", S:\"%s\", get %d, should be %d\n",
               p, s, result, answer);
    }
    return;
}

int main (void)
{
    test_is_match("ab", "abc", FALSE);
    test_is_match("abc", "abc", TRUE);
    test_is_match("?", "", FALSE);
    test_is_match("*", "", TRUE);
    test_is_match("a?c*", "abc", TRUE);
    test_is_match("a?c*", "abcc", TRUE);
    test_is_match("*a?c*", "abcc", TRUE);
    test_is_match("*a?c*", "xabcc", TRUE);
    return 0;
}

