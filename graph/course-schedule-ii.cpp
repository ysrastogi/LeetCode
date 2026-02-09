class Solution {
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int> adj[numCourses];
        vector<int>p;
        int count =0;
        int m = prerequisites.size();
        for(int i=0; i<m; i++){
            adj[prerequisites[i][1]].push_back(prerequisites[i][0]);
        }
        int indegree[numCourses];
        for(int i=0; i<numCourses; i++){
            indegree[i] =0;
        }
        for(int i=0; i<numCourses; i++){
            for(auto it: adj[i]){
                indegree[it]++;
            }
        }
        queue<int> q;
        for(int i=0; i<numCourses; i++){
            if(indegree[i] == 0){
                q.push(i);
            }
        }
        vector<int> ans;
        while(!q.empty()){
            int node = q.front();
            ans.push_back(node);
            count++;
            q.pop();
            for(auto it: adj[node]){
                indegree[it]--;
                if(indegree[it] == 0){
                    q.push(it);
                }
            }
        }
        if(count != numCourses) return p;
        return ans;
        
    }
};