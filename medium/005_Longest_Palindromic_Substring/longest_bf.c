#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

static bool
is_palindrome (char *s, int len)
{
    int i, j;

    for (i = 0, j = len-1; i < j; i++, j--) {
        if (s[i] != s[j]) {
            return FALSE;
        }
    }
    return TRUE;
}

char *
longestPalindrome (char *s)
{
    int len, sub_len, start, max_len = 1;
    char *p = s, *dst;

    if (s == NULL) {
        return NULL;
    }

    if (*s == '\0') {
        max_len = 0;
    }

    len = strlen(s);
    for (sub_len = 2; sub_len <= len; sub_len++) {
        start = 0;
        for (start = 0; start+sub_len <= len; start++) {
            if (is_palindrome(&s[start], sub_len) == FALSE) {
                continue;
            }
            if (sub_len > max_len) {
                max_len = sub_len;
                p = &s[start];
            }
        }
    }

    dst = calloc(max_len+1, sizeof(char));
    if (dst == NULL) {
        return NULL;
    }

    strncpy(dst, p, max_len);
    return dst;
}

#include <stdio.h>

static void
test_longest (char *s)
{
    char *p;

    printf("Longest for \"%s\"\n", s);

    p = longestPalindrome(s);
    if (p == NULL) {
        return;
    }

    printf("It's \"%s\"\n", p);
    free(p);
    return;
}

int main (void)
{
    test_longest("a");
    test_longest("abcxycba");
    test_longest("abbxyaba");
    test_longest("abcdcba");
    return 0;
}

