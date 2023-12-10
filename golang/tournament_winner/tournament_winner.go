package main

import (
	"fmt"
	"reflect"
)

/*
Title: Tournament Winner
Difficulty: Easy
Technique:

---

Question: There's an algorithms tournament taking place in which teams of programmers compete against each other to solve algorithmic problems as fast as possible. Teams compete in a round robin, where each team faces off against all other teams. Only two teams compete against each other at a time, and for each competition, one team is designated the home team, while the other team is the away team. In each competition there's always one winner and one loser; there are no ties. A team receives 3 points if it wins and 0 points if it loses. The winner of the tournament is the team that receives the most amount of points.

Given an array of pairs representing the teams that have competed against each other and an array containing the results of each competition, write a function that returns the winner of the tournament. The input arrays are named competitions and results, respectively. The competitions array has elements in the form of [homeTeam, awayTeam], where each team is a string of at most 30 characters representing the name of the team. The results array contains information about the winner of each corresponding competition in the competitions array. Specifically, results[i] denotes the winner of competitions[i], where a 1 in the results array means that the home team in the corresponding competition won and a 0 means that the away team won.

It's guaranteed that exactly one team will win the tournament and that each team will compete against all other teams exactly once. It's also guaranteed that the tournament will always have at least two teams.

Sample Input
competitions = [

	["HTML", "C#"],
	["C#", "Python"],
	["Python", "HTML"],

]
results = [0, 0, 1]

Sample Output
"Python"
// C# beats HTML, Python Beats C#, and Python Beats HTML.
// HTML - 0 points
// C# -  3 points
// Python -  6 points

---
*/
func tournamentWinner(competitions [][]string, results []int) string {
	// Problem
	// Team receives 3 points if win / 0 points if lose
	// Input - array of pairs (homeTeam, awayTeam)
	// Input - In results 1 = home team win / 0 = away team win
	// Return - Name of winner

	// Solution
	// Loop over competitions
	// Add teams to a map where key is name and value is score
	// If team is not in map then add with beginning score
	// Else increment the teams score (+3)
	// After competitions loop then loop over score map and return winner
	// Time complexity: O(n + m) | Space complexity: O(n)
	scores := make(map[string]int)
	var winner string
	var highest string
	for i, v := range competitions {
		home := v[0]
		away := v[1]
		result := results[i]

		if result == 1 {
			winner = home
		} else {
			winner = away
		}

		if _, ok := scores[winner]; !ok {
			scores[winner] = 3
		} else {
			scores[winner] += 3
		}

		if len(highest) == 0 {
			highest = winner
		} else if scores[home] > scores[highest] {
			highest = home
		} else if scores[away] > scores[highest] {
			highest = away
		}
	}

	return highest
}

func main() {
	tests := []struct {
		competitions [][]string
		results      []int
		want         string
	}{
		{competitions: [][]string{[]string{"HTML", "C#"}, []string{"C#", "Python"}, []string{"Python", "HTML"}}, results: []int{0, 0, 1}, want: "Python"},
		{competitions: [][]string{[]string{"HTML", "Java"}, []string{"Java", "Python"}, []string{"Python", "HTML"}}, results: []int{0, 1, 1}, want: "Java"},
		{competitions: [][]string{
			[]string{"HTML", "Java"},
			[]string{"Java", "Python"},
			[]string{"Python", "HTML"},
			[]string{"C#", "Python"},
			[]string{"Java", "C#"},
			[]string{"C#", "HTML"},
		}, results: []int{0, 1, 1, 1, 0, 1}, want: "C#"},
		{competitions: [][]string{
			[]string{"HTML", "Java"},
			[]string{"Java", "Python"},
			[]string{"Python", "HTML"},
			[]string{"C#", "Python"},
			[]string{"Java", "C#"},
			[]string{"C#", "HTML"},
			[]string{"SQL", "C#"},
			[]string{"HTML", "SQL"},
			[]string{"SQL", "Python"},
			[]string{"SQL", "Java"},
		}, results: []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0}, want: "C#"},
		{competitions: [][]string{
			[]string{"Bulls", "Eagles"},
		}, results: []int{1}, want: "Bulls"},
		{competitions: [][]string{
			[]string{"Bulls", "Eagles"},
			[]string{"Bulls", "Bears"},
			[]string{"Bears", "Eagles"},
		}, results: []int{0, 0, 0}, want: "Eagles"},
		{competitions: [][]string{
			[]string{"Bulls", "Eagles"},
			[]string{"Bulls", "Bears"},
			[]string{"Bulls", "Monkeys"},
			[]string{"Eagles", "Bears"},
			[]string{"Eagles", "Monkeys"},
			[]string{"Bears", "Monkeys"},
		}, results: []int{1, 1, 1, 1, 1, 1}, want: "Bulls"},
		{competitions: [][]string{
			[]string{"AlgoMasters", "FrontPage Freebirds"},
			[]string{"Runtime Terror", "Static Startup"},
			[]string{"WeC#", "Hypertext Assassins"},
			[]string{"AlgoMasters", "WeC#"},
			[]string{"Static Startup", "Hypertext Assassins"},
			[]string{"Runtime Terror", "FrontPage Freebirds"},
			[]string{"AlgoMasters", "Runtime Terror"},
			[]string{"Hypertext Assassins", "FrontPage Freebirds"},
			[]string{"Static Startup", "WeC#"},
			[]string{"AlgoMasters", "Static Startup"},
			[]string{"FrontPage Freebirds", "WeC#"},
			[]string{"Hypertext Assassins", "Runtime Terror"},
			[]string{"AlgoMasters", "Hypertext Assassins"},
			[]string{"WeC#", "Runtime Terror"},
			[]string{"FrontPage Freebirds", "Static Startup"},
		}, results: []int{1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0}, want: "AlgoMasters"},
		{competitions: [][]string{
			[]string{"HTML", "Java"},
			[]string{"Java", "Python"},
			[]string{"Python", "HTML"},
			[]string{"C#", "Python"},
			[]string{"Java", "C#"},
			[]string{"C#", "HTML"},
			[]string{"SQL", "C#"},
			[]string{"HTML", "SQL"},
			[]string{"SQL", "Python"},
			[]string{"SQL", "Java"},
		}, results: []int{0, 0, 0, 0, 0, 0, 1, 0, 1, 1}, want: "SQL"},
		{competitions: [][]string{[]string{"A", "B"}}, results: []int{0}, want: "B"},
	}

	for _, tc := range tests {
		got := tournamentWinner(tc.competitions, tc.results)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAILED! Got: %v Want: %v\n", got, tc.want)
		} else {
			fmt.Printf("PASSED! Got: %v Want: %v\n", got, tc.want)
		}
	}
}
