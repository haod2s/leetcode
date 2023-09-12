class Solution {
public:
    bool validateBinaryTreeNodes(int n, vector<int>& leftChild, vector<int>& rightChild) {
		vector<int> in(n, 0);
		for (int i = 0; i < n; i++) {
			if (leftChild[i] != -1) {
				in[leftChild[i]]++;
			}
			if (rightChild[i] != -1) {
				in[rightChild[i]]++;
			}
		}
		int root = -1;
		for (int i = 0; i < n; i++) {
			if (in[i] >= 2 || (in[i] == 0 && root != -1)) {
				return false;
			} else if (in[i] == 0) {
				root = i;
			}
		}
		if (root == -1) {
			return false;
		}
		vector<bool> vis(n, false);
		queue<int> q;
		q.emplace(root);
		while (!q.empty()) {
			int idx = q.front();
			if (vis[idx]) {
				return false;
			}
			vis[idx] = true;
			q.pop();
			if (leftChild[idx] != -1) {
				q.emplace(leftChild[idx]);
			}
			if (rightChild[idx] != -1) {
				q.emplace(rightChild[idx]);
			}
		}
		for (auto v : vis) {
			if (!v) {
				return false;
			}
		}
		return true;
    }
};