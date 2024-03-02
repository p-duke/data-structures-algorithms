package main

import (
  "fmt"
  "reflect"
) 

// Educative 
// The time complexity of the algorithm is O(n) , where n is the number of nodes in the linked list.
// The space complexity of the algorithm above is O(1)
func detectCycle(head *EduLinkedListNode) bool {
    if head == nil {
        return false
    }

    // Initialize two pointers, slow and fast, to the head of the linked list
    slow, fast := head, head

    // Run the loop until we reach the end of the linked list or find a cycle
    for fast != nil && fast.next != nil {
        // Move the slow pointer one step at a time
        slow = slow.next
        // Move the fast pointer two steps at a time
        fast = fast.next.next

        // If there is a cycle, the slow and fast pointers will meet
        if slow == fast {
            return true
        }
    }

    // If we reach the end of the linked list and haven't found a cycle, return false
    return false
}

// O(n) time | O(1) space
func detectCycle(head *EduLinkedListNode) bool {
  var slow *EduLinkedListNode = head
  var fast *EduLinkedListNode = head.next

  for fast != nil {
    if slow == fast {
      return true
    }

    slow = slow.next
    fast = fast.next.next
  }
	return false
}

func main() {
	tests := []struct {
		linkedList []int
    tailPointer int
		want   bool
	}{
    {linkedList: []int{2, 4, 6, 8, 10}, tailPointer: 2, want: true},
    {linkedList: []int{1, 3, 5, 7, 9}, tailPointer: -1, want: false},
    {linkedList: []int{1, 2, 3, 4, 5}, tailPointer: 3, want: true},
    {linkedList: []int{0, 2, 3, 5, 6}, tailPointer: -1, want: false},
    {linkedList: []int{3, 6, 9, 10, 11}, tailPointer: 0, want: true},
    {linkedList: []int{4, 4, 4, 4, 4, 4, 4}, tailPointer: -1, want: false},
	}

	for _, tc := range tests {
    eduLinkedList := &EduLinkedList{}
    eduLinkedList.CreateLinkedList(tc.linkedList)
    eduLinkedList.DisplayLinkedList()

    if (tc.tailPointer > -1) {
      var tailNode = eduLinkedList.GetNode(eduLinkedList.head, len(tc.linkedList) - 1)
      var tailNodeNext = eduLinkedList.GetNode(eduLinkedList.head, tc.tailPointer)
      tailNode.next = tailNodeNext
    }

    got := detectCycle(eduLinkedList.head)
    if !reflect.DeepEqual(tc.want, got) {
      fmt.Printf("FAIL! Assert: %d, Expected: %t --- Got: %t\n", tc.linkedList, tc.want, got)
    } else {
      fmt.Printf("PASS Assert: %d, Expected: %t --- Got: %t\n", tc.linkedList, tc.want, got)
    }
	}
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    fast := head
    slow := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
