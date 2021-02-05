class Solution {
public:
    vector<string> letterCombinations(string digits) {
        vector<string> results;
        if (digits.empty()) {
            return results;
        }
        unordered_map<char, string> m{
            {'2', "abc"},
            {'3', "def"},
            {'4', "ghi"},
            {'5', "jkl"},
            {'6', "mno"},
            {'7', "pqrs"},
            {'8', "tuv"},
            {'9', "wxyz"}
        };
        string result;
        backtrack(results, m, digits, 0, result);
        return results;
    }

private:
    void backtrack(vector<string>& results, const unordered_map<char, string>& m, const string& digits, int i, string& result) {
        if (digits.size() == i) {
            results.push_back(result);
            return;
        }
        char digit = digits[i];
        const string& letters = m.at(digit);
        for (const char& letter: letters) {
            result.push_back(letter);
            backtrack(results, m, digits, i+1, result);
            result.pop_back();
        }
    }
};