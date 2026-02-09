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
    bool isSameTree(TreeNode* p, TreeNode* q) {
        bool postorder, inorder = false;
        if(p== NULL && q== NULL) return true;
        else if(p == NULL || q== NULL) return false;
        
        if(p->val == q->val) postorder = true;
        bool left = isSameTree(p->left, q->left);
        if(p->val == q->val) inorder = true;
        bool right = isSameTree(p->right, q->right);
        return (postorder && inorder && left && right);
        
    }
};