function sortedSquaredArray(array) {
	let final = [];
	for (let i = 0; i < array.length; i++) {
	  const num = array[i];
		let squared = num * num;
		if (final.length === 0) {
		  final.push(squared);
			continue;
		}
		
		for (let y = 0; y < final.length; y++) {
		  let curr = final[y];
			let next = final[y + 1];
		  let fIndex = final.indexOf(curr);
			
			if (squared <= curr) {
				final.splice(fIndex, 0, squared);
				break;
			} else if (next === undefined) {
				final.push(squared);
				break;
			} 
		}
	}
	
  return final;
}

// Do not edit the line below.
exports.sortedSquaredArray = sortedSquaredArray;
