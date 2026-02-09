/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    void treeToArray(TreeNode* root, vector<int> &vector){
        if(root == NULL){
            return;
        }
        vector.push_back(root->val);
        treeToArray(root->left, vector);
        treeToArray(root->right, vector);
    }
    
    int kthSmallest(TreeNode* root, int k) {
        vector<int> vector;
        priority_queue<int> pq;
        treeToArray(root, vector);
        for(int i= 0; i<k; i++){
            pq.push(vector[i]);
        }
        for(int i= k; i<vector.size(); i++){
            if(pq.top()>vector[i]){
                pq.pop();
                pq.push(vector[i]);
            }
        }
        return pq.top();
        
        
    }
};