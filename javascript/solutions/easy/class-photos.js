// Solution 3
function classPhotos(redShirtHeights, blueShirtHeights) {
  redShirtHeights.sort((a,b) => b - a);
  blueShirtHeights.sort((a,b) => b - a);
	const shirtColorInFirstRow = redShirtHeights[0] < blueShirtHeights[0] ? 'red' : 'blue';
	for (let i = 0; i < redShirtHeights.length; i++) {
		const redShirtHeight = redShirtHeights[i];
		const blueShirtHeight = blueShirtHeights[i];
		
		if (shirtColorInFirstRow  === 'red') {
			if (redShirtHeight >= blueShirtHeight) {
				return false;
			}
		} else {
			if (blueShirtHeight >= redShirtHeight) {
				return false;
			}
		}
	}
	
	return true;
}

// Do not edit the line below.
exports.classPhotos = classPhotos;


// Solution 2
function classPhotos(redShirtHeights, blueShirtHeights) {
	redShirtHeights.sort((a,b) => b - a);
	blueShirtHeights.sort((a,b) => b - a);

  let valid = false;	
	let backRow = redShirtHeights[0] > blueShirtHeights[0] ? 'red' : 'blue';
	for (let i = 0; i < redShirtHeights.length; i++) {
		let backRowRed = backRow === 'red' && redShirtHeights[i] > blueShirtHeights[i];
		let backRowBlue = backRow === 'blue' && blueShirtHeights[i] > redShirtHeights[i];
		if (backRowRed || backRowBlue) {
			valid = true;
		} else {
			return false;
		}
	}
	
  return valid;
}

// Do not edit the line below.
exports.classPhotos = classPhotos;
