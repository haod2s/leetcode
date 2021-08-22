/**
 * Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.
 * 
 * Example 1:
 * 
 * Input: l1 = [1,2,4], l2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 * Example 2:
 * 
 * Input: l1 = [], l2 = []
 * Output: []
 * Example 3:
 * 
 * Input: l1 = [], l2 = [0]
 * Output: [0]
 *  
 * 
 * Constraints:
 * 
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both l1 and l2 are sorted in non-decreasing order
 */


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var ans, p *ListNode
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p == nil {
			ans = &ListNode{}
			p = ans
		} else {
			p.Next = &ListNode{}
			p = p.Next
		}
		if p1.Val > p2.Val {
			p.Val = p2.Val
			p2 = p2.Next
		} else {
			p.Val = p1.Val
			p1 = p1.Next
		}
	}
	for p1 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = p1.Val
		p1 = p1.Next
	}
	for p2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = p2.Val
		p2 = p2.Next
	}
	return ans
}