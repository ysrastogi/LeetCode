# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    
        min_heap = []
        for i, node in enumerate(lists):
            if node:
                heapq.heappush(min_heap, (node.val, i, node))

        dummy = ListNode(0)
        current = dummy
        
        while min_heap:
            val, idx, node = heapq.heappop(min_heap)
            current.next = ListNode(val)
            current = current.next
            if node.next:  # push the next node from the same list
                heapq.heappush(min_heap, (node.next.val, idx, node.next))

        return dummy.next
            