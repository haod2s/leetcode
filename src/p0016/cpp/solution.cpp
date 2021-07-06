/**
 * Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
 * 
 *  
 * 
 * Example 1:
 * 
 * Input: nums = [-1,2,1,-4], target = 1
 * Output: 2
 * Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *  
 * 
 * Constraints:
 * 
 * 3 <= nums.length <= 10^3
 * -10^3 <= nums[i] <= 10^3
 * -10^4 <= target <= 10^4
 */

class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
		int best = 1e5;
		sort(nums.begin(), nums.end());
		for(int i = 0; i < nums.size(); i++) {
			if(i > 0 && nums[i] == nums[i-1]) { continue; }
			int j = i+1, k = nums.size()-1;
			while(j < k) {
				int sum = nums[i] + nums[j] + nums[k];
				if(sum == target) {
					return sum;
				}
				update(sum, target, best);
				if(sum > target) {
					int k_ = k - 1;
					while(k_ > j && nums[k_] == nums[k]) { --k_; }
					k = k_;
				} else {
					int j_ = j + 1;
					while(j_ < k && nums[j_] == nums[j]) { ++j_; }
					j = j_;
				}				
			}
		}

		return best;
    }

private:
    void update(int sum, int target, int& best) {
		if(abs(sum - target) < abs(best - target)) {
			best = sum;
		}
	}
    
};