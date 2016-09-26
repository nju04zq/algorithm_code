#include <stddef.h>

int
lengthOfLastWord (char *s)
{
    int prev_len = 0, cur_len = 0;

    if (s == NULL) {
        return 0;
    }

    for (; *s; s++) {
        if (*s != ' ') {
            cur_len++;
        } else if (cur_len != 0) {
            prev_len = cur_len;
            cur_len = 0;
        }
    }

    if (cur_len != 0) {
        return cur_len;
    } else {
        return prev_len;
    }
}

#include <stdio.h>

static void
test_len_of_last_word (char *s)
{
    int len;

    len = lengthOfLastWord(s);
    printf("\"%s\", last word len %d\n", s, len);
    return;
}

int main (void)
{
    test_len_of_last_word("");
    test_len_of_last_word("  ");
    test_len_of_last_word("123");
    test_len_of_last_word(" 123");
    test_len_of_last_word(" 123 ");
    test_len_of_last_word(" 123 1234");
    test_len_of_last_word(" 123 1234    ");
    return 0;
}
