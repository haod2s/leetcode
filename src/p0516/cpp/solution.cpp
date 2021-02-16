//给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
//
//
//
// 示例 1:
//输入:
//
// "bbbab"
//
//
// 输出:
//
// 4
//
//
// 一个可能的最长回文子序列为 "bbbb"。
//
// 示例 2:
//输入:
//
// "cbbd"
//
//
// 输出:
//
// 2
//
//
// 一个可能的最长回文子序列为 "bb"。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 只包含小写英文字母
//
// Related Topics 动态规划
// 👍 386 👎 0


class Solution {
public:
    int longestPalindromeSubseq(string s) {
        vector<vector<int>> dp(s.size(), vector<int>(s.size()));
        for (int i = 0; i < dp.size(); i++) {
            dp[i][i] = 1;
        }
        for (int i = s.size() - 2; i >= 0; i--) {
            for (int j = i + 1; j < s.size(); j++) {
                dp[i][j] = s[i] == s[j] ? dp[i+1][j-1] + 2 : max(dp[i][j-1], dp[i+1][j]);
            }
        }
        return dp[0][s.size()-1];
    }
};
