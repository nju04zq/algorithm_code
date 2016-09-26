#include <stddef.h>
#include <string.h>

int
strStr (char *haystack, char *needle)
{
    int len_haystack, len_needle, i, j;

    if (haystack == NULL || needle == NULL) {
       return -1;
    }

    len_haystack = strlen(haystack);
    len_needle = strlen(needle);

    if (len_needle == 0) {
        return 0;
    }

    for (i = 0; i < len_haystack; i++) {
       for (j = 0; j < len_needle; j++) {
           if (haystack[i+j] != needle[j]) {
               break;
           }
       }
       if (j == len_needle) {
           return i;
       }
    }

    return -1;
}

#include <stdio.h>

static void
test_strstr (char *s, char *p)
{
    int rc;

    printf("Haystack %s\n", s);
    printf("Needle %s\n", p);

    rc = strStr(s, p);
    printf("result: %d\n", rc);
    return;
}

int main (void)
{
    test_strstr("", "");
    test_strstr("abc", "xy");
    test_strstr("abac", "ba");
    test_strstr("abac", "bac");
    test_strstr("abac", "abac");
    return 0;
}

