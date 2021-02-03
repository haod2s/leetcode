class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        stringstream ss;
        ss << x;
        string s = ss.str();
        int i = 0;
        int j = s.length()-1;
        while (i <= j) {
            if (s[i++] != s[j--]) {
                return false;
            }
        }
        return true;
    }
};