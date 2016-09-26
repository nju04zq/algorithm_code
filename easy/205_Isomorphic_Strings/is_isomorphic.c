#include <stddef.h>
#include <string.h>
#include "../../common/bool.h"

bool
isIsomorphic (char *s, char *t)
{
    char mapping[256], mapped[256];
    int s_idx, t_idx;

    if (s == NULL || t == NULL) {
        return FALSE;
    }
    if (strlen(s) != strlen(t)) {
        return FALSE;
    }
    //TODO validate string, whether it's alphabetic
    
    memset(mapping, 0, 256*sizeof(char));
    memset(mapped, 0, 256*sizeof(char));

    for (; *s && *t; s++, t++) {
        s_idx = (int)*s;
        t_idx = (int)*t;
        if (mapping[s_idx] == '\0') {
            if (mapped[t_idx] != '\0') {
                return FALSE;
            }
            mapping[s_idx] = *t;
            mapped[t_idx] = *s;
        } else if (mapping[s_idx] != *t) {
            return FALSE;
        }
    }

    return TRUE;
}

#include <stdio.h>

static void
test_is_isomorphic (char *s, char *t, bool answer)
{
    bool result;

    result = isIsomorphic(s, t);
    if (result != answer) {
        printf("\"%s\", \"%s\", get %d, should be %d\n", s, t, result, answer);
    }
    return;
}

int main (void)
{
    test_is_isomorphic("add", "egg", TRUE);
    test_is_isomorphic("adg", "egg", FALSE);
    test_is_isomorphic("foo", "bar", FALSE);
    test_is_isomorphic("paper", "title", TRUE);
    return 0;
}

