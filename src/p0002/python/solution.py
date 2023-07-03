# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        p1, p2 = l1, l2
        plus = 0
        dummy = ListNode()
        p = dummy
        while p1 or p2 or plus:
            p.next = ListNode()
            if p1:
                p.next.val += p1.val
                p1 = p1.next
            if p2:
                p.next.val += p2.val
                p2 = p2.next
            p.next.val += plus
            plus = p.next.val // 10
            p.next.val %= 10
            p = p.next
        
        return dummy.next