/**
 * Given a string s, return true if a permutation of the string could form a palindrome.
 * 
 * 
 * Example 1:
 * 
 * Input: s = "code"
 * Output: false
 * Example 2:
 * 
 * Input: s = "aab"
 * Output: true
 * Example 3:
 * 
 * Input: s = "carerac"
 * Output: true
 *  
 * 
 * Constraints:
 * 
 * 1 <= s.length <= 5000
 * s consists of only lowercase English letters
 */

class Solution {
public:
    bool canPermutePalindrome(string s) {
        unsigned int r = 0;
        for(char c: s) {
            r ^= 1 << (c - 'a');
        }

        return s.size() % 2 ? (r & (r-1)) == 0 : r == 0;
    }
};