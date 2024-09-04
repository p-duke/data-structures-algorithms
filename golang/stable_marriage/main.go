package main

import "fmt"

/*

Problem Description:
- The Stable Marriage Problem involves finding a stable matching between two sets of elements.
- The problem is typically described in terms of two sets of people: one set of "men" and one set of "women."
- Each person in both sets ranks all members of the opposite set based on their preference for marriage.

The goal is to find a "stable matching," where:

1. Every man is paired with exactly one woman, and every woman is paired with exactly one man.
2. No pair of people (man and woman) would both prefer each other over their current partners. If such a pair exists, the matching is not stable, because those two people would prefer to leave their current partners to be with each other.

Example Scenario:
- Imagine there are three men (A, B, C) and three women (X, Y, Z).
- Each man ranks the women in order of preference, and each woman ranks the men similarly.
- The problem is to pair each man with a woman in such a way that there is no instability (i.e., no two people would prefer each other over their current partners).

Preferred Technique:
- The preferred technique for solving the Stable Marriage Problem is Gale-Shapley Algorithm

General Approach:
- Proposals and Rejections: The algorithm iteratively has one set of individuals (e.g., men) propose to their top choice among the available members of the opposite set (e.g., women).
- If a person from the opposite set is free, they accept the proposal.
- If they are already engaged, they compare the new proposal with their current engagement and may decide to reject their current partner in favor of the new proposal if the new proposer is preferred.

*/

func wPrefersM1OverM(prefer [][]int, w, m, m1 int) bool {
	return false
}

func stableMarriage(prefer [][]int) {
	// Solution
	// Divide into men and woman
	// Loop over men and try to give them their first choice
	// For each man loop through choices
	// 
	for _, v := range prefer {
		fmt.Println("val", v)
	}
}

func main() {
	prefer := [][]int{ 
		{7, 5, 6, 4}, // m
        {5, 4, 6, 7}, // m
        {4, 5, 6, 7}, // m
        {4, 5, 6, 7}, // m
        {0, 1, 2, 3}, // f
        {0, 1, 2, 3}, // f
        {0, 1, 2, 3}, // f
        {0, 1, 2, 3}, // f
    }

    stableMarriage(prefer)
}
