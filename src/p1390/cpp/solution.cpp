class Solution {
public:
    int sumFourDivisors(vector<int>& nums) {
		int ans = 0;
		for (auto e : nums) {
			map<int, int> m;
			for (int i = 1; i <= e; i++) {
				if (e % i == 0) {
					if (m.count(i) > 0) {
						break;
					}
					m.insert({e/i, (e/i==i ? 0 : i)});
					if (m.size() > 2) {
						break;
					}
				}	
			}
			if (m.size() > 2) {
				continue;
			}
			if (m.size() == 2) {
				int t = 0;
				for (auto it = m.begin(); it != m.end(); ++it) {
					if (it->second != 0) {
						t += it->first + it->second;
					} else {
						t = 0;
						break;
					}
				}
				ans += t;
			}
		}
		return ans;
    }
};