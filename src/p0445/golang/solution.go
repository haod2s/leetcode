/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	p1, p2 := l1, l2
	for p1 != nil {
		s1 = append(s1, p1.Val)
		p1 = p1.Next
	}
	for p2 != nil {
		s2 = append(s2, p2.Val)
		p2 = p2.Next
	}
	var ans *ListNode
	flag := 0
	i, j := len(s1)-1, len(s2)-1
	for i >= 0 || j >= 0 || flag != 0 {
		val := 0
		if i >= 0 {
			val += s1[i]
			i--
		}
		if j >= 0 {
			val += s2[j]
			j--
		}
		val += flag
		flag = val / 10
		val %= 10
		cur := &ListNode{
			Val:  val,
			Next: ans,
		}
		ans = cur
	}

	return ans
}