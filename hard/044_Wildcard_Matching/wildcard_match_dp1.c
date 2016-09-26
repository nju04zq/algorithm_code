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

static bool
is_match_internal (char *s, char *p)
{
    int len_s, len_p, i, j;
    bool *buf, *temp, *ri, *ri1, is_match;

    len_s = strlen(s);
    len_p = strlen(p);

    buf = calloc(2*(len_p+1), sizeof(bool));
    if (buf == NULL) {
        return FALSE;
    }
    ri = &buf[0];
    ri1 = &buf[len_p+1];

    ri[0] = TRUE;
    for (i = 0; i < len_p; i++) {
        if (p[i] == '*' && ri[i]) {
            ri[i+1] = TRUE;
        }
    }

    for (i = 0; i < len_s; i++) {
        memset(ri1, 0, sizeof(bool) * (len_p+1));
        for (j = 0; j < len_p; j++) {
            if (p[j] != '*') {
                if (is_same_char(s[i], p[j])) {
                    ri1[j+1] = ri[j];
                }
            } else if (ri[j] || ri[j+1] || ri1[j]) {
                ri1[j+1] = TRUE;
            }
        }
        temp = ri;
        ri = ri1;
        ri1 = temp;
    }

    is_match = ri[len_p];
    free(buf);
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

