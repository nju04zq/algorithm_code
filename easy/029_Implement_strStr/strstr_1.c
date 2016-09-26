#include <stdlib.h>
#include <stddef.h>
#include <string.h>

static void
fail (int *failure, char *p)
{
    int i, j, len;

    len = strlen(p);

    failure[0] = -1;
    for (i = 1; i < len; i++) {
        j = failure[i-1];
        while (j > 0 && p[i] != p[j+1]) {
            j = failure[j];
        }
        if (p[i] == p[j+1]) {
            failure[i] = j+1;
        } else {
            failure[i] = -1;
        }
    }
    return;
}

int
strStr (char *s, char *p)
{
    int *failure, i, j, len_s, len_p, rc;

    if (s == NULL || p == NULL) {
        return -1;
    }

    len_s = strlen(s);
    len_p = strlen(p);

    if (len_p == 0) {
        return 0;
    }

    failure = calloc(len_p, sizeof(int));
    if (failure == NULL) {
        return -1;
    }

    fail(failure, p);
    j = failure[0];
    rc = -1;
    for (i = 0; i < len_s; i++) {
        while (j > 0 && s[i] != p[j+1]) {
            j = failure[j];
        }
        if (s[i] == p[j+1]) {
            j++;
        }
        if (j == (len_p- 1)) {
            rc = i - len_p + 1;
            break;
        }
    }

    free(failure);
    return rc;
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

