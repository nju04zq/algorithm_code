#include <stdint.h>
#include <stddef.h>
#include <limits.h>

typedef unsigned char bool;

#define TRUE 1
#define FALSE 0

#define IS_NUM(a) ((a) >= '0' && (a) <= '9')

static int
convert_to_int (char *str, bool is_negative)
{
    int64_t x = 0;

    for (; IS_NUM(*str); str++) {
        x = x*10 + (*str - '0');
        if (x >= UINT32_MAX) {
            break;
        }
    }

    if (is_negative) {
        x = -x;
    }

    if (x > INT_MAX) {
        x = INT_MAX;
    } else if (x < INT_MIN) {
        x = INT_MIN;
    }

    return (int)x;
}

static bool
check_negative (char *str)
{
    if (*str == '-') {
        return TRUE;
    }
    return FALSE;
}

static bool
is_sign (char *str)
{
    if (*str == '+' || *str == '-') {
        return TRUE;
    }
    return FALSE;
}

static char *
strip_space (char *str)
{
    while (*str == ' ') {
        str++;
    }
    return str;
}

int
myAtoi (char *str)
{
    int x;
    bool is_negative = FALSE;

    if (str == NULL) {
        return 0;
    }

    str = strip_space(str);

    if (is_sign(str)) {
        if (check_negative(str)) {
            is_negative = TRUE;
        }
        str++; //skip over sign
    }

    x = convert_to_int(str, is_negative);
    return x;
}

#include <stdio.h>

static void
test_atoi (char *str)
{
    printf("%s, %d\n", str, myAtoi(str));
    return;
}

int main (void)
{
    test_atoi(NULL);
    test_atoi("");
    test_atoi("123");
    test_atoi("-123");
    test_atoi("+123");
    test_atoi("+123a");
    test_atoi("abc");
    test_atoi("   100");
    test_atoi("2147483648");
    test_atoi("-2147483648");
    test_atoi("0011");
    return 0;
}
