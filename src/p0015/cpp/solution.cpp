/**
 * Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
 * 
 * Notice that the solution set must not contain duplicate triplets.
 * 
 *  
 * 
 * Example 1:
 * 
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Example 2:
 * 
 * Input: nums = []
 * Output: []
 * Example 3:
 * 
 * Input: nums = [0]
 * Output: []
 *  
 * 
 * Constraints:
 * 
 * 0 <= nums.length <= 3000
 * -105 <= nums[i] <= 105
 */


class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
		sort(nums.begin(), nums.end());
		vector<vector<int>> ret;
		for(int i = 0; i < nums.size(); i++) {
			if(i > 0 && nums[i] == nums[i-1]) { continue; }
			for(int j = i+1, k=nums.size()-1; j < nums.size(); j++) {
				if(j > i+1 && nums[j] == nums[j-1]) { continue; }
				while(k > j && nums[j]+nums[k] > -1*nums[i]) { k--; }
				if(k == j) { break; }
				if(nums[i] + nums[j] + nums[k] == 0) {
					ret.push_back({nums[i], nums[j], nums[k]});
				}
			}
		}

		return ret;
    }
};