class Solution:
    def buddyStrings(self, s: str, goal: str) -> bool:
        if len(s) != len(goal) or len(s) == 1:
            return False
        if s == goal:
            if len(set(s)) < len(s):
                return True
        diff = [(x,y) for x, y in zip(s, goal) if x != y]

        return len(diff) == 2 and diff[0][0] == diff[1][1] and diff[0][1] == diff[1][0]