function nonConstructibleChange(coins) {
	const sortedCoins = coins.sort((a,b) => a - b);
	let currentChange = 0;

	for (let i = 0; i < sortedCoins.length; i++) {
		let coin = sortedCoins[i];
		if (coin > (currentChange + 1)) {
			return currentChange + 1;
		}
		currentChange += coin;
	}
	
	return currentChange + 1;
}

// Do not edit the line below.
exports.nonConstructibleChange = nonConstructibleChange;

