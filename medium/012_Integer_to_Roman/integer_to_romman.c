#include <stddef.h>
#include <stdlib.h>
#include "../../common/bool.h"
#include "../../common/stack.h"

#define ROMAN_MAX 3999

#define ROMAN_BASE_MIN 10
#define ROMAN_BASE_MAX 10000
#define ROMAN_BASE_INC 10

typedef struct roman_base_s {
    int num;
    char ch;
} roman_base_t;

static roman_base_t base_1 = {1, 'I'};
static roman_base_t base_5 = {5, 'V'};
static roman_base_t base_10 = {10, 'X'};
static roman_base_t base_50 = {50, 'L'};
static roman_base_t base_100 = {100, 'C'};
static roman_base_t base_500 = {500, 'D'};
static roman_base_t base_1000 = {1000, 'M'};

#define ROMAN_BASE_GROUP_SIZE 3

static void
fill_roman_base_group (int base, roman_base_t **base_group)
{
    switch (base) {
    case 10:
        base_group[0] = &base_1;
        base_group[1] = &base_5;
        base_group[2] = &base_10;
        break;
    case 100:
        base_group[0] = &base_10;
        base_group[1] = &base_50;
        base_group[2] = &base_100;
        break;
    case 1000:
        base_group[0] = &base_100;
        base_group[1] = &base_500;
        base_group[2] = &base_1000;
        break;
    case 10000:
        base_group[0] = &base_1000;
        base_group[1] = NULL;
        base_group[2] = NULL;
        break;
    default:
        base_group[0] = NULL;
        base_group[1] = NULL;
        base_group[2] = NULL;
        break;
    }
    return;
}

static void
remain_to_roman (ustack_t *stack, int remain, roman_base_t **base_group)
{
    int i, val;
    bool need_middle_base = FALSE;

    if (remain == 0) {
        return;
    }

    for (i = 0; i < ROMAN_BASE_GROUP_SIZE; i++) {
        if (base_group[i] == NULL) {
            break;
        }
        if (remain == base_group[i]->num) {
            push(stack, (void *)(unsigned long)(base_group[i]->ch));
            return;
        }
    }

    for (i = 1; i < ROMAN_BASE_GROUP_SIZE; i++) {
        if (base_group[i] == NULL) {
            break;
        }
        val = base_group[i]->num - base_group[0]->num;
        if (remain == val) {
            push(stack, (void *)(unsigned long)(base_group[i]->ch));
            push(stack, (void *)(unsigned long)(base_group[0]->ch));
            return;
        }
    }

    if (base_group[1] && remain > base_group[1]->num) {
        remain -= base_group[1]->num;
        need_middle_base = TRUE;
    }

    for (i = 1; i <= ROMAN_BASE_GROUP_SIZE && remain > 0; i++) {
        push(stack, (void *)(unsigned long)(base_group[0]->ch));
        remain -= base_group[0]->num;
    }

    if (need_middle_base) {
        push(stack, (void *)(unsigned long)(base_group[1]->ch));
    }

    return;
}

static void
int_to_roman_internal (ustack_t *stack, int num)
{
    int base = 10, remain;
    roman_base_t *base_group[ROMAN_BASE_GROUP_SIZE];

    for (base = ROMAN_BASE_MIN;
         base <= ROMAN_BASE_MAX;
         base*= ROMAN_BASE_INC) {
        if (num == 0) {
            break;
        }
        remain = num % base;
        fill_roman_base_group(base, base_group);
        remain_to_roman(stack, remain, base_group);
        num -= remain;
    }

    return;
}

static char *
copy_stack_str (ustack_t *stack)
{
    int i, size;
    char *s;

    size = get_stack_size(stack);
    s = calloc(size+1, sizeof(char));
    if (s == NULL) {
        return NULL;
    }

    for (i = 0; i < size; i++) {
        s[i] = (char)(unsigned long)pop(stack);
    }

    return s;
}

char *
intToRoman (int num)
{
    ustack_t stack;
    char *s = NULL;
    int rc;

    if (num <= 0 || num > ROMAN_MAX) {
        return NULL;
    }

    rc = init_stack(&stack);
    if (rc != 0) {
        return NULL;
    }

    int_to_roman_internal(&stack, num);

    s = copy_stack_str(&stack);
    clean_stack(&stack);
    return s;
}

#include <stdio.h>
#include <string.h>

static void
test_int_to_roman (int num, char *answer)
{
    char *result;

    result = intToRoman(num);
    if (result == NULL) {
        printf("INT %d, get <NULL>\n", num);
        return;
    }
    if (strcmp(result, answer) != 0) {
        printf("INT %d, get %s, should be %s\n", num, result, answer);
    }
    free(result);
    return;
}

int main (void)
{
    test_int_to_roman(1, "I");
    test_int_to_roman(2, "II");
    test_int_to_roman(3, "III");
    test_int_to_roman(4, "IV");
    test_int_to_roman(5, "V");
    test_int_to_roman(6, "VI");
    test_int_to_roman(7, "VII");
    test_int_to_roman(8, "VIII");
    test_int_to_roman(9, "IX");
    test_int_to_roman(10, "X");
    test_int_to_roman(11, "XI");
    test_int_to_roman(22, "XXII");
    test_int_to_roman(33, "XXXIII");
    test_int_to_roman(44, "XLIV");
    test_int_to_roman(50, "L");
    test_int_to_roman(55, "LV");
    test_int_to_roman(66, "LXVI");
    test_int_to_roman(77, "LXXVII");
    test_int_to_roman(88, "LXXXVIII");
    test_int_to_roman(99, "XCIX");
    test_int_to_roman(100, "C");
    test_int_to_roman(102, "CII");
    test_int_to_roman(199, "CXCIX");
    test_int_to_roman(400, "CD");
    test_int_to_roman(500, "D");
    test_int_to_roman(900, "CM");
    test_int_to_roman(1000, "M");
    test_int_to_roman(1437, "MCDXXXVII");
    test_int_to_roman(3333, "MMMCCCXXXIII");
    return 0;
}

