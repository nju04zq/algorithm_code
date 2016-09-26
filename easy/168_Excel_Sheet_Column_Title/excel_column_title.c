/*
 * f(n) = 26 * f(n-1) + a(n)
 * 1 <= a(n) <= 26
 * A = 1, Z = 26
 * AA = 1*26 + 1
 * BZ = 2*26 + 26
 * ZZ = 26*26 + 26
 */

#include <stddef.h>
#include <stdlib.h>

#define EXCEL_COL_RADIX 26

static int
calc_title_len (int n)
{
    int len = 0, offset;

    while (n > 0) {
        offset = n % EXCEL_COL_RADIX;
        if (offset == 0) {
            offset = EXCEL_COL_RADIX;
        }
        len++;
        n = (n - offset)/EXCEL_COL_RADIX;
    }
    return len;
}

static char *
convert_to_title_internal (int n)
{
    int len, i, offset;
    char *title;

    len = calc_title_len(n);
    title = calloc(len+1, sizeof(char));
    if (title == NULL) {
        return NULL;
    }

    i = len-1;
    while (n > 0) {
        offset = n % EXCEL_COL_RADIX;
        if (offset == 0) {
            offset = EXCEL_COL_RADIX;
        }
        title[i--] = 'A' + offset - 1;
        n = (n - offset)/EXCEL_COL_RADIX;
    }

    return title;
}

char *
convertToTitle (int n)
{
    char *title;

    if (n <= 0) {
        return NULL;
    }

    title = convert_to_title_internal(n);
    return title;
}

#include <stdio.h>
#include <string.h>

static void
test_excel_col_to_title (int n, char *answer)
{
    char *title;

    title = convertToTitle(n);
    if (title == NULL) {
        printf("Fail on convert.\n");
        return;
    }

    if (strcmp(title, answer) != 0) {
        printf("%d, get %s, should be %s\n", n, title, answer);
    }

    free(title);
    return;
}

int main (void)
{
    test_excel_col_to_title(703, "AAA");
    test_excel_col_to_title(677, "ZA");
    test_excel_col_to_title(1, "A");
    test_excel_col_to_title(2, "B");
    test_excel_col_to_title(3, "C");
    test_excel_col_to_title(26, "Z");
    test_excel_col_to_title(27, "AA");
    test_excel_col_to_title(28, "AB");
    test_excel_col_to_title(677, "ZA");
    test_excel_col_to_title(678, "ZB");
    return 0;
}

