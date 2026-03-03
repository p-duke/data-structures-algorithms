# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        """
        Solution
        - we iterate through the linked list
        - use a fast slow pointer technique
        - declar slow as the head of the list
        - declare fast as head.next.next
        - if slow == fast then return true
        - else slow = slow.next and fast = fast.next.next
        - return false at the end
        """
        slow = self
        print("slow", self)
        while slow != None:
            slow = slow.next

        return False
