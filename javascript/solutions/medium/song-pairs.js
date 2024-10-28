'use strict';

/*
 * Complete the 'playlist' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY songs as parameter.
 * 
 * Notes
 * - input is an array integers (seconds)
 * - We're looking for a pair (two el) that equal 60
 * - Want to count all pairs that equal 60
 * - Array is not sorted
 * 
 * Solution
 * - declare a count
 * - decalre a hashmap seen
 * - loop through songs
 * - take a diff against the target which is 60
 * - if not in hashmap then add
 * - if in hashmap then increment count
 * - return count
 * 
 * Trade offs 
 * - Two pointer - But we need to sort - O(n log(n))
 * - hashmap = space complexity would be affected
 * 
 * Edge Case
 * - Could be duplicate numbers (e.g. two 10s)
 * 
 * [30a, 30b, 30c]
 * 
 * [30a, 30b, 30a, 30c]
 * 
 * [30b, 30c]
 * 
 * Sample case 1
 * [30, 20, 30, 40, 40, 40]
 * [30, 30, 30]
 * 
 *  count - 2
 * {
 *  30: 1,
 * }
 */

// TODO: INCOMPLETE
function playlist(songs) {
   let count = 0;
   let seen = {};
   const target = 60;
   
   for (let i = 0; i < songs.length; i++) {
       let song = songs[i];
       if (seen[song] !== undefined) {
           seen[song] += 1;
       } else {
           seen[song] = 1;
       }
   }
   // count - 2
   console.log("seen", seen);
   for (let i = 0; i < songs.length; i++) {
       let curr = songs[i];
       let diff = target - curr;
       if (seen[diff] !== undefined && seen[diff] > 0) {
           count += 1;
           seen[curr] -= 1;
       }
   }
   
   return count;
}


const songs = [30, 30,30]; // 3
const songsTwo = [10, 20, 40, 50]; // 2
const songsFour = [10,20,40,40]; // 2
const songsThree = [20,20,20,20,40,40,40,40]; // 16

console.log("output:", playlist(songsThree));
