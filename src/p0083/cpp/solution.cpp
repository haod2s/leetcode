class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if (head == NULL) {
            return head;
        }

        ListNode* cur = head;
        ListNode* next = head->next;

        while (next != NULL) {
            if (cur->val == next->val) {
                cur->next = next->next;
                delete next;
            } else {
                cur = cur->next;
            }

            next = cur->next;
        }

        return head;
    }
};