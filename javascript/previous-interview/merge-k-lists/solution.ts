/*
Create a function that accepts any number of sorted integer array and returns a single sorted integer array containing all of the elements of the input array.

Example: Three input slices are given,
{3, 9},
{0, 5, 8}, and
{2, 6, 7, 8}.

The output of the function that provides a valid solution should be
{0, 2, 3, 5, 6, 7, 8, 8, 9}.

 */

// Solution
// - option 1 - loop over and add each array together into one
//    - Then run a merge sort
// - option 2 - create a k-way merge function to merge two together
//    - Push the first array into the "result"
//    - Loop over the arrays and merge together two at a time
const mergeKArrays = (arrays: number[][]): number[] => {
  if (arrays.length === 0) {
    return [];
  }

  // O(n) time | O(n) space
  const unsortedMerged: number[] = [];
  for (let [_, v] of arrays.entries()) {
    unsortedMerged.push(...v);
  }

  // Implement merge sort here
  // - separate function to kick off mergeSort
  // - recursively break down the array into single items
  // - pass those into the sort (k-way merge) function
  return mergeSort(unsortedMerged);
}

// first have base case that returns element when length of 1
// find the mid, left, and right
// call mergeSort on left
// call mergeSort on right
// pass those into sort
const mergeSort = (arr: number[]): number[] => {
  if (arr.length === 1) {
    return arr;
  }

  let result: number[] = [];
  const mid = arr.length / 2;
  const leftArr = mergeSort(arr.slice(0, mid));
  const rightArr = mergeSort(arr.slice(mid, arr.length));

  result = sort(leftArr, rightArr, result);

  return result;
};

// Sort takes left and right arr and a result
const sort = (left: number[], right: number[], result: number[]): number[] => {
  let i = 0, j = 0;
  while (i < left.length && j < right.length) {
    if (left[i] < right[j]) {
      result.push(left[i]);
      i++;
    } else {
      result.push(right[j]);
      j++;
    }
  }

  while (i < left.length) {
    result.push(left[i]);
    i++;
  }

  while (j < right.length) {
    result.push(right[j]);
    j++;
  }

  return result;
};

const main = () => {
  const testCases: { input: number[][]; expected: number[] }[] = [
    {
      input: [
        [3, 9],
        [0, 5, 8],
        [2, 6, 7, 8],
      ],
      expected: [0, 2, 3, 5, 6, 7, 8, 8, 9],
    },
    {
      input: [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
      ],
      expected: [1, 2, 3, 4, 5, 6, 7, 8, 9],
    },
    {
      input: [
        [],
        [0],
        [],
      ],
      expected: [0],
    },
    {
      input: [
        [5, 10],
        [1, 1, 2],
        [3, 4, 5],
      ],
      expected: [1, 1, 2, 3, 4, 5, 5, 10],
    },
    {
      input: [],
      expected: [],
    },
    {
      input: [
        [100],
        [50, 60, 70],
        [5, 6, 7],
      ],
      expected: [5, 6, 7, 50, 60, 70, 100],
    },
  ];

  for (const [i, testCase] of testCases.entries()) {
    const result = mergeKArrays(testCase.input);
    const passed = JSON.stringify(result) === JSON.stringify(testCase.expected);
    console.log(`Test case #${i + 1}: ${passed ? '✅ Passed' : '❌ Failed'}`);
    if (!passed) {
      console.log(`  Input: ${JSON.stringify(testCase.input)}`);
      console.log(`  Expected: ${JSON.stringify(testCase.expected)}`);
      console.log(`  Got: ${JSON.stringify(result)}`);
    }
  }
}

// Run the tests
main();
