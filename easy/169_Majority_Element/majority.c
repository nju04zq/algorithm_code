int
majorityElement (int *nums, int size)
{
    int cnt, i, j;

    for (i = 0, cnt = 0; i < size; i++) {
        if (cnt == 0) {
            j = i;
            cnt++;
            continue;
        }

        if (nums[i] == nums[j]) {
            cnt++;
        } else {
            cnt--;
        }
    }

    if (cnt == 0) {
        return -1;
    }

    for (i = 0, cnt = 0; i < size; i++) {
        if (nums[i] == nums[j]) {
            cnt++;
        }
    }

    if (cnt > size/2) {
        return nums[j];
    } else {
        return -1;
    }
}

#include <stdio.h>

static void
test_majority (int *nums, int size, int answer)
{
    int majority, i;

    majority = majorityElement(nums, size);
    if (majority == answer) {
        return;
    }

    printf("Majority element in ");
    for (i = 0; i < size; i++) {
        printf("%d ", nums[i]);
    }
    printf(", get %d, should be %d\n", majority, answer);
    return;
}

#define ARRAY_LEN(a) (sizeof(a)/sizeof((a)[0]))

int main (void)
{
    int a0[] = {1};
    int a1[] = {1, 2};
    int a2[] = {1, 1};
    int a3[] = {1, 2, 3};
    int a4[] = {1, 1, 3};
    int a5[] = {1, 1, 3, 2};
    int a6[] = {1, 1, 1, 2};

    test_majority(a0, ARRAY_LEN(a0), 1);
    test_majority(a1, ARRAY_LEN(a1), -1);
    test_majority(a2, ARRAY_LEN(a2), 1);
    test_majority(a3, ARRAY_LEN(a3), -1);
    test_majority(a4, ARRAY_LEN(a4), 1);
    test_majority(a5, ARRAY_LEN(a5), -1);
    test_majority(a6, ARRAY_LEN(a6), 1);
    return 0;
}

