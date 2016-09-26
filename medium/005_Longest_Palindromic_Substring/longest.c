#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include "../../common/bool.h"

static char *
check_palindrome_even (char *s, int mid, int len, int *max_len)
{
    int left, right;
    char *p = NULL;

    *max_len = 0;

    for (left=mid, right=mid+1; left>=0 && right<len; left--, right++) {
        if (s[left] == s[right]) {
            (*max_len) += 2;
            p = &s[left];
        } else {
            break;
        }
    }
    return p;
}

static char *
check_palindrome_odd (char *s, int mid, int len, int *max_len)
{
    int left, right;
    char *p;

    p = &s[mid];
    *max_len = 1;

    for (left= mid-1, right=mid+1; left>=0 && right<len; left--, right++) {
        if (s[left] == s[right]) {
            (*max_len) += 2;
            p = &s[left];
        } else {
            break;
        }
    }
    return p;
}

char *
longestPalindrome (char *s)
{
    int i, len, max_len = 0, max_len_1;
    char *p, *p1, *dst;

    if (s == NULL) {
        return NULL;
    }

    len = strlen(s);
    for (i = 0; i < len; i++) {
        p1 = check_palindrome_even(s, i, len, &max_len_1);
        if (max_len_1 > max_len) {
            max_len = max_len_1;
            p = p1;
        }
        p1 = check_palindrome_odd(s, i, len, &max_len_1);
        if (max_len_1 > max_len) {
            max_len = max_len_1;
            p = p1;
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

