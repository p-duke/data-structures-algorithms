package main

import (
  "fmt"
  "reflect"
)

// Educative solution
// O(n) time | O(1) space
func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {
	// Initially slow and fast pointers point to head
	slow := head
	fast := head

	// Traverse the linked list until fast reaches at the last node or NULL
	for fast != nil && fast.next != nil {
		// Move slow pointer one step ahead
		slow = slow.next

		// Move fast pointer two step ahead
		fast = fast.next.next
	}

	// Return the slow pointer
	return slow
}

// getMiddleNode function to find the middle node of the linked list
// O(n) time | O(1) space
func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {
  var slow, fast *EduLinkedListNode = head, head

  for fast.next != nil {
    slow = slow.next

    if fast.next.next == nil {
      return slow
    }

    fast = fast.next.next
  }

	return slow
}

func main() {
	tests := []struct {
		linkedList []int
		want   int
	}{
    {linkedList: []int{1,2,3,4,5}, want: 3},
    {linkedList: []int{1,2,3,4,5,6}, want: 4},
    {linkedList: []int{3,2,1}, want: 2},
    {linkedList: []int{10}, want: 10},
    {linkedList: []int{1,2}, want: 2},
	}

	for _, tc := range tests {
    eduLinkedList := &EduLinkedList{}
    eduLinkedList.CreateLinkedList(tc.linkedList)
    var expected *EduLinkedListNode = InitLinkedListNode(tc.want)

    got := getMiddleNode(eduLinkedList.head)
    if !reflect.DeepEqual(expected.data, got.data) {
      fmt.Printf("FAIL! Assert: %d, Expected: %d --- Got: %d\n", tc.linkedList, expected.data, got.data)
    } else {
      fmt.Printf("PASS Assert: %d, Expected: %d --- Got: %d\n", tc.linkedList, expected.data, got.data)
    }
	}
}
