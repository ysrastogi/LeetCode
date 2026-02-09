class Solution {
public:
    int myAtoi(string s) {
        int i=0;
        while(i<s.size() && s[i]== ' '){
            i++;
        }
        bool positive = false; bool negative = false;
        if(s[i] == '+'){
            positive = true;
            i++;
        }
        else if(s[i] == '-'){
            negative = true;
            i++;
        }
            
        double ans =0;
        while(i<s.size() && s[i]>= '0' && s[i]<= '9'){
            ans = ans*10+(s[i] - '0');
            i++;
        }
        if(negative)
            ans = -ans;
        
        if(ans>INT_MAX)
            ans = INT_MAX;
        
        if(ans<INT_MIN)
            ans = INT_MIN;

        return ans;
    }
};