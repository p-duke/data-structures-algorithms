// Solution 1
class Node {
  constructor(name) {
    this.name = name;
    this.children = [];
  }

  addChild(name) {
    this.children.push(new Node(name));
    return this;
  }

  depthFirstSearch(array) {
		array.push(this.name);
		const final = this.search(this.children, array);
		return final;
  }
	
	search(children, result) {
		for (let i = 0; i < children.length; i++) {
			let node = children[i];
			if (!node.visited) {
				node.visited = true;
				result.push(node.name);
				this.search(node.children, result);
			}
		}
		
		return result;
	}
}

// Do not edit the line below.
exports.Node = Node;
