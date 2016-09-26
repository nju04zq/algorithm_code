#include <stdlib.h>
#include <string.h>

static int
transform_one_roman (char r)
{
    switch (r) {
    case 'I':
        return 1;
    case 'V':
        return 5;
    case 'X':
        return 10;
    case 'L':
        return 50;
    case 'C':
        return 100;
    case 'D':
        return 500;
    case 'M':
        return 1000;
    default:
        return -1;
    }
    return -1;
}

static void
transform (char *str, int *a)
{
    int x, prefix_start = 0, i = 0, prev = 0;
    int str_len;

    str_len = strlen(str);

    for (; *str; str++) {
        x = transform_one_roman(*str);
        if (x < 0) {
            continue;
        }

        if (prev == 0 || x <= prev) {
            if (x < prev) {
                prefix_start = i;
            }
            a[i++] = x;
            prev = x;
            continue;
        }

        a[prefix_start++] = x-prev;
        i = prefix_start;
        prev = x;
    }

    for (; i < str_len; i++) {
        a[i] = 0;
    }
    return;
}

static int
calc_sum (int *a, int cnt)
{
    int i, sum = 0;

    for (i = 0; i < cnt; i++) {
        sum += a[i];
    }

    return sum;
}

int
romanToInt (char *s)
{
    int *a, x;

    if (s == NULL || *s == '\0') {
        return 0;
    }

    a = calloc(strlen(s), sizeof(int));    
    if (a == NULL) {
        return 0;
    }

    transform(s, a);
    x = calc_sum(a, strlen(s));

    free(a);
    return x;
}

#include <stdio.h>

static void
test_romanToInt (char *s, int x)
{
    int y;
    
    y = romanToInt(s);
    if (y != x) {
        printf("%s, calc %d, should be %d\n", s, y, x);
    }
    return;
}

int main (void)
{
    test_romanToInt("XIIX", 19);
    test_romanToInt("AXIIXxxi", 19);


    test_romanToInt("I", 1);
    test_romanToInt("II", 2);
    test_romanToInt("III", 3);
    test_romanToInt("IV", 4);
    test_romanToInt("V", 5);
    test_romanToInt("VI", 6);
    test_romanToInt("VII", 7);
    test_romanToInt("VIII", 8);
    test_romanToInt("IX", 9);
    test_romanToInt("X", 10);
    test_romanToInt("XI", 11);
    test_romanToInt("XII", 12);
    test_romanToInt("XIII", 13);
    test_romanToInt("XIV", 14);
    test_romanToInt("XV", 15);
    test_romanToInt("XVI", 16);
    test_romanToInt("XVII", 17);
    test_romanToInt("XVIII", 18);
    test_romanToInt("XIX", 19);
    test_romanToInt("XX", 20);
    test_romanToInt("XXX", 30);
    test_romanToInt("XL", 40);
    test_romanToInt("L", 50);
    test_romanToInt("LX", 60);
    test_romanToInt("LXX", 70);
    test_romanToInt("LXXX", 80);
    test_romanToInt("XC", 90);
    test_romanToInt("XCIX", 99);
    test_romanToInt("C", 100);
    test_romanToInt("CI", 101);
    test_romanToInt("CII", 102);
    test_romanToInt("CXCIX", 199);
    test_romanToInt("CC", 200);
    test_romanToInt("CCC", 300);
    test_romanToInt("CD", 400);
    test_romanToInt("D", 500);
    test_romanToInt("DC", 600);
    test_romanToInt("DCCC", 800);
    test_romanToInt("CM", 900);
    test_romanToInt("M", 1000);
    test_romanToInt("MCD", 1400);
    test_romanToInt("MCDXXXVII", 1437);
    test_romanToInt("MD", 1500);
    test_romanToInt("MDCCC", 1800);
    test_romanToInt("MDCCCLXXX", 1880);
    test_romanToInt("MCM", 1900);
    test_romanToInt("MM", 2000);
    test_romanToInt("MMM", 3000);
    test_romanToInt("MMMCCCXXXIII", 3333);
}

