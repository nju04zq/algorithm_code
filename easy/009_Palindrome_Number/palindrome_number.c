typedef unsigned char bool;
#define TRUE 1
#define FALSE 0

static int
calc_div_high (int x)
{
    int div_high = 1;

    while (x >= 10) {
        x = x/10;
        div_high *= 10;
    }

    return div_high;
}

static int
discard_high (int x, int div_high)
{
    return x % div_high;
}

static int
extract_high (int x, int div_high)
{
    return x/div_high;
}

static int
discard_low (int x)
{
    return x/10;
}

static int
extract_low (int x)
{
    return x % 10;
}

static bool
check_palindrome (int x)
{
    int high, low, div_high;

    div_high = calc_div_high(x);

    while (x > 0) {
        low = extract_low(x);
        high = extract_high(x, div_high);
        if (low != high) {
            return FALSE;
        }
        x = discard_high(x, div_high);
        x = discard_low(x);
        div_high /= 100;
    }

    return TRUE;
}

bool
isPalindrome (int x)
{
    bool result;

    if (x < 0) {
        return FALSE;
    }

    result = check_palindrome(x);
    return result;
}

#include <stdio.h>

static void
test_check (int x)
{
    printf("%d, %d\n", x, isPalindrome(x));
    return;
}

int main (void)
{
//    test_check(0);
//    test_check(10);
//    test_check(-1);
//    test_check(123);
//    test_check(121);
//    test_check(1221);
    test_check(1000021);
    return 0;
}

