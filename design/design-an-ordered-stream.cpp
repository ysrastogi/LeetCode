class OrderedStream {
public:
    vector<string> val;
    int ptr = 0;
    OrderedStream(int n) {
        val.resize(n+1, "");
    }
    
    vector<string> insert(int idKey, string value) {
        vector<string> ans;
        val[idKey-1]=value;
        while(val[ptr] != ""){
            ans.push_back(val[ptr]);
            ptr++;
        }
        return ans;
        
    }
};

/**
 * Your OrderedStream object will be instantiated and called as such:
 * OrderedStream* obj = new OrderedStream(n);
 * vector<string> param_1 = obj->insert(idKey,value);
 */