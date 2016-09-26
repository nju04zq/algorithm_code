#include <math.h>

static float
mySqrt_float (float x)
{
    float xhalf = 0.5f*x;
    int i = *(int*)&x; // get bits for floating VALUE 
    i = 0x5f375a86- (i>>1); // gives initial guess y0
    x = *(float*)&i; // convert bits BACK to float
    x = x*(1.5f-xhalf*x*x); // Newton step, repeating increases accuracy
    x = x*(1.5f-xhalf*x*x); // Newton step, repeating increases accuracy
    x = x*(1.5f-xhalf*x*x); // Newton step, repeating increases accuracy
    return 1/x;
}

int mySqrt(int x)
{
    float result;
    int ret, sqr1, sqr2, sqr3;

    result = mySqrt_float((float)x);
    ret = (int)result;
    
    sqr2 = ret * ret;
    sqr3 = sqr2 + 2*ret + 1;
    sqr1 = sqr2 - 2*ret + 1;
    if (sqr1 == x) {
        return ret-1;
    } else if (sqr2 == x) {
        return ret;
    } else if (sqr3 == x) {
        return ret+1;
    }
    if (sqr2 > x) {
        return ret-1;
    } else {
        return ret;
    }
}

#include <stdio.h>
#include <limits.h>

int
test_sqrt (int x)
{
    int result, answer, rc = 0;

    result = mySqrt(x);
    answer = (int)sqrt((double)x);
    if (result != answer) {
        printf("Fail on %d, get %d, should be %d\n", x, result, answer);
        rc = -1;
    }
    return rc;
}

#define TEST_MAX_NUM 10000

int main (void)
{
    int i, rc;

    for (i = 0; i <= TEST_MAX_NUM; i++) {
        rc = test_sqrt(i);
        if (rc != 0) {
            break;
        }
    }
    test_sqrt(INT_MAX);
    test_sqrt(2147395599);
    return 0;
}

