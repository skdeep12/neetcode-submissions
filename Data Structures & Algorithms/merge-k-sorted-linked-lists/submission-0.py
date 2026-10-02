# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:    
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        h = []
        for idx, ls in enumerate(lists):
            if ls is not None:
                heapq.heappush(h, (ls.val, idx))
        
        ans = None
        current = None
        while len(h) > 0:
            val, idx = heapq.heappop(h)
            if ans is None:
                ans = lists[idx]
                current = lists[idx]
            else:
                current.next = lists[idx]
                current = current.next
            
            if lists[idx].next is not None:
                lists[idx] = lists[idx].next
                heapq.heappush(h, (lists[idx].val,idx))
        return ans
            