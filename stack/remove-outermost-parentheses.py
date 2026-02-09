class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        result = []
        st = []

        for char in s:
            if char == '(':
                if st:  # if stack is not empty, it's not outermost
                    result.append('(')
                st.append('(')
            else:  # char == ')'
                st.pop()
                if st:  # if stack not empty after popping, it's not outermost
                    result.append(')')

        return "".join(result)
