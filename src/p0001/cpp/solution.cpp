class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> m;
        map<int, int>::iterator it;
        vector<int> result;
        for (int i = 0; i < nums.size(); i++) {
            it = m.find(target-nums[i]);
            if (it == m.end()) {
                m[nums[i]] = i;
            } else {
                result.push_back(it->second);
                result.push_back(i);
                break;
            }
        }
        return result;
    }
};