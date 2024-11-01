"use strict";

/*

Problem
- reading until a choice (iterating)
- choice contains a page to check the choice
- the next two elements are what page to go to
- ending pages in a sorted list 
- you are always given the option that you will take
- return the ending or -1 if stuck in a loop

Solution
- iterate starting from page 1
- keep track of choice with j = 0
- when a page === choices[j][1] then pick option
- set page to selection, increment j and continue looping
- while iterating check if page number equals ending and return number
- keep a hash of previously seen choices if we've seen the choice before then return -1

 */

const chooseAdventure = function(endings, choices, option) {

  const seen = {};
  let j = 0;
  for (let i = 1; i < 100; i++) {
    console.log("i is", i);
    console.log("j is", j);

    let choice = choices[j][option];
    console.log("choice is ", choice);

    if (seen[i] === 1) {
      return -1;
    }

    if (endings.includes(i)) {
      return i;
    }

    if (choices[j][0] === i) {
      i = choices[j][option];
      console.log("i is now", i);

      if (j < choices.length - 1) {
        j++
      }

      seen[i] = 1;
    }
  }

  return -1;
};

module.exports = chooseAdventure;
