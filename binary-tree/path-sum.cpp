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
    bool pathSum(TreeNode* node, int target){
        if(node == NULL){
           return false;
        }
        target -= node->val;
        if(node->left== NULL && node->right == NULL){
            if(target == 0){
                return true;
            }
            return false;
        }
        // if(target == 0) return true;
        // if(target<0)return false;
        bool left = pathSum(node->left, target);
        bool right = pathSum(node->right, target);

        return left||right;
    }
    bool hasPathSum(TreeNode* root, int targetSum) {
        return pathSum(root, targetSum);
    }
};