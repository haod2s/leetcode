/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	p1, p2 := l1, l2
	var ans, p *ListNode
	for p1 != nil || p2 != nil {
		if p == nil {
			ans = &ListNode{}
			p = ans
		} else {
			p.Next = &ListNode{}
			p = p.Next
		}
		if p1 != nil {
			p.Val += p1.Val
		}
		if p2 != nil {
			p.Val += p2.Val
		}
		p.Val += carry
		carry = p.Val / 10
		p.Val %= 10
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if carry != 0 {
		p.Next = &ListNode{
			Val: 1,
		}
	}
	return ans
}