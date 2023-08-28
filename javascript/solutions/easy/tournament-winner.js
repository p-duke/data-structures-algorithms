// Solution 2
function tournamentWinner(competitions, results) {
  let currentWinner;
	let bestWinner;
	let scores = {};
	competitions.forEach((comp, idx) => {
	  let result = results[idx];
		currentWinner = result === 1 ? comp[0] : comp[1];
		if (scores[currentWinner] === undefined) {
		  scores[currentWinner] = result
			bestWinner = currentWinner;
		}
		
		if (scores[currentWinner] > scores[bestWinner]) {
		  bestWinner = currentWinner;
			scores[currentWinner] += 1;
		} else {
		  scores[currentWinner] += 1;
		}
	});
  return bestWinner;
}

// Do not edit the line below.
exports.tournamentWinner = tournamentWinner;

// Solution 1
function tournamentWinner(competitions, results) {
	let score = {};
	competitions.forEach((comp, idx) => {
	  if (results[idx] === 1) {
		  const winner = comp[0];
			if (score[winner] === undefined) {
				score[winner] = 1;
			} else {
			  score[winner] += 1;
			}
		} else if (results[idx] === 0) {
		  const winner = comp[1];
			if (score[winner] === undefined) {
				score[winner] = 1;
			} else {
			  score[winner] += 1;
			}
		}
	});
	
	let highest = 0;
	let winningTeam = '';
	Object.keys(score).forEach((team) => {
		if (score[team] > highest) {
		  winningTeam = team;
			highest = score[team];
		}
	});
	
  return winningTeam;
}

// Do not edit the line below.
exports.tournamentWinner = tournamentWinner;

