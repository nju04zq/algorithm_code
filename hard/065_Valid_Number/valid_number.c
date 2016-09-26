#include <stddef.h>
#include <string.h>
#include <stdlib.h>
#include "../../common/bool.h"


#define IS_NUMERIC(a) ((a) >= '0' && (a) <= '9')

static char *
strip_heading_spaces (char *s)
{
    while (*s == ' ') {
        s++;
    }
    return s;
}

static void
strip_trailing_spaces (char *s)
{
    int i, len;

    len = strlen(s);
    for (i = len-1; i >= 0; i--) {
        if (s[i] == ' ') {
            s[i] = '\0';
        } else {
            break;
        }
    }
    return;
}

static bool
is_prev_allowed_before_dot (char prev)
{
    char ch = prev;

    if (ch == '\0' || ch == '+' || ch == '-' || IS_NUMERIC(ch)) {
        return TRUE;
    } else {
        return FALSE;
    }
}

static bool
is_prev_allowed_before_e (char prev)
{
    char ch = prev;

    if (IS_NUMERIC(ch) || ch == '.') {
        return TRUE;
    } else {
        return FALSE;
    }
}

static bool
is_prev_allowed_before_sign (char prev)
{
    char ch = prev;

    if (ch == '\0' || ch == 'e') {
        return TRUE;
    } else {
        return FALSE;
    }
}

static bool
is_prev_allowed_before_end (char prev)
{
    char ch = prev;

    if (IS_NUMERIC(ch) || ch == '.') {
        return TRUE;
    } else {
        return FALSE;
    }
}

static bool
is_number_internal (char *s)
{
    int i, len;
    char ch, prev = '\0';
    bool has_sign = FALSE, has_e = FALSE, has_dot = FALSE, has_num = FALSE;

    len = strlen(s);
    if (len <= 1) {
        if (IS_NUMERIC(s[0])) {
            return TRUE;
        } else {
            return FALSE;
        }
    }

    for (i = 0; i <= len; i++) {
        ch = s[i];
        if (ch == '.') {
            if (is_prev_allowed_before_dot(prev) == FALSE) {
                return FALSE;
            }
            if (has_dot || has_e) {
                return FALSE;
            }
            has_dot = TRUE;
        } else if (ch == 'e') {
            if (is_prev_allowed_before_e(prev) == FALSE) {
                return FALSE;
            }
            if (!has_num || has_e) {
                return FALSE;
            }
            has_e = TRUE;
            has_sign = FALSE;
        } else if (ch == '+' || ch == '-') {
            if (is_prev_allowed_before_sign(prev) == FALSE) {
                return FALSE;
            }
            if (has_sign) {
                return FALSE;
            }
            has_sign = TRUE;
        } else if (ch == '\0') {
            if (is_prev_allowed_before_end(prev) == FALSE) {
                return FALSE;
            }
            if (has_num == FALSE) {
                return FALSE;
            }
        } else if (IS_NUMERIC(ch)) {
            has_num = TRUE;
        } else {
            return FALSE;
        }
        prev = ch;
    }

    return TRUE;
}

bool
isNumber (char *s)
{
    bool result;
    char *p, *p_start;
    int len;

    if (s == NULL) {
        return FALSE;
    }

    len = strlen(s);
    p = calloc(len+1, sizeof(char));
    if (p == NULL) {
        return FALSE;
    }
    strncpy(p, s, len);

    p_start = strip_heading_spaces(p);
    strip_trailing_spaces(p_start);

    result = is_number_internal(p_start);

    free(p);
    return result;
}

#include <stdio.h>

static void
test_num (char *s, bool answer)
{
    bool result;

    result = isNumber(s);
    if (result != answer) {
        printf("%s, get %d, should be %d\n", s, result, answer);
    }
    return;
}

int main (void)
{
    test_num("0", TRUE);
    test_num("0.1", TRUE);
    test_num(".", FALSE);
    test_num("abc", FALSE);
    test_num("1 a", FALSE);
    test_num("2e10", TRUE);
    test_num("+", FALSE);
    test_num("+1", TRUE);
    test_num("+1.0", TRUE);
    test_num("++1.0", FALSE);
    test_num("+1+0", FALSE);
    test_num("1..0", FALSE);
    test_num("1.0.", FALSE);
    test_num("1.0e", FALSE);
    test_num("+1e+10", TRUE);
    test_num("+1.2e-10", TRUE);
    test_num("-22.1e+11.02", FALSE);
    test_num("+1.2e-10e1", FALSE);
    test_num(" 12345689 ", TRUE);
    test_num("1.", TRUE);
    test_num(".1", TRUE);
    test_num("-.1", TRUE);
    test_num("1.e1", TRUE);
    test_num(".e1", FALSE);
    test_num("+.", FALSE);
    test_num("", FALSE);
    return 0;
}

