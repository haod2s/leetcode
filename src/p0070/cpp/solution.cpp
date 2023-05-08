class Solution {
public:
    int climbStairs(int n) {
        if (n == 1) {
            return 1;
        } else if (n == 2) {
            return 2;
        }

        int ways = 0, i = 2, x = 1, y = 2;
        while (i++ < n) {
            ways = x + y;
            x = y;
            y = ways;
        }

        return ways;
    }
};
