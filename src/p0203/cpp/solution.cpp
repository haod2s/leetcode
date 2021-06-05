/**
 * 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
 *  
 * 输入：head = [1,2,6,3,4,5,6], val = 6
 * 输出：[1,2,3,4,5]
 * 示例 2：
 * 
 * 输入：head = [], val = 1
 * 输出：[]
 * 示例 3：
 * 
 * 输入：head = [7,7,7,7], val = 7
 * 输出：[]
 *  
 * 
 * 提示：
 * 
 * 列表中的节点在范围 [0, 104] 内
 * 1 <= Node.val <= 50
 * 0 <= k <= 50
 */
 
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        if(head == NULL) {
            return NULL;
        }
        vector<int> v;
        ListNode* p = head;
        ListNode* pre = NULL;
        while(p != NULL) {
            if(p->val != val) {
                v.push_back(p->val);
            }
            p = p->next;
        }
        if(v.size() == 0) {
            return NULL;
        }
        p = head;
        for(int e : v) {
            p->val = e;
            pre = p;
            p = p->next;
        }
        pre->next = NULL;

        return head;
    }
};