// Solution 3
function tandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest) {
  
}



// Solution 1
function tandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest) {
	const red = redShirtSpeeds.sort((a, b) => {
		if (fastest) return b - a;
		return a - b;
	});
	const blue = blueShirtSpeeds.sort((a, b) => a - b);
	let result = 0;
	for (let i = 0; i < red.length; i++) {
		const redRider = red[i];
		const blueRider = blue[i];
		result += Math.max(redRider, blueRider);
	}
	
  return result;
}

// Do not edit the line below.
exports.tandemBicycle = tandemBicycle;

