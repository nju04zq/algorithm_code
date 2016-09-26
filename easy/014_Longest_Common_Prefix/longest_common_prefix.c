#include <stdlib.h>
#include <string.h>

static char *
clone (char *s, int len)
{
    char *p;

    if (len == -1) {
        len = strlen(s);
    }

    p = calloc(len+1, sizeof(char));
    if (p == NULL) {
        return p;
    }

    strncpy(p, s, len);
    return p;
}

static int
check_input (char **strs, int cnt)
{
    int i;

    if (strs == NULL) {
        return -1;
    }

    if (cnt == 0) {
        return -1;
    }

    for (i = 0; i < cnt; i++) {
        if (strs[i] == NULL) {
            return -1;
        }
    }

    return 0;
}

char *
longestCommonPrefix (char **strs, int strsSize)
{
    int rc, i, j;
    char *prefix;

    rc = check_input(strs, strsSize); 
    if (rc == -1) {
        return clone("", -1);
    }

    if (strsSize == 1) {
        return clone(strs[0], -1);
    }

    for (i = 0;; i++) {
        for (j = 1; j < strsSize; j++) {
            if (strs[j][i] != strs[0][i]) {
                break;
            }
        }
        if (j < strsSize) {
            break;
        }
    }

    prefix = clone(strs[0], i);
    return prefix;
}

#include <stdio.h>

static void
test_LCP (char **strs, int cnt)
{
    char *p;
    int i;

    p = longestCommonPrefix(strs, cnt);
    if (p == NULL) {
        printf("ERROR!\n");
        return;
    }

    printf("LCP among is \"%s\"\n", p);
    for (i = 0; i < cnt; i++) {
        printf("%s\n", strs[i]);
    }

    free(p);
    return;
}

int main (void)
{
    char *p0[] = {"abc"};
    char *p1[] = {"abce", "abcd", "abdc"};
    char *p2[] = {"abc", "", "abdc"};

    test_LCP(p0, 1);
    test_LCP(p1, 3);
    test_LCP(p2, 3);
    return 0;
}
