// Quick Sort - CTCI
var CTCI = {
  quickSort: function (arr, left, right) {
    let index = this.partition(arr, left, right);
    if (left < index - 1) { // Sort the left half
      this.quickSort(arr, index, right);
    }
    if (index < right) { // sort the right half
      this.quickSort(arr, index, right);
    }
  },
  partition: function (arr, left, right) {
    let pivot = arr[(left + right) / 2]; // Pick pivot point
    while (left <= right) {
      // Find element on left that should be on right
      while (arr[left] < pivot) left++;

      // Find element on right that should be on left
      while (arr[right] > pivot) right--;

      // Swap elements, and move left and right indices
      if (left <= right) {
        swap(arr, left, right); // swaps elements
        left++;
        right--;
      }
    }
    return left;
  }
};


