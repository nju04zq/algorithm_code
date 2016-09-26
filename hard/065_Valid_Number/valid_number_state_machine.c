#include <stddef.h>
#include <string.h>
#include <stdlib.h>
#include "../../common/bool.h"

#define IS_NUMERIC(a) ((a) >= '0' && (a) <= '9')

enum {
    CH_INVALID = -1,
    CH_NUM = 0,
    CH_SIGN,
    CH_DOT,
    CH_EXPONENT,
    CH_NULL,
    CH_MAX,
};

static int
state_transition[CH_MAX][CH_MAX] = //[prev][cur]
{
//  NUM SIGN DOT E  NULL
    {1,  0,   1, 1, 1  }, //NUM
    {1,  0,   1, 0, 0  }, //SIGN
    {1,  0,   0, 1, 1  }, //DOT
    {1,  1,   0, 0, 0  }, //E
    {1,  1,   1, 0, 0  }, //NULL
};

static int
ch_count_max_init[CH_MAX] =
// NUM, SIGN, DOT, E, NULL
{  -1,   2,    1,  1, 1};

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

static int
get_ch_type (char ch)
{
    if (ch == 'e') {
        return CH_EXPONENT;
    } else if (ch == '+') {
        return CH_SIGN;
    } else if (ch == '-') {
        return CH_SIGN;
    } else if (ch == '.') {
        return CH_DOT;
    } else if (ch == '\0') {
        return CH_NULL;
    } else if (IS_NUMERIC(ch)) {
        return CH_NUM;
    } else {
        return CH_INVALID;
    }
}

static bool
is_number_internal (char *s)
{
    int prev_type = CH_NULL, cur_type;
    int count_max, i, len;
    int ch_count[CH_MAX] = {0,}, ch_count_max[CH_MAX];

    memcpy(ch_count_max, ch_count_max_init, sizeof(ch_count_max));

    len = strlen(s);
    for (i = 0; i <= len; i++) {
        cur_type = get_ch_type(s[i]);
        if (cur_type == CH_INVALID) {
            return FALSE;
        }
        if (state_transition[prev_type][cur_type] == 0) {
            return FALSE;
        }

        ch_count[cur_type]++;
        count_max = ch_count_max[cur_type];
        if (count_max != -1 && ch_count[cur_type] > count_max) {
            return FALSE;
        }

        if (cur_type == CH_EXPONENT) {
           if(ch_count[CH_NUM] == 0) {
                return FALSE; //for .e1
           }
           ch_count_max[CH_DOT] = 0;
        }

        prev_type = cur_type;
    }

    if (ch_count[CH_NUM] == 0) {
        return FALSE;
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


