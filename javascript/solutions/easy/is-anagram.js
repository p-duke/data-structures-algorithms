/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
    // Solution
    // if length of s is not equal to t then return false
    if (s.length !== t.length) {
        return false;
    }

    // loop through once and increment s and decrement t
    // loop through again and check if all are zero
    // if not then return false else return true
    const record = {};
    for (let i = 0; i < s.length; i++) {
        const sLet = s[i];
        const tLet = t[i];

        if (!record[sLet]) {
            record[sLet] = 1;
        } else {
            record[sLet] += 1;
        }

        if (!record[tLet]) {
            record[tLet] = -1;
        } else {
            record[tLet] -= 1;
        }
    }
    
    return Object.values(record).every((rec) => rec == 0);
};
