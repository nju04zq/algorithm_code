#include <string.h>

#define VER_NUM_MAX_LEN 63

#define VER_NUM_SEP '.'

static char *
skip_heading_zero (char *str)
{
    if (*str == '\0') {
        return str;
    }

    for (; *str; str++) {
        if (*str != '0') {
            break;
        }
    }
    if (*str == '\0') {
        str--;
    }
    return str;
}

static int
compare_ver_num (char *num1, char *num2)
{
    int len1, len2, i;

    num1 = skip_heading_zero(num1);
    num2 = skip_heading_zero(num2);

    len1 = strlen(num1);
    len2 = strlen(num2);
    if (len1 < len2) {
        return -1;
    } else if (len1 > len2) {
        return 1;
    }

    for (i = 0; i < len1; i++) {
        if (num1[i] < num2[i]) {
            return -1;
        } else if (num1[i] > num2[i]) {
            return 1;
        }
    }

    return 0;
}

static int
get_ver_num (char *str, char *num)
{
    char *p;
    int len = 0;

    num[0] = '0';
    num[1] = '\0';

    for (p = str; *p; p++) {
        if (*p == VER_NUM_SEP) {
            break;
        }
        len++;
    }

    if (len == 0 || len > VER_NUM_MAX_LEN) {
        return 0;
    }

    strncpy(num, str, len);
    num[len] = '\0';
    return len;
}

static char *
move_to_next_ver_num (char *str, int len)
{
    str += len;
    if (*str == VER_NUM_SEP) {
        str++;
    }
    return str;
}

int
compareVersion (char *ver_str1, char *ver_str2)
{
    char num1[VER_NUM_MAX_LEN+1], num2[VER_NUM_MAX_LEN+1];
    int len1, len2, result;

    //TODO verify version string

    result = 0;
    for (;;) {
        len1 = get_ver_num(ver_str1, num1);
        len2 = get_ver_num(ver_str2, num2);
        if (len1 == 0 && len2 == 0) {
            break;
        }

        result = compare_ver_num(num1, num2);
        if (result != 0) {
            break;
        }

        ver_str1 = move_to_next_ver_num(ver_str1, len1);
        ver_str2 = move_to_next_ver_num(ver_str2, len2);
    }

    return result;
}

#include <stdio.h>

static void
test_cmp_ver (char *str1, char *str2, int answer)
{
    int result;

    result = compareVersion(str1, str2);
    if (result != answer) {
        printf("%s, %s, result %d, should be %d\n", str1, str2, result, answer);
    }
    return;
}

int main (void)
{
    test_cmp_ver("1", "1.1", -1);
    test_cmp_ver("1", "1.0", 0);
    test_cmp_ver("0.1", "1.1", -1);
    test_cmp_ver("1.1", "1.2", -1);
    test_cmp_ver("1.1", "1.1", 0);
    test_cmp_ver("1.2", "13.37", -1);
    test_cmp_ver("1", "2", -1);
    test_cmp_ver("1", "0", 1);
    test_cmp_ver("1", "1", 0);
    test_cmp_ver("1.002", "1.37", -1);
    test_cmp_ver("1.002", "1.2", 0);
    test_cmp_ver("1.00", "1.0", 0);
    return 0;
}

