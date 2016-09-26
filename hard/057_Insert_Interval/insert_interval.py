class Solution(object):
    def interval_intersect(self, p1, p2):
        if p1.start <= p2.start <= p1.end:
            return True
        if p2.start <= p1.start <= p2.end:
            return True
        return False

    def insert(self, intervals, p):
        """
        :type intervals: List[Interval]
        :type newInterval: Interval
        :rtype: List[Interval]
        """
        result = []
        for cur in intervals:
            if p is None:
                result.append(cur)
            elif self.interval_intersect(cur, p):
                p.start = min(p.start, cur.start)
                p.end = max(p.end, cur.end)
            elif p.start < cur.start:
                result.append(p)
                result.append(cur)
                p = None
            else:
                result.append(cur)
        if p is not None:
            result.append(p)
        return result
