class Solution {
public:
    static bool cmp(pair<char, int>a, pair<char, int>b){
        return a.second>b.second;
    }
    vector<pair<char,int>> sort_map(map<char, int> &M){
        vector<pair<char,int>> A;
        for(auto& it : M){
            A.push_back(it);
        }
        sort(A.begin(), A.end(), cmp);
        return A;
    }
    string frequencySort(string s) {
        map<char, int>mp;
        vector<pair<char,int>> A;
        string ans;
        for(int i=0; i<s.size(); i++){
            mp[s[i]]++;
        }
        A = sort_map(mp);
        for(auto& it: A){
            cout<<it.first<<it.second<<endl;
        }
        for(auto& it :A){
            while(it.second){
                ans+= it.first;
                it.second--;
            }
        }
        return ans;

        
    }
};