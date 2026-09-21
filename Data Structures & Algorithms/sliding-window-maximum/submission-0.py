from collections import deque


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        q = deque()
        l = 0
        ans = []
        for r in range(0, len(nums)):
            while len(q) > 0 and q[-1][1] <= nums[r]:
                q.pop()
            while len(q) > 0 and q[0][0] < r-k+1:
                q.popleft()

            q.append((r, nums[r]))
            if r >= k-1:
                ans.append(q[0][1])
        return ans
