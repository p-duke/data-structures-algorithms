// Best - O(n log (n)) time | O(1) space
function twoNumberSum(array, targetSum) {
  const sorted = array.sort((a,b) => a - b);
  let left = 0;
  let right = sorted.length - 1;

  while (left !== right) {
    let sum = sorted[left] + sorted[right];
    if (sum === targetSum) {
      return [sorted[left], sorted[right]];
    } else if (sum > targetSum) {
      right -= 1;
    } else if (sum < targetSum) {
      left += 1;
    }
  }
  return [];
}

// Incomplete
function twoNumberSum(array, targetSum) {
	for (let i = 0; i <= array.length; i += 1) {
		let firstNum = array[i];
		let secondNum = targetSum - firstNum;
		if (firstNum !== secondNum && array.includes(secondNum)) {
			return [firstNum, secondNum];
		}
	}
	
	return [];
}

// Worst - O(n^2) time | O(1) space
function twoNumberSum(array, targetSum) {
  for (let i = 0; i <= array.length; i += 1) {
    let compare = array[i];

    for (let k = 0; k <= array.length; k += 1) {
      let next = array[k+1];
      if (compare === next) {
        next = array[k+2];
      }

      let sum = compare + next;
      if (sum === targetSum) {
        return [compare, next];
      }
    }
  }

  return [];
}
