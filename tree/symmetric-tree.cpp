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
    bool checkSymmetry(TreeNode* p , TreeNode* q){
        bool preorder= false;
        if(p== NULL && q == NULL) return true;
        else if(p == NULL || q == NULL) return false;
        if(p->val == q->val) preorder = true;
        bool leftRight = checkSymmetry(p->left, q->right);
        bool rightLeft = checkSymmetry(p->right, q->left);
        return (preorder && leftRight && rightLeft);
    }
    bool isSymmetric(TreeNode* root) {
        if(!root) return true;
        return checkSymmetry(root->left, root->right);
    }
};