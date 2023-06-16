// Solution 2

// Solution 1
// This is an input class. Do not edit.
class LinkedList {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

function removeDuplicatesFromLinkedList(linkedList) {
	const values = [];
	const list = [];
	let node = linkedList;
	let prev;
	while (node.next !== null) {
		let duplicate = values.includes(node.value);
		if (duplicate) {
			prev.next = node.next;
		} else {
			values.push(node.value);
			prev = node;
		}
		
		node = node.next
		duplicate = values.includes(node.value);
		
		if (duplicate && node.next === null) {
			prev.next = null;
		}
	}
  return linkedList;
}

// Do not edit the lines below.
exports.LinkedList = LinkedList;
exports.removeDuplicatesFromLinkedList = removeDuplicatesFromLinkedList;

