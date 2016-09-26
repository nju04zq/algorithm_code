#include <stddef.h>

int
removeElement (int *nums, int cnt, int val)
{
    int i, j;

    if (nums == NULL || cnt == 0) {
        return 0;
    }
    
    for (i = 0, j = 0; i < cnt; i++) {
        if (nums[i] != val) {
            nums[j++] = nums[i];
        }
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
test_remove (int *nums, int cnt, int val)
{
    printf("Before remove %d:", val);
    dump_nums(nums, cnt);

    cnt = removeElement(nums, cnt, val);

    printf("After  remove %d:", val);
    dump_nums(nums, cnt);
    return;
}

int main (void)
{
    int a0[] = {};
    int a1[] = {1};
    int a2[] = {3, 2, 3, 4, 2};
    int a3[] = {3, 2, 3, 4, 2};

    test_remove(a0, 0, 1);
    test_remove(a1, 1, 2);
    test_remove(a1, 1, 1);
    test_remove(a2, 5, 1);
    test_remove(a3, 5, 2);
    return 0;
}
