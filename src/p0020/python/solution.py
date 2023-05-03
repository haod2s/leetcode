class Solution:
    def isValid(self, s: str) -> bool:
        st = list()
        for c in s:
            if c in ['(', '{', '[']:
                st.append(c)
            elif len(st) > 0:
                if (st[-1] == '(' and c == ')') or (st[-1] == '{' and c == '}') or (st[-1] == '[' and c == ']'):
                    st.pop()
                else:
                    return False
            else:
                return False
                    
        return len(st) == 0