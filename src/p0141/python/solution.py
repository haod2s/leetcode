# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head == None:
            return False
        fast, slow = head, head
        while fast != None and slow != None:
            if fast.next == None:
                return False
            fast = fast.next.next
            slow = slow.next
            if slow == fast:
                return True
            
        return False    