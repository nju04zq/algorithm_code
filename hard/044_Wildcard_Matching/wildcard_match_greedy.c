#include <stddef.h>
#include <string.h>
#include "../../common/bool.h"

static bool
is_char_same (char p, char s)
{
    if (p == '?') {
        return TRUE;
    }
    return (p == s);
}

static bool
is_match_internal (char *s, char *p)
{
    int i, j, len_s, len_p;
    int star = -1, mark = -1;

    len_s = strlen(s);
    len_p = strlen(p);

    for (i = 0, j = 0; i < len_s; ) {
        if (j < len_p && p[j] == '*') {
            star = j;
            j++;
            mark = i;
        } else if (j < len_p && is_char_same(p[j], s[i])) {
            i++;
            j++;
        } else if (star != -1) {
            j = star + 1;
            i = ++mark;
        } else {
            return FALSE;
        }
    }

    while (j < len_p && p[j] == '*') {
        j++;
    }

    return (j == len_p);
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

