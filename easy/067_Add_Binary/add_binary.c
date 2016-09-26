#include <stddef.h>
#include <stdlib.h>
#include <string.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

static char
get_last_digit (char *s, int i, int len)
{
    if (i > len) {
        return '0';
    } else {
        return s[len - i];
    }
}

static char
add_binary_digit (char a, char b, char *flag)
{
    if (a == '1' && b == '1') {
        *flag = '1';
        return '0';
    } else if (a == '0' && b == '0') {
        *flag = '0';
        return '0';
    } else {
        *flag = '0';
        return '1';
    }
}

static char
add_binary_digit_with_flag (char a, char b, char *flag)
{
    char flag1, flag2, flag3, c;

    c = add_binary_digit(a, b, &flag1);
    c = add_binary_digit(c, *flag, &flag2);
    *flag = add_binary_digit(flag1, flag2, &flag3);
    return c;
}

char *
addBinary (char *a, char *b)
{
    char digit_a, digit_b, digit_c, flag, *c;
    int i, len_a, len_b, len_c;

    if (a == NULL || b == NULL) {
        return 0;
    }

    //Check input binary string, TODO

    len_a = strlen(a);
    len_b = strlen(b);
    len_c = MAX(len_a, len_b) + 1; //reserver 1 for carrying flag
    c = calloc(len_c+1, sizeof(char));
    if (c == NULL) {
        return NULL;
    }

    flag = '0';
    for (i = 1; i <= len_a || i <= len_b; i++) {
        digit_a = get_last_digit(a, i, len_a);
        digit_b = get_last_digit(b, i, len_b);
        digit_c = add_binary_digit_with_flag(digit_a, digit_b, &flag);
        c[len_c-i-1] = digit_c;
    }

    if (len_a == 0 && len_b == 0) {
        c[0] = '0';
        return c;
    }

    if (flag == '1') {
        memmove(&c[1], &c[0], len_c);
        c[0] = flag;
    }

    return c;
}

#include <stdio.h>

static void
test_binary_add (char *a, char *b)
{
    char *c;

    printf("%s + %s = ", a, b);

    c = addBinary(a, b);
    if (c == NULL) {
        printf("???\n");
    } else {
        printf("%s\n", c);
        free(c);
    }
    return;
}

int main (void)
{
    test_binary_add("", "");
    test_binary_add("1", "1");
    test_binary_add("0", "0");
    test_binary_add("1", "0");
    test_binary_add("100", "1");
    test_binary_add("101", "1");
    test_binary_add("111", "1");
    return 0;
}

