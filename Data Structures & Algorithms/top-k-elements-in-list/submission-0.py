class Solution:
    def topKFrequent(self, nums: List[int], K: int) -> List[int]:
        h = []
        m = {}
        for n in nums:
            m[n] = m.get(n,0) + 1 
        for k,v in m.items():
            heapq.heappush(h, (-v, k))
        ans = []
        for i in range(0,K):
            top = heapq.heappop(h)
            ans.append(top[1])
        return ans