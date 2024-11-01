var chooseAdventure = require('./index.js');

describe("it should pass all tests", () => {

    // Endings --------------------------------
    const endings1 = [6, 15, 21, 30];
    const endings2 = [11];
    const endings3 = [4, 11];
    const endings4 = [20];

    // Choices --------------------------------
    const choices1_1 = [ [3, 7, 8], [9, 4, 2] ];
    const choices1_2 = [ [3, 14, 2] ];
    const choices1_3 = [
        [5, 11, 28],
        [9, 19, 29],
        [14, 16, 20],
        [18, 7, 22],
        [25, 6, 30]
    ];
    const choices1_4 = [
        [2, 10, 15],
        [3, 4, 10],
        [4, 3, 15],
        [10, 3, 15]
    ];

    const choices2_1 = [
        [2, 3, 4],
        [5, 10, 2]
    ];
    const choices2_2 = [];
    const choices3_1 = [ [10, 6, 8] ];
    const choices4_1 = [
        [2, 6, 3],
        [3, 1, 4],
        [4, 10, 5],
        [6, 3, 7]
    ];

    tests = [
        { endings: endings1, choices: choices1_1, option: 1, expected: 6 },
        { endings: endings1, choices: choices1_1, option: 2, expected: -1 },
        // { endings: endings1, choices: choices1_2, option: 1, expected: 15 },
        // { endings: endings1, choices: choices1_2, option: 2, expected: -1 },
        // { endings: endings1, choices: choices1_3, option: 2, expected: 30 },
        // { endings: endings1, choices: choices1_4, option: 2, expected: 15 },
        // { endings: endings2, choices: choices2_1, option: 1, expected: 11 },
        // { endings: endings2, choices: choices2_1, option: 2, expected: -1 },
        // { endings: endings2, choices: choices2_2, option: 1, expected: 11 },
        // { endings: endings2, choices: choices2_2, option: 2, expected: 11 },
        // { endings: endings3, choices: choices3_1, option: 1, expected: 4 },
        // { endings: endings3, choices: choices3_1, option: 2, expected: 4 },
        // { endings: endings1, choices: choices1_4, option: 1, expected: -1 },
        // { endings: endings4, choices: choices4_1, option: 1, expected: -1 },
    ];

    for (let idx in tests) {
        let tc = tests[idx];
        let result = chooseAdventure(tc.endings, tc.choices, tc.option);
        test(`ending: ${tc.endings}; choice: ${tc.choices} should equal ${tc.expected}`, () => {
            expect(result).toEqual(tc.expected);
        });
    }
});
