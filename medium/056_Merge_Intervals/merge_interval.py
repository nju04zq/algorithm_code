# Definition for an interval.
class Interval(object):
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution(object):
    def interval_intersect(self, p1, p2):
        if p1.start <= p2.start <= p1.end:
            return True
        if p2.start <= p1.start <= p2.end:
            return True
        return False

    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        result = []
        intervals.sort(key=lambda x:x.start)

        prev = None
        for cur in intervals:
            if prev is not None and self.interval_intersect(cur, prev):
                prev.start = min(prev.start, cur.start)
                prev.end = max(prev.end, cur.end)
            else:
                result.append(cur)
                prev = cur
        return result

def dump_intervals(intervals):
    for p in intervals:
        print "[{}, {}]".format(p.start, p.end)

def test_merge(intervals):
    print "before merge"
    dump_intervals(intervals)
    result = solution.merge(intervals)
    print "after merge"
    dump_intervals(result)

def test_case_1():
    i0 = Interval(3, 4)
    i1 = Interval(1, 2)
    i2 = Interval(5, 6)
    test_merge([i0, i1, i2])

def test_case_2():
    i0 = Interval(2, 6)
    i1 = Interval(1, 5)
    i2 = Interval(3, 7)
    i3 = Interval(4, 6)
    test_merge([i0, i1, i2, i3])

solution = Solution()
test_case_1()
test_case_2()

