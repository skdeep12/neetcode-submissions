class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        h = []
        for s in stones:
            heapq.heappush(h,-s)
        
        while len(h) > 1:
            first = heapq.heappop(h)
            second = heapq.heappop(h)
            left = abs(first-second)
            if left!=0:
                heapq.heappush(h,-left)
        if len(h) == 0:
            return 0
        else:
            return -h[0]