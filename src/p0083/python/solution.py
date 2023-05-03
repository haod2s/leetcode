# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return head
        p1, p2 = head, head.next
        while p2 != None:
            if p1.val != p2.val:
                p1.next = p2
                p1 = p2
                p2 = p1.next
            else:
                p2 = p2.next
        p1.next = None
            
        return head