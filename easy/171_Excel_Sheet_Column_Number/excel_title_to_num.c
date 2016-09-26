#include <stddef.h>
#include <string.h>

#define EXCEL_COL_RADIX 26

int
titleToNumber (char *s)
{
    int i, len, num, offset;

    if (s == NULL) {
        return 0;
    }

    len = strlen(s);

    for (i = 0, num = 0; i < len; i++) {
        offset = s[i] - 'A' + 1;
        num = num * EXCEL_COL_RADIX + offset;
    }
    return num;
}

#include <stdio.h>

static void
test_title_to_num (char *s, int answer)
{
    int num;

    num = titleToNumber(s);
    if (num != answer) {
        printf("%s, get %d, should be %d\n", s, num, answer);
    }
    return;
}

int main (void)
{
    test_title_to_num("A", 1);
    test_title_to_num("B", 2);
    test_title_to_num("C", 3);
    test_title_to_num("Z", 26);
    test_title_to_num("AA", 27);
    test_title_to_num("AB", 28);
    test_title_to_num("ZA", 677);
    test_title_to_num("ZB", 678);
    test_title_to_num("AAA", 703);
    return 0;
}

