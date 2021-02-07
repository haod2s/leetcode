/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* sortedListToBST(ListNode* head) {
        if (head == nullptr) {
            return nullptr;
        }
        int n = 1;
        ListNode* cur = head;
        while (cur->next) {
            cur = cur->next;
            n++;
        }
        
        return sortedListToBSTHelper(head, n);
    }

private:
    TreeNode* sortedListToBSTHelper(ListNode* head, int length) {
        if (head == nullptr || length <= 0) {
            return nullptr;
        }
        ListNode* cur = head;
        int leftLen = length >> 1;
        while (cur->next && leftLen--) {
            cur = cur->next;
        }
        TreeNode* left = sortedListToBSTHelper(head, length >> 1);
        TreeNode* root = new TreeNode(cur->val);
        TreeNode* right = sortedListToBSTHelper(cur->next, length-1-(length >> 1));
        root->left = left;
        root->right = right;

        return root;
    }
};