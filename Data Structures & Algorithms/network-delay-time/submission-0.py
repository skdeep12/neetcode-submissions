class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        g = [[] for i in range(0,n+1)]
        for u,v,t in times:
            g[u].append((v,t))
        h = [(0,k)]
        visited = [False]*(n+1)
        ans = None
        while len(h) > 0:
            dist, u  = heapq.heappop(h)
            
            if visited[u]:
                continue
            else:
                visited[u] = True
            if ans is None or ans < dist:
                ans = dist
            for v,t in g[u]:
                if not visited[v]:
                    heapq.heappush(h,(dist+t, v))
        for i in range(1,n+1):
            if not visited[i]:
                return -1
        return ans
        