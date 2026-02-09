class Solution {
public:
    vector<vector<int>> groupThePeople(vector<int>& groupSizes) {
        map<int,vector<int>> m;
        vector<vector<int>> vt1;
        for(int i=0; i<groupSizes.size(); i++){
            m[groupSizes[i]].push_back(i);
        }
        map<int, vector<int>>::iterator it = m.begin();
        while(it!= m.end()){
            int i=0;
            while( i<it->second.size()){
                int count = 0;
                vector<int> v;
                while(count!= it->first){
                    v.push_back(it->second[i]);
                    count++;
                    i++;
                }
                vt1.push_back(v);
            }
            it++;
        }
        return vt1;
        
    }
};