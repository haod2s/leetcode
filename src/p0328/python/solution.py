# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return head
        odd, even = head, head.next
        odd.next = even.next
        even.next = None
        p, q = odd.next, even
        while p != None and p.next != None:
            q.next = p.next
            p.next = p.next.next
            q = q.next
            q.next = None
            p = p.next
        if p == None:
            p = odd
            while p.next != None:
                p = p.next
        p.next = even
        
        return odd
    
# 华为