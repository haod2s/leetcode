// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
//
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
// 示例 2：
//
// 输入：head = [5], left = 1, right = 1
// 输出：[5]
//
//
// 提示：
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
//
// 通过次数143,173提交次数265,530

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	s := make([]int, 0)
	p, i := head, 0
	left, right = left-1, right-1
	for p != nil && i <= right {
		if i >= left {
			s = append(s, p.Val)
		}
		p = p.Next
		i++
	}
	p, i = head, 0
	for p != nil && i <= right {
		if i >= left {
			p.Val = s[len(s)-1]
			s = s[:len(s)-1]
		}
		p = p.Next
		i++
	}

	return head
}