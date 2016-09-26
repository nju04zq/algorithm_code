/*
 * The outer loop finds a window that has all the required chars
 * The inner loop shrink the window to find the min one
 * When the inner loop breaks, it's assured that window start at l won't contain
 * a window has all those chars with len less than max_len
 */

#include <limits.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>

#define WIN_LEN(l, r) ((r) - (l) + 1)

static char *
clone_str (char *s, int len)
{
    char *t;

    t = calloc(len+1, sizeof(char));
    if (t == NULL) {
        return NULL;
    }

    strncpy(t, s, len);
    return t;
}

static void
setup_hash_map (char *s, int *hash_map)
{
    int idx;

    for (; *s; s++) {
        idx = (int)(*s);
        hash_map[idx]++;
    }
    return;
}

static char *
min_window_internal (char *s, char *t)
{
    int hash_map[256], win_map[256], t_len, s_len;
    int idx, occured = 0, l = 0, r = 0;
    int min_l = -1, min_r = -1, min_len = INT_MAX;

    memset(hash_map, 0, sizeof(hash_map));
    memset(win_map, 0, sizeof(win_map));
    setup_hash_map(t, hash_map);

    s_len = strlen(s);
    t_len = strlen(t);

    for (l = 0, r = 0; r < s_len; r++) {
        idx = (int)(s[r]);
        if (hash_map[idx] == 0) {
            continue;
        }
        win_map[idx]++;
        if (win_map[idx] <= hash_map[idx]) {
            occured++;
        }
        if (occured < t_len) {
            continue;
        }
        while (l <= r) {
            if (WIN_LEN(l, r) < min_len) {
                min_l = l;
                min_r = r;
                min_len = WIN_LEN(l, r);
            }

            idx = (int)(s[l++]);
            if (hash_map[idx] > 0) {
                win_map[idx]--;
                if (win_map[idx] < hash_map[idx]) {
                    occured--;
                    break;
                }
            }
        }
    }

    if (min_len == INT_MAX) {
        return clone_str("", 0);
    } else {
        return clone_str(&s[min_l], min_len);
    }
}

char *
minWindow (char *s, char *t)
{
    char *win;

    if (s == NULL || t == NULL || s[0] == '\0' || t[0] == '\0') {
        return clone_str("", 0);
    }

    win = min_window_internal(s, t);
    return win;
}

#include <stdio.h>

static void
test_min_window (char *s, char *t)
{
    char *win;

    win = minWindow(s, t);
    if (win == NULL) {
        printf("Fail to get min window.\n");
        return;
    }

    printf("s: %s\nt: %s\nw: %s\n", s, t, win);

    free(win);
    return;
}

int main (void)
{
    test_min_window("ADOBECODEBANC", "ABC");
    test_min_window("ABC", "ADOBECODEBANC");
    return 0;
}

