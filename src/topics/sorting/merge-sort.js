// Merge Sort - CTCI
function main(array) {
  let helper = array.length;
  mergeSort(array, helper, 0, array.length - 1);
}

function mergeSort(array, helper, low, high) {
  if (low < high) {
    let middle = (low + high) / 2;
    mergeSort(array, helper, low, middle); // Sort left half
    mergeSort(array, helper, middle+1, high); // Sort right half
    merge(array, helper, low, middle, high); // Sort right half
  }
}

function merge(array, helper, low, middle, high) {
  // Copy both halves inter helper array
  for (let i = low; i <= high; i++) {
    helper[i] = array[i];
  }

  let helperLeft = low;
  let helperRight = middle + 1;
  let current = low;

  /* 
   * Iterate though helper array. Compare the left and right half, copying back
   * the smaller element from the two halves into the original array
   */
  while(helperLeft <= middle && helperRight <= high) {
    if (helper[helperLeft] <= helper[helperRight]) {
      array[current] = helper[helperLeft];
      helperRight++;
    } else { // If right element is smaller than left element
      array[current] = helper[helperRight];
      helperRight++;
    }

    current++;
  }

  /* Copy the rest of the left side of the array into the target array */
  let remaining = middle - helperLeft;
  for (let i = 0; i <= remaining; i++) {
    array[current + i] = helper[helperLeft + i];
  }
}

// InterviewCake
var interviewCake = {
  combineSortedArrays: function (arrayOne, arrayTwo) {
    let arrayOneIndex = 0;
    let arrayTwoIndex = 0;
    let mergedArrayIndex = 0;
    const mergedArray = [];

    // Both arrays have some items left in them
    while (arrayOneIndex < arrayOne.length && arrayTwoIndex < arrayTwo.length) {
      // Choose the smaller of the two items and add it to the merged array
      if (arrayOne[arrayOneIndex] <= arrayTwo[arrayTwoIndex]) {
        mergedArray[mergedArrayIndex] = arrayOne[arrayOneIndex];
        arrayOneIndex += 1;
      } else {
        mergedArray[mergedArrayIndex] = arrayTwo[arrayTwoIndex];
        arrayTwoIndex += 1;
      }
      mergedArrayIndex += 1;
    }

    // Grab any lingering items in the first array after we've exhausted the second array
    while (arrayOneIndex < arrayOne.length) {
      mergedArray[mergedArrayIndex] = arrayOne[arrayOneIndex];
      mergedArrayIndex += 1;
      arrayOneIndex += 1;
    }

    // Grab any lingering items in the second array after we've exhausted the first array
    while (arrayTwoIndex < arrayTwo.length) {
      mergedArray[mergedArrayIndex] = arrayTwo[arrayTwoIndex];
      mergedArrayIndex += 1;
      arrayTwoIndex += 1;
    }

    return mergedArray;
  },
  mergeSort: function (theArray) {
    // Base case - single element array
    if (theArray.length <= 1) {
      return theArray;
    }

    // Split the input in half
    const middleIndex = theArray.length / 2;
    const left = theArray.slice(0, middleIndex);
    const right = theArray.slice(middleIndex, theArray.length);

    // Sort each half
    const leftSorted = this.mergeSort(left);
    const rightSorted = this.mergeSort(right);

    // Merge the sorted halves
    return this.combineSortedArrays(leftSorted, rightSorted);
  }
};

console.log(interviewCake.mergeSort([30, 50, 70, 3, 13, 2, 100]));
