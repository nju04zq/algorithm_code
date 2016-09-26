#include <stddef.h>
#include <stdlib.h>
#include <string.h>

#define MIN(a, b) ((a) < (b) ? (a) : (b))

static int
min_distance_internal (char *s, int len_s, char *t, int len_t)
{
    int *dp, result, i, j, prev, temp;

    dp = calloc(len_t+1, sizeof(int));
    if (dp == NULL) {
        return -1;
    }

    for (i = 0; i <= len_t; i++) {
        dp[i] = i;
    }

    for (i = 1; i <= len_s; i++) {
        prev = dp[0];
        dp[0] += 1;
        for (j = 1; j <= len_t; j++) {
            temp = dp[j];
            if (s[i-1] == t[j-1]) {
                dp[j] = prev;
            } else {
                dp[j] = MIN(dp[j-1]+1, dp[j]+1);
                dp[j] = MIN(prev+1, dp[j]);
            }
            prev = temp;
        }
    } 

    result = dp[len_t];
    free(dp);
    return result;
}

int
minDistance (char *s, char *t)
{
    int len_s, len_t, result;

    if (s == NULL || t == NULL) {
        return 0;
    }

    len_s = strlen(s);
    len_t = strlen(t);
    if (len_s >= len_t) {
        result = min_distance_internal(s, len_s, t, len_t);
    } else {
        result = min_distance_internal(t, len_t, s, len_s);
    }
    return result;
}

#include <stdio.h>

static void
test_min_distance (char *s, char *t)
{
    int result;

    result = minDistance(s, t);
    printf("%s, %s, min distance %d\n", s, t, result);
    return;
}

int main (void)
{
    test_min_distance("", "");
    test_min_distance("", "a");
    test_min_distance("xxxab", "abyyy");
    return 0;
}

