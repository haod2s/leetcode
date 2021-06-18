use std::cmp;

impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
		let mut dp = vec![0; nums.len()];
		dp[0] = nums[0];
		let mut maxSum = dp[0];
		for i in 1..nums.len() {
			dp[i] = cmp::max(nums[i], dp[i-1] + nums[i]);
			if dp[i] > maxSum {
				maxSum = dp[i];
			}
		}
		
		maxSum
    }
}