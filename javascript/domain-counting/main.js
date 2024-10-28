/*
 * Problem
 * csv input of clicks per domain
 * return number of clicks that were recorded on each domain and subdomain
 *
 * 
 * Example
 * mail.yahoo.com = mail.yahoo.com, yahoo.com and .com
 *
 * Solution
 * we need to sum up start from each domain on the left
 * declare a hash to store entries
 * loop over the list and split each domain
 * Then loop over each subdomain and add to the hashmap with count
 * declare a latest that concatenates the last domain
 * return the hashmap at the end
 *
 */

const calculateClicksByDomain = (domains) => {
    const totalClicks = {};

    for (let i = 0; i < domains.length; i++) {
        let curr = domains[i];
        let [count, domain] = curr.split(",");
        let splitD = domain.split(".");
        let latest = "";

        for (let j = splitD.length - 1; j >= 0; j--) {
            let sub = splitD[j];
            if (latest === "") {
                latest += sub
            } else {
                latest = sub + "." + latest
            }

            if (totalClicks[latest] !== undefined) {
                totalClicks[latest] += parseInt(count);
            } else {
                totalClicks[latest] = parseInt(count);
            }
        }
    };

    return totalClicks;
}

module.exports = calculateClicksByDomain;
