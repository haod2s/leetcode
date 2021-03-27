// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
// 
// 返回同样按升序排列的结果链表。
// 
//  
// 
// 示例 1：
// 
// 
// 输入：head = [1,1,2]
// 输出：[1,2]
// 示例 2：
// 
// 
// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]
//  
// 
// 提示：
// 
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
// 通过次数234,197提交次数436,633

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	s := make([]int, 0)
	for p != nil {
		if len(s) == 0 || s[len(s)-1] != p.Val {
			s = append(s, p.Val)
		}
		p = p.Next
	}
	p = head
	for i := 0; i < len(s); i++ {
		p.Val = s[i]
		if i == len(s)-1 {
			p.Next = nil
		} else {
			p = p.Next
		}
	}

	return head
}