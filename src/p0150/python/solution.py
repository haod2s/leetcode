class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        st = []
        for v in tokens:
            if v not in "+-*/":
                st.append(int(v))
            else:
                b = st.pop()
                a = st.pop()
                if v == '+': st.append(a+b)
                if v == '-': st.append(a-b)
                if v == '*': st.append(a*b)
                if v == '/': st.append((-1*(abs(a)//abs(b))) if a*b < 0 else (a//b))
        
        return st[-1]