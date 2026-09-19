class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        q = [(0,0)]
        visited = [False]*n
        cost = 0
        while len(q) > 0:
            front = heapq.heappop(q)
            if visited[front[1]]:
                continue
            cost += front[0]
            visited[front[1]] = True
            for i in range(0,n):
                if not visited[i]:
                    dist = abs(points[i][0]-points[front[1]][0]) + abs(points[i][1]-points[front[1]][1])
                    heapq.heappush(q, (dist,i))
        return cost
