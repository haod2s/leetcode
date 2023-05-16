class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head == None:
            return None
        elif head.next == None:
            return head
        p = head
        h = self.reverseList(p.next)
        q = h
        while q.next != None: 
            q = q.next
        q.next = p
        p.next = None
        
        return h

# 华为