/* Run time on leetcode 60ms */
#include <stddef.h>
#include <string.h>
#include "../../common/bool.h"

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
is_match_internal (char *s, int i, char *p, int j)
{
    int len_s, len_p;
    bool result;

    len_s = strlen(s);
    len_p = strlen(p);
    if (j >= len_p) {
        return (i >= len_s);
    }

    if ((j+1) >= len_p || p[j+1] != '*') {
        if (i == len_s) {
            return FALSE;
        }
        if (is_char_same(s[i], p[j]) == FALSE) {
            return FALSE;
        }
        return is_match_internal(s, i+1, p, j+1);
    }

    result = is_match_internal(s, i, p, j+2);
    if (result == TRUE) {
        return TRUE;
    }

    for (; i < len_s; i++) {
        if (is_char_same(s[i], p[j]) == FALSE) {
            break;
        }
        result = is_match_internal(s, i+1, p, j+2);
        if (result) {
            return TRUE;
        }
    }

    return FALSE;
}

bool
isMatch (char *s, char *p)
{
    bool result;

    if (s == NULL || p == NULL) {
        return FALSE;
    }

    result = is_match_internal(s, 0, p, 0);
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
    test_is_match(".*", "", TRUE);
    return 0;
}

