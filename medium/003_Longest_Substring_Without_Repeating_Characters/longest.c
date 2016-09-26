#include <string.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

int
lengthOfLongestSubstring (char *s)
{
    int marks[256], max_len, cur_len, start, idx, i;
    char *s1 = s;

    start = 0;
    max_len = 0;
    memset(marks, -1, sizeof(marks));

    for (i = 0; *s; s++, i++) {
        idx = (int)(*s);
        if (marks[idx] != -1 && marks[idx] >= start) {
            cur_len = i - start;
            max_len = MAX(max_len, cur_len);
            start = marks[idx] + 1;
            //printf("max_len %d, start %02d, i %02d, %20s\n",
            //       max_len, start, i, &s1[start]);
        }
        marks[idx] = i;
    }

    cur_len = i - start;
    max_len = MAX(max_len, cur_len);
    return max_len;
}

#include <stdio.h>

static void
test_longest (char *s, int answer)
{
    int max_len;

    max_len = lengthOfLongestSubstring(s);
    if (max_len != answer) {
        printf("Longest substring for \"%s\", get %d, should be %d\n",
               s, max_len, answer);
    }
    return;
}

int main (void)
{
    //test_longest("abcd", 4);
    //test_longest("aaaa", 1);
    test_longest("aababcaeabcdabcaba", 5);
    return 0;
}

