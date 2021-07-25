/**
 * Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
 * 
 *  
 * 
 * Example 1:
 * 
 * Input: s = "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 * Example 2:
 * 
 * Input: s = "God Ding"
 * Output: "doG gniD"
 *  
 * 
 * Constraints:
 * 
 * 1 <= s.length <= 5 * 104
 * s contains printable ASCII characters.
 * s does not contain any leading or trailing spaces.
 * There is at least one word in s.
 * All the words in s are separated by a single space
 */

class Solution {
public:
    string reverseWords(string s) {
		int i = 0;
		while(1) {
			int p = s.find_first_of(' ', i);
			if(p == string::npos) {
				reverseSubWords(s, i, s.size()-1);
				break;
			}
			reverseSubWords(s, i, p-1);
			i = p + 1;
		}

		return s;
    }

private:
    void reverseSubWords(string&s, int i, int j) {
		int m = i, n = j;
		while(m < n) {
			char t = s[m];
			s[m] = s[n];
			s[n] = t;
			++m, --n;
		}
	}
};