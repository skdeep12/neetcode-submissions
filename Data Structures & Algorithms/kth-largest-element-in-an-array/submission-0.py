class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        n = len(nums)
        s = n-k+1
        h = []
        for n in nums:
            heapq.heappush(h,-n)
            if len(h) > s:
                heapq.heappop(h)
        return -h[0]