# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 == None and l2 == None:
            return None
        p1, p2 = l1, l2
        temp = list()
        while p1 != None and p2 != None:
            if p1.val < p2.val:
                temp.append(p1.val)
                p1 = p1.next
            else:
                temp.append(p2.val)
                p2 = p2.next
        while p1 != None:
            temp.append(p1.val)
            p1 = p1.next
        while p2 != None:
            temp.append(p2.val)
            p2 = p2.next
        ans = ListNode(temp[0])
        p = ans
        for i in range(1, len(temp)):
            p.next = ListNode(temp[i])
            p = p.next
        
        return ans