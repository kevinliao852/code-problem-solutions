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
  ListNode *middleNode(ListNode *head) {
    vector<ListNode *> v;

    while (head != NULL) {
      v.push_back(head);
      head = head->next;
    }

    int index = floor(v.size() / 2);

    return v[index];
  }
};
