/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
  // Solution
  // fast and slow pointer
  // declare two variables for each
  // slow will be head
  // fast will be head.next.next
  // loop through the linked list while fast pointer next is not null
  // solve for simple case first
  // if head.next == null return false
  if (head == null || head.next === null) {
    return false;
  }

  let slow = head;
  let fast = head.next.next;

  while (fast !== null && fast.next !== null) {
    if (slow === fast) {
      return true;
    }

    if (fast.next != null) {
      fast = fast.next.next;
    }

    slow = slow.next;
  }

  return false;
};
