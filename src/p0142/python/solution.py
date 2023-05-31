# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None or head.next == None:
            return None
        slow, fast = head, head.next.next
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if fast == slow:
                break
        if fast == None or fast.next == None:
            return None
        existed = set()
        while fast not in existed:
            existed.add(fast)
            fast = fast.next
        if head in existed:
            return head
        slow = head
        while slow.next not in existed:
            slow = slow.next

        return slow.next
        
# 腾讯