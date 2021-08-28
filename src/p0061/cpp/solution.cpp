class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
		if(!head || !head->next || !k) {
			return head;
		}
		ListNode* p = head;
		int n = 1;
		while(p->next) {
			n++;
			p = p->next;
		}
		p->next = head;
		k = n - k % n;
		while(k--) {
			p = p->next;
		}
		ListNode* h = p->next;
		p->next = nullptr;
		return h;
    }
};