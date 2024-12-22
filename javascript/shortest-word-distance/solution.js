/**
 * @param {string[]} wordsDict
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
var shortestDistance = function(wordsDict, word1, word2) {
    let minDist = wordsDict.length;
    let widx1;
    let widx2;
    for (let i = 0; i < wordsDict.length; i++) {
        let word = wordsDict[i];

        if (word === word1) {
            widx1 = i;
        } else if (word === word2) {
            widx2 = i;
        }

        let diff = Math.abs(widx1 - widx2);
        if (diff < minDist) {
            minDist = diff;
        }
    }

    return minDist;
};
