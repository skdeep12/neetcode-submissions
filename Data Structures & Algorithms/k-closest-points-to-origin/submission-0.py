class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        h = []
        for p in points:
            dist = math.sqrt(p[0]*p[0]+p[1]*p[1])
            heapq.heappush(h,(-dist, p))
            if len(h) > k:
                heapq.heappop(h)
        ans = []
        for item in h:
            ans.append(item[1])
        return ans