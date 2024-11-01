const chooseAdventure = function(endings1, choices1_1, option) {
    let choice = 0;
    let pagesSeen = {};
    let endingPage;
    // 1 -> 2 -> 3 -> 4 ->
    for (let page = 1; page <= endings1[endings1.length - 1]; ) {
        console.log("current page", page);
        if (pagesSeen[page] === 1) {
            console.log("saw page",page);
            return -1;
        }
        pagesSeen[page] = 1
        // Possible error
        if (choices1_1[choice]) {
            if (page === choices1_1[choice][0]) {
                console.log("getting here");
                page = choices1_1[choice][option];
                choice += 1;
            } else {
                page++;
            }
        } else {
            page++;

        }
        if (endings1.indexOf(page) > -1) {
            endingPage = page;
            break;
        }
    }
    return endingPage;
};
