#include <stddef.h>
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

static bool
is_match_internal (char *s, int i, char *p, int j)
{
    int len_s, len_p;
    bool result;

    len_s = strlen(s);
    len_p = strlen(p);

    if (j >= len_p) {
        return i >= len_s;
    }

    if (p[j] != '*') {
        if (i >= len_s) {
            return FALSE;
        }
        if (is_same_char(s[i], p[j]) == FALSE) {
            return FALSE;
        }
        return is_match_internal(s, i+1, p, j+1);
    }

    for (; i <= len_s; i++) {
        result = is_match_internal(s, i, p, j+1);
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
    test_is_match("?", "", FALSE);
    test_is_match("*", "", TRUE);
    test_is_match("a?c*", "abc", TRUE);
    test_is_match("a?c*", "abcc", TRUE);
    test_is_match("*a?c*", "abcc", TRUE);
    test_is_match("*a?c*", "xabcc", TRUE);
    return 0;
}

