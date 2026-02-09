class Solution {
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        int m = prerequisites.size();
        int n = numCourses;
        vector<int>adj[numCourses];
        queue<int>q;
        int count=0;
        for(int i=0; i<m; i++){
            adj[prerequisites[i][0]].push_back(prerequisites[i][1]);
        }
        int indegree[numCourses];
        for(int i=0; i<numCourses; i++){
            indegree[i] =0;
        }
        for(int i=0; i<numCourses;i++){
            for(auto it: adj[i]){
                indegree[it]++;
            }
        }
        for(int i=0; i<numCourses; i++){
            if(indegree[i] == 0){
                q.push(i);
            }
        }
        while(!q.empty()){
            int node = q.front();
            count++;
            q.pop(); 
            for(auto it: adj[node]){
                indegree[it]--;
                if(indegree[it] ==0){
                    q.push(it);
                }
            }
        }
        if(count == numCourses) return true;
        return false;

        
    }
};