class Solution {
public:
    string longestPalindrome(string s) {
		int n = s.size();
		vector<vector<bool>> dp(n, vector<bool>(n));
		for(int i = 0; i < n; i++) {
			dp[i][i] = true;
		}
		int x = 0, y = 0, maxSize = INT_MIN;
		for(int i = n - 2; i >= 0; --i) {
			for(int j = i+1; j < n; ++j) {
                if(s[i] == s[j]) {
                   dp[i][j] = dp[i+1][j-1]; 
                }
				if(dp[i][j]) {
                    maxSize = max(j-i+1, maxSize);
                }
			}
		}

		return s.substr(x, maxSize);
    }
};