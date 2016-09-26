import random

def max_subarray_dp(a):
    dp_prev, max_sub_total = 0, float("-inf")
    for i in xrange(0, len(a)):
        if dp_prev > 0:
            dp_now = dp_prev + a[i]
        else:
            dp_now = a[i]
        max_sub_total = max(dp_now, max_sub_total)
        dp_prev = dp_now
    return max_sub_total

def max_subarray_bf(a):
    max_sub_total = float("-inf")
    for i in xrange(0, len(a)):
        total = 0
        for j in xrange(i, len(a)):
            total += a[j]
            max_sub_total = max(total, max_sub_total)
    return max_sub_total

def generate_a_len():
    return random.randint(1, 20)

def generate_a(n):
    a = [0 for i in xrange(n)]
    for i in xrange(n):
        a[i] = random.randint(-29, 29)
    return a

def test_once(i):
    n = generate_a_len()
    a = generate_a(n)
    max_dp = max_subarray_dp(a)
    max_bf = max_subarray_bf(a)
    if max_dp != max_bf:
        raise Exception("Test fail on case {}, a {}, max_dp {}, max_bf {}".\
                        format(i, a, max_dp, max_bf))

def test_solution():
    for i in xrange(10000):
        test_once(i)
    print "Passed {} cases".format(i+1)

test_solution()
