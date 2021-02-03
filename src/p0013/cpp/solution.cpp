class Solution {
public:
    int romanToInt(string s) {
        int result = 0, i = 0;
        while (i < s.length()-1) {
            if (toNumber(s[i]) < toNumber(s[i+1])) {
                result += toNumber(s[i+1]) - toNumber(s[i]);
                i += 2;
                continue;
            }
            result += toNumber(s[i++]);
        }
        return result + (i == s.length()-1 ? toNumber(s[s.length()-1]) : 0);
    }

private:
    int toNumber(char c) {
        switch (c) {
            case 'I':
                return 1;
            case 'V':
                return 5;
            case 'X':
                return 10;
            case 'L':
                return 50;
            case 'C':
                return 100;
            case 'D':
                return 500;
            case 'M':
                return 1000;
        }
        return 0;
    }
};