#include <stdint.h>
#include <stdlib.h>

#define UINT8_BITS 4 //can change it to 4 or 8
#define UINT32_BITS 32
#define REVERSE_MASK ((1<<UINT8_BITS) - 1)

#define LOOKUP_TBL_SIZE (1<<UINT8_BITS)
static uint8_t *lookup_tbl = NULL;

static uint8_t
reverse_uint8 (uint8_t x)
{
    uint8_t y = 0;
    int i;

    for (i = 0; i < UINT8_BITS; i++) {
        y = (y << 1) | (x & 0x1);
        x = x >> 1;
    }
    return y;
}

static void
setup_lookup_tbl (void)
{
    uint8_t *p;
    int i;

    p = calloc(LOOKUP_TBL_SIZE, sizeof(uint8_t));
    if (p == NULL) {
        return;
    }

    for (i = 0; i < LOOKUP_TBL_SIZE; i++) {
        p[i] = reverse_uint8(i);
    }
    lookup_tbl = p;
    return;
}

uint32_t
reverseBits (uint32_t x)
{
    int i, cnt;
    uint32_t y = 0, mask, reversed;

    if (lookup_tbl == NULL) {
        setup_lookup_tbl();
    }
    if (lookup_tbl == NULL) {
        return 0;
    }

    mask = REVERSE_MASK;
    cnt = UINT32_BITS/UINT8_BITS;
    for (i = 0; i < cnt; i++) {
        reversed = lookup_tbl[x & mask];
        y = (y << UINT8_BITS) | reversed;
        x = x >> UINT8_BITS;
    }

    return y;
}

#include <stdio.h>

static void
test_reverse (uint32_t x, uint32_t answer)
{
    uint32_t y;

    y = reverseBits(x);
    if (y != answer) {
        printf("Reverse bits of %x, get %x, should be %x", x, y, answer);
    }
    return;
}

int main (void)
{
    test_reverse(0x02941e9c, 0x39782940);
    return 0;
}

