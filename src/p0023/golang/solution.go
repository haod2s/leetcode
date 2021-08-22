/**
 * You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
 * 
 * Merge all the linked-lists into one sorted linked-list and return it.
 * 
 *  
 * 
 * Example 1:
 * 
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 *   1->4->5,
 *   1->3->4,
 *   2->6
 * ]
 * merging them into one sorted list:
 * 1->1->2->3->4->4->5->6
 * Example 2:
 * 
 * Input: lists = []
 * Output: []
 * Example 3:
 * 
 * Input: lists = [[]]
 * Output: []
 *  
 * 
 * Constraints:
 * 
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length won't exceed 10^4
 */

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}

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