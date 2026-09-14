class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.h = []
        self.k = k
        for n in nums:
            heapq.heappush(self.h, n)
            if len(self.h) > k:
                heapq.heappop(self.h)
        

    def add(self, val: int) -> int:
        heapq.heappush(self.h, val)
        if len(self.h) > self.k:
                heapq.heappop(self.h)
        return self.h[0]
