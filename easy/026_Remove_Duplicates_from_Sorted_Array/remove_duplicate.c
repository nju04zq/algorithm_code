#include <stddef.h>

int
removeDuplicates (int *nums, int cnt)
{
    int i, j, prev;

    if (nums == NULL || cnt == 0) {
        return 0;
    }

    prev = nums[0];
    for (i = 1, j = 1; i < cnt; i++) {
        if (nums[i] == prev) {
            continue;
        }
        nums[j++] = nums[i];
        prev = nums[i];
    }

    return j;
}

#include <stdio.h>

static void
dump_nums (int *nums, int cnt)
{
    int i;

    for (i = 0; i < cnt; i++) {
        printf("%d ", nums[i]);
    }
    printf("\n");
    return;
}

static void
test_remove_duplicate (int *nums, int cnt)
{
    printf("Before remove: ");
    dump_nums(nums, cnt);

    cnt = removeDuplicates(nums, cnt);

    printf("After remove: ");
    dump_nums(nums, cnt);
    return;
}

int main (void)
{
    int a0[] = {};
    int a1[] = {1, 2, 3, 4};
    int a2[] = {1, 1, 2, 2, 3, 3, 4};

    test_remove_duplicate(a0, 0);
    test_remove_duplicate(a1, 4);
    test_remove_duplicate(a2, 7);
    return 0;
}
