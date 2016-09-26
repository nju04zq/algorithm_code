# array contains number from 0 - 29, now put those in the range [0, 9] first,
# then those in range [10, 19], in the end those in range [20, 29].
# Do not need keep the relative order in original array.
# Should run in O(n) and use space with O(1).

import sys
import random

def sort_number1(a):
    j, k = 0, 0
    for i in xrange(len(a)):
        if a[i] < 10:
            a[i], a[j] = a[j], a[i]
            j += 1
        if a[i] < 20:
            a[i], a[k] = a[k], a[i]
            k += 1
        if k < j:
            k += 1

def sort_number2(a):
    j, k = 0, 0
    for i in xrange(len(a)):
        temp = a[i]
        if a[i] < 10:
            a[i], a[j] = a[j], a[i]
        if a[i] < 20:
            a[i], a[j+k] = a[j+k], a[i]
        if temp < 10:
            j += 1
        elif temp < 20:
            k += 1

def split_array(a, start, end, mid_value):
    j = start
    for i in xrange(start, end):
        if a[i] < mid_value:
            a[i], a[j] = a[j], a[i]
            j += 1
    return j

def sort_number3(a):
    start, end = 0, len(a)
    mid_value = 10
    while mid_value < 30:
        if start >= end:
            break
        next_start = split_array(a, start, end, mid_value)
        start = next_start
        mid_value += 10

def test_sorted(i, sorted_a, original):
    lower, upper = 0, 9
    for i in xrange(len(sorted_a)):
        if a[i] < lower: 
            print "failed on #{} {}, {}".format(i, original, sorted_a)
            sys.exit()
        if a[i] > upper:
            lower += 10
            upper += 10

def generate_a_len():
    return random.randint(1, 20)

def generate_a(n):
    a = [0 for i in xrange(n)]
    for i in xrange(n):
        a[i] = random.randint(0, 29)
    return a

for i in xrange(100000):
    n = generate_a_len()
    a = generate_a(n)
    b = a[:]
    sort_number1(a)
    test_sorted(i, a, b)

    a = b[:]
    sort_number2(a)
    test_sorted(i, a, b)

    a = b[:]
    sort_number3(a)
    test_sorted(i, a, b)

print "Passed #{} cases".format(i+1)

