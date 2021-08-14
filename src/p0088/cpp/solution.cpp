class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
		vector<int> t(m);
		for(int i = 0; i < m; i++) {
			t[i] = nums1[i];
		}
		int i = 0, j = 0, k = 0;
		while(i < m && j < n) {
			nums1[k++] = t[i] > nums2[j] ? nums2[j++] : t[i++];
		}
		while(i < m) { nums1[k++] = t[i++]; }
		while(j < n) { nums1[k++] = nums2[j++]; }
    }
};