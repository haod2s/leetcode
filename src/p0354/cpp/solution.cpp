//给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如
//同俄罗斯套娃一样。
//
// 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 说明:
//不允许旋转信封。
//
// 示例:
//
// 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出: 3
//解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//
// Related Topics 二分查找 动态规划
// 👍 269 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    int maxEnvelopes(vector<vector<int>>& envelopes) {
        sort(envelopes.begin(), envelopes.end(), less);
        vector<int> nums(envelopes.size());
        for (int i = 0; i < envelopes.size(); i++) {
            nums[i] = envelopes[i][1];
        }
        return lengthOfLIS(nums);
    }

private:
    static bool less(const vector<int>& lhs, const vector<int>& rhs) {
        if (lhs[0] == rhs[0]) {
            return lhs[1] > rhs[1];
        }
        return lhs[0] < rhs[0];
    }

    int lengthOfLIS(vector<int>& nums) {
        vector<int> dp(nums.size(), 1);
        for(int i = 0; i < nums.size(); i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = max(dp[i], dp[j]+1);
                }
            }
        }
        auto it = max_element(dp.begin(), dp.end());
        return it == dp.end() ? 0 : *it;
    }
};
//leetcode submit region end(Prohibit modification and deletion)
