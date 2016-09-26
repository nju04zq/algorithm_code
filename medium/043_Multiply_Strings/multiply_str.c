#include <stddef.h>
#include <stdlib.h>
#include <string.h>

#define CHAR2INT(a) ((a)=='\0'? 0 : ((a)-'0'))
#define INT2CHAR(a) ((a) + '0')

static void
save_int_c (char *s, int i, int c)
{
    for (;;) {
        c = CHAR2INT(s[i]) + c;
        s[i] = INT2CHAR(c % 10);
        c /= 10;
        i--;
        if (c == 0) {
            break;
        }
    }
    return;
}

static void
remove_heading_zero (char *s, int len)
{
    char *start;
    int real_len;

    start = s;
    while (*start == '\0') {
        start++;
    }
    while (*start == '0') {
        if (*(start+1) != '0') {
            break;
        }
        start++;
    }
    if (s == start) {
        return;
    }
    real_len = (unsigned long)start - (unsigned long)s;
    real_len = len + 1 - real_len;
    memmove(s, start, real_len);
    return;
}

static char *
multiply_internal (char *a, char *b)
{
    int i, j, len_a, len_b, len_c;
    int int_a, int_b, int_c;
    char *c;

    len_a = strlen(a);
    len_b = strlen(b);
    len_c = len_a + len_b;

    c = calloc(len_c+1, sizeof(char));
    if (c == NULL) {
        return NULL;
    }

    for (i = 0; i < len_a; i++) {
        int_a = CHAR2INT(a[len_a-i-1]);
        for (j = 0; j < len_b; j++) {
            int_b = CHAR2INT(b[len_b-j-1]);
            int_c = int_a * int_b;
            save_int_c(c, len_c-(i+j)-1, int_c);
        }
    }

    remove_heading_zero(c, len_c);
    return c;
}

static char *
multiply (char *a, char *b)
{
    char *c;
    int len_a, len_b;

    if (a == NULL || b == NULL) {
        return NULL;
    }

    len_a = strlen(a);
    len_b = strlen(b);
    if (len_a == 0 || len_b == 0) {
        return NULL;
    }

    if (len_a >= len_b) {
        c = multiply_internal(a, b);
    } else {
        c = multiply_internal(b, a);
    }
    return c;
}

#include <stdio.h>

static int
calc_len (int a)
{
    int len = 1;

    a /= 10;
    while (a > 0) {
        len++;
        a /= 10;
    }
    return len;
}

static char *
int_to_str (int a)
{
    char *s;
    int i, len;

    len = calc_len(a);
    s = calloc(len+1, sizeof(char));
    if (s == NULL) {
        return NULL;
    }

    if (a == 0) {
        s[0] = '0';
        return s;
    }

    i = len-1;
    while (a > 0) {
        s[i] = INT2CHAR(a%10);
        a /= 10;
        i--;
    }

    return s;
}

static int
test_multiply_int (int int_a, int int_b)
{
    char *answer, *result, *a, *b;
    int rc;

    a = int_to_str(int_a);
    b = int_to_str(int_b);
    answer = int_to_str(int_a*int_b);

    result = multiply(a, b);
    if (strcmp(result, answer) != 0) {
        printf("%s*%s, get %s, should be %s\n", a, b, result, answer);
        rc = -1;
    } else {
        rc = 0;
    }

    free(a);
    free(b);
    free(answer);
    free(result);
    return rc;
}

static void
test_case (void)
{
    int i, j, rc;

    for (i = 0; i <= 1000; i++) {
        for (j = 0; j <= 1000; j++) {
            rc = test_multiply_int(i, j);
            if (rc != 0) {
                return;
            }
        }
    }
    return;
}

int main (void)
{
    test_case();
    return 0;
}

