/**
 * Given the head of a singly linked list, reverse the list, and return the reversed list.
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * Input: head = [1,2,3,4,5]
 * Output: [5,4,3,2,1]
 * Example 2:
 * 
 * 
 * Input: head = [1,2]
 * Output: [2,1]
 * Example 3:
 * 
 * Input: head = []
 * Output: []
 *  
 * 
 * Constraints:
 * 
 * The number of nodes in the list is the range [0, 5000].
 * -5000 <= Node.val <= 5000
 */


func reverseList(head *ListNode) *ListNode {
	h, _ := reverseListHelper(head)
	return h
}

func reverseListHelper(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	} else if head.Next == nil {
		return head, head
	} else if head.Next.Next == nil {
		p := head.Next
		p.Next = head
		head.Next = nil
		return p, head	
	}
	p, q := reverseListHelper(head.Next)
	q.Next = head
	head.Next = nil
	return p, head
}