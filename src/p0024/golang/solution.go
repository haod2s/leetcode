/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	q := swapPairs(p2.Next)
	p1.Next = q
	p2.Next = p1

	return p2
}

// 腾讯