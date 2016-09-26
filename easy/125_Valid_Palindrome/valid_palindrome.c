#include <stddef.h>
#include <string.h>
#include "../../common/bool.h"

static bool
is_alphanumeric (char c)
{
    if (c >= 'a' && c <= 'z') {
        return TRUE;
    }

    if (c >= 'A' && c <= 'Z') {
        return TRUE;
    }

    if (c >= '0' && c <= '9') {
        return TRUE;
    }

    return FALSE;
}

static char
to_lower (char c)
{
    if (c >= '0' && c <= '9') {
        return c;
    }

    if (c >= 'a' && c <= 'z') {
        return c;
    }

    if (c >= 'A' && c <= 'Z') {
        return c - 'A' + 'a';
    }

    return c;
}

bool
isPalindrome (char* s)
{
    int i, j, len;
    bool the_same = TRUE;

    if (s == NULL || s[0] == '\0') {
        return TRUE;
    }

    len = strlen(s);
    for (i = 0, j = len-1; i < j;) {
        if (is_alphanumeric(s[i]) == FALSE) {
            i++;
            continue;
        }
        if (is_alphanumeric(s[j]) == FALSE) {
            j--;
            continue;
        }
        if (to_lower(s[i]) != to_lower(s[j])) {
            the_same = FALSE;
            break;
        }
        i++;
        j--;
    }

    return the_same;
}

#include <stdio.h>

static void
test_is_palindrome (char *s)
{
    bool rc;

    rc = isPalindrome(s);
    printf("\"%s\" is palindrome, %d\n", s, rc);
    return;
}

int main (void)
{
    test_is_palindrome("");
    test_is_palindrome("A man, a plan, a canal: Panama");
    test_is_palindrome("race a car");
    return 0;
}

