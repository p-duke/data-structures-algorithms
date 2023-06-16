// Solution 1
function minimumWaitingTime(queries) {
	let finalWait = 0;
	let runningTotal = 0;
	const sorted = queries.sort((a,b) => a - b);
	for (let i = 0; i < sorted.length - 1; i++) {
		let num = sorted[i];
		runningTotal = runningTotal + num;
		finalWait = finalWait + runningTotal;
	}
	
  return finalWait;
}

// Do not edit the line below.
exports.minimumWaitingTime = minimumWaitingTime;

