// Better O(n^2) time | O(n) space``
function threeNumberSum(array, targetSum) {
  const final = [];
  const sorted = array.sort((a,b) => a - b);
  for (let i = 0; i < sorted.length - 1; i++) {
    let left = i + 1;
    let right = sorted.length - 1;

    while (left < right) {
      const currSum = sorted[i] + sorted[left] + sorted[right];
      if (currSum === targetSum) {
        final.push([sorted[i], sorted[left], sorted[right]]);
        left++;
        right--;
      } else if (currSum > targetSum) {
        right--;
      } else if (currSum < targetSum) {
        left++;
      }
    }
  }

  return final;
}

// Good
function threeNumberSum(array, targetSum) {
  const final = [];
  const sorted = array.sort((a,b) => a - b);
  let curr = 0;
  let left = curr + 1;
  let right = sorted.length - 1;

  while (curr < sorted.length - 1) {
    let sum = sorted[curr] + sorted[left] + sorted[right];
    if (left === right) {
      curr += 1;
      left = curr + 1;
      right = sorted.length - 1;

    } else if (sum === targetSum) {
      let temp = [sorted[curr], sorted[left], sorted[right]].sort((a,b) => a - b);
      final.push(temp);
      left += 1;
      right = sorted.length - 1;
    } else if (sum > targetSum) {
      right -= 1;
    } else if (sum < targetSum) {
      left += 1;
    }
  }

  return final;
}
// Bad
function threeNumberSum(array, targetSum) {
	const sortedArr = array.sort((a,b) => a - b);
	let final = [];
	for (let i = 0; i < sortedArr.length; i++) {
	  let curr = sortedArr[i];
		
		for (let y = 0; y < sortedArr.length; y++) {
		  let next = sortedArr[y];
		  if (curr === next) continue; 
			let third = targetSum - (curr + next);
			if (sortedArr.includes(third)) {
				let triplet;
				if (curr < next && curr < third && next < third) {
				  triplet = [curr, next, third];
				  if (!subArray(final, triplet)) {
					  final.push(triplet);
					}
				} else if (next < curr && next < third && curr < third) {
				  triplet = [next, curr, third];
				  if (!subArray(final, triplet)) {
					  final.push(triplet);
					}
				} else if (third < next && third < curr && next < curr) {
				  triplet = [third, next, curr];
				  if (!subArray(final, triplet)) {
						final.push(triplet);
					}
				}
			}
		}
	}
	
	return final;
}

const subArray = function(main, sub) {
  let included = false;
  for (let i = 0; i < main.length; i++) {
	  let subArr = main[i];
		let one = subArr.includes(sub[0]);
		let two = subArr.includes(sub[1]);
		let three = subArr.includes(sub[2]);
		
		if (one && two && three) {
		  included = true;
			break;
		}
	}
	
	return included;
}


// Do not edit the line below.
exports.threeNumberSum = threeNumberSum;
