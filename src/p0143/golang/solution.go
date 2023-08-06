/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	a := make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		a = append(a, p)
	}
	p := head
	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			p.Next = a[i]
			p = a[i]
			break
		}
		if i != 0 {
			p.Next = a[i]
		}
		a[i].Next = a[j]
		p = a[j]
	}
	p.Next = nil
}