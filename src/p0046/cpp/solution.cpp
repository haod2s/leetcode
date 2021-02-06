class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> results;
        if (nums.empty()) {
            return results;
        }
        vector<int> result;
        backtrack(nums, 0, results, result);
        return results;
    }

private:
    void backtrack(vector<int>& nums, int pos, vector<vector<int>>& results, vector<int>& result) {
        if (nums.size() == pos) {
            results.push_back(result);
            return;
        }
        for (int i = pos; i < nums.size(); i++) {
            result.push_back(nums[i]);
            swap(nums[pos], nums[i]);
            backtrack(nums, pos+1, results, result);
            result.pop_back();
            swap(nums[i], nums[pos]);
        }
    }
};