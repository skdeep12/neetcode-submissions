class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        f = dict()
        for t in tasks:
            f[t] = f.get(t, 0) + 1
        h = []
        for k,v in f.items():
            heapq.heappush(h,[-v,k])
        # print(h)
        k = 0
        while len(h) > 0:
            r = []
            for i in range(0,n+1):
                if len(h) == 0:
                    if len(r) > 0:
                        k+=1
                    continue 
                k+=1
                item = heapq.heappop(h)
                item[0]+=1
                if item[0] < 0:
                    r.append(item)
            # print(r,k)
            for i in r:
                heapq.heappush(h,i)
        return k
                