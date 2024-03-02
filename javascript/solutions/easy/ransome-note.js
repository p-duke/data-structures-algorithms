/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
var canConstruct = function(ransomNote, magazine) {
   // Problem
   // Return true if the ransom note can be made from letters in magazine
   // Each letter can only be used once

   // Solution
   // Can loop through the ransomNote and count letters
   const count = {};
   for (let i = 0; i < magazine.length; i++) {
       const letter = magazine[i];
       if (count[letter] !== undefined) {
           count[letter] += 1;
       } else {
           count[letter] = 1;
       }
   }

   for (let i = 0; i < ransomNote.length; i++) {
       const letter = ransomNote[i];
       if (count[letter] !== undefined) {
           count[letter] -= 1;
       } else {
           count[letter] = -1;
       }
   }

    for (let key in count) {
        if (count[key] < 0) {
            return false;
        }
    }

    return true;
};
