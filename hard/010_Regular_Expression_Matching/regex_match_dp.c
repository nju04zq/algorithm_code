/* Run time on leetcode 4ms */
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

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
    free(a);
    return;
}

static void *
calloc_2d_array (int m, int n, int element_size)
{
    int i;
    void **a;

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
is_char_same (char s, char p)
{
    if (s == p) {
        return TRUE;
    }
    if (p == '.') {
        return TRUE;
    }
    return FALSE;
}

static bool
is_match_internal (char *s, char *p)
{
    int i, j, len_p, len_s;
    bool **result, is_match;

    len_s = strlen(s);
    len_p = strlen(p);

    result = calloc_2d_array(len_s+1, len_p+1, sizeof(bool));
    if (result == NULL) {
        return FALSE;
    }

    result[0][0] = TRUE;
    for (i = 1; i < len_p; i++) {
        if (p[i] == '*') {
            result[0][i+1] = result[0][i] || result[0][i-1];
        }
    }

    for (i = 0; i < len_s; i++) {
        for (j = 0; j < len_p; j++) {
            if (p[j] != '*') {
                if (is_char_same(s[i], p[j])) {
                    result[i+1][j+1] = result[i][j];
                }
                continue;
            }
            if (j == 0) { // need to check p[j-1]/[j-2]
                continue;
            }
            if (result[i+1][j]) { //x*, x occur once
                result[i+1][j+1] = TRUE;
                continue;
            }
            //x*, x occur >= 2
            if (is_char_same(s[i], p[j-1]) && result[i][j+1]) {
                result[i+1][j+1] = TRUE;
                continue;
            }
            if (j <= 1) {
                continue;
            }
            //x*, x not occur
            if (result[i+1][j-1]) {
                result[i+1][j+1] = TRUE;
            }
        }
    }

    is_match =  result[len_s][len_p];
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
    test_is_match("a.*bb", "abbbb", TRUE);
    test_is_match("ab*bb", "abbbb", TRUE);
    test_is_match("ab*bb", "abbbc", FALSE);
    test_is_match("ab*", "a", TRUE);
    test_is_match(".*a*a", "bbbba", TRUE);
    test_is_match(".*a*aa*.*b*.c*.*a*", "aabcbcbcaccbcaabc", TRUE);
    test_is_match("c*.*a*", "c", TRUE);
    test_is_match(".*", "c", TRUE);
    test_is_match(".*", "", TRUE);
    return 0;
}

