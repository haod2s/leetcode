/**
 * Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
 * 
 * Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
 * 
 *  
 * 
 * Example 1:
 * 
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 * Example 2:
 * 
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *  
 * 
 * Constraints:
 * 
 * 1 <= num1.length, num2.length <= 200
 * num1 and num2 consist of digits only.
 * Both num1 and num2 do not contain any leading zero, except the number 0 itself.
 */

#include <algorithm>

class Solution {
public:
    string multiply(string num1, string num2) {
		string ans = "0";
		for(int i = num2.size()-1; i >= 0; --i) {
			for(int j = num1.size()-1; j >= 0; --j) {
				int t = (num2[i] - '0') * (num1[j] - '0');
				string n = "";
				if(t / 10) { n.append(1, (t / 10) + '0'); }
				n.append(1, (t % 10) + '0');
				if(n != "0") {
					n.append(num1.size()+num2.size()-j-i-2, '0');
				}
				ans = add(ans, n);
			}

		}

		return ans;
    }

private:
    string add(const string& num1, const string& num2) {
		string ans = "";
		int flag = 0;
		int i = num1.size() - 1, j = num2.size() - 1;
		while(i >= 0 || j >= 0) {
			int t = 0;
			if(i >= 0) { t += num1[i--] - '0'; }
			if(j >= 0) { t += num2[j--] - '0'; }
			t += flag;
			flag = t / 10 ? 1 : 0;
			ans.append(1, t % 10 + '0');
		}
		if(flag) { ans.append(1, '1'); }
		reverse(ans.begin(), ans.end());

		return ans;
	}
};