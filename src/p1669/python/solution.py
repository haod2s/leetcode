class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        p, q = list1, list1
        for i in range(a-1):
            p = p.next
        for i in range(b):
            q = q.next
        r = list2
        while r.next != None:
            r = r.next
        p.next = list2
        r.next = q.next
        
        return list1
    
# 华为