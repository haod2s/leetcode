class Solution {
public:
    bool containsPattern(vector<int>& arr, int m, int k) {
		for (int start = 0; start < arr.size()-m; start++) {
			if (dfs(arr, start, start, m, k)) {
				return true;
			}
		}
		return false;
    }

	bool dfs(vector<int>& arr, int pBegin, int start, int m, int k) {
		if (start+m > arr.size() && k > 0) {
			return false;
		}
        if (k == 0) {
            return true;
        }
		for (int i = 0; i < m; i++) {
			if (arr[start+i] != arr[pBegin + i]) {
				return false;
			}
		}
		return dfs(arr, pBegin, start+m, m, k-1);
	}
};