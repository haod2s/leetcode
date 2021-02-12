//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。 
//
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。 
//
// 说明： 
//
// 
// 字母异位词指字母相同，但排列不同的字符串。 
// 不考虑答案输出的顺序。 
// 
//
// 示例 1: 
//
// 
//输入:
//s: "cbaebabacd" p: "abc"
//
//输出:
//[0, 6]
//
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
// 
//
// 示例 2: 
//
// 
//输入:
//s: "abab" p: "ab"
//
//输出:
//[0, 1, 2]
//
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
// 
// Related Topics 哈希表 
// 👍 463 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        vector<int> r;
        unordered_map<char, int> window, need;
        for (const char& c: p) {
            need[c]++;
        }
        int left = 0, right = 0, valid = 0;
        while (right < s.size()) {
            char c = s[right++];
            if (need.count(c) != 0) {
                window[c]++;
                if (window[c] == need[c]) {
                    valid++;
                }
            }
            while (right - left + 1 > p.size()) {
                if (valid == need.size()) {
                    r.push_back(left);
                }
                char d = s[left++];
                if (need.count(d) != 0) {
                    if (need[d] == window[d]) {
                        valid--;
                    }
                    window[d]--;
                }
            }
        }

        return r;
    }
};
//leetcode submit region end(Prohibit modification and deletion)
