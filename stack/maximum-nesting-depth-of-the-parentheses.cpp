class Solution {
public:
    int maxDepth(string s) {
        stack<char> st;
        int maxi = 0;
        for(int i=0; i<s.size(); i++){
            if(s[i] == '(')
                st.push(s[i]);
            
            else if(s[i] == ')'){
                if(!st.empty()){
                    int curr = st.size();
                    maxi = max(curr, maxi);
                    st.pop();
                }
            }
        }
        if(!st.empty()){
            return -1;
        }
        return maxi;
        
    }
};