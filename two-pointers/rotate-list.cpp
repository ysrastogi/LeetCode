/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
        int count = 1;
        ListNode* tmp = head;
        if(head == NULL){
            return head;
        }
        while(tmp->next != NULL){
            tmp = tmp->next;
            count++;
        }
        tmp->next = head;
        int n = count -(k%count);
        tmp = head;
        for(int i=1; i<n; i++){
            tmp = tmp->next;
        }
        head = tmp->next;
        tmp->next = NULL;
        return head;
    }
};