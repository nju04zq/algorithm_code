import Queue
class Solution(object):
    def scheduleCourse(self, courses):
        """
        :type courses: List[List[int]]
        :rtype: int
        """
        courses.sort(key=lambda x:x[1])
        pq = Queue.PriorityQueue()
        timeline = 0
        for course in courses:
            if timeline + course[0] <= course[1]:
                pq.put(-course[0])
                timeline += course[0]
            elif not pq.empty():
                longest = -pq.get()
                if course[0] >= longest:
                    pq.put(-longest)
                else:
                    pq.put(-course[0])
                    timeline += (course[0]-longest)
        return pq.qsize()
