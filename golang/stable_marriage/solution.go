package main

import "fmt"

// prefersNewMan checks if a woman prefers a new man over her current partner.
func prefersNewMan(preferences [][]int, woman int, newMan int, currentMan int, n int) bool {
	// Iterate through the woman's preference list.
	for _, man := range preferences[n+woman] {
		// If she prefers the new man, return true.
		if man == newMan {
			return true
		}
		// If she reaches the current man first, she prefers the current man.
		if man == currentMan {
			return false
		}
	}
	return false
}

// stableMarriage finds the stable matching using the Gale-Shapley algorithm.
func stableMarriage(preferences [][]int) []int {
	n := len(preferences) / 2 // Number of men and women.

	// womenPartners stores the current partner for each woman. -1 means no partner yet.
	womenPartners := make([]int, n)
	
	// menFree keeps track of whether each man is free or not.
	menFree := make([]bool, n)
	
	// nextProposals tracks the index of the next woman each man will propose to.
	nextProposals := make([]int, n)

	// Initialize all women as not engaged (-1 means no partner).
	for i := 0; i < n; i++ {
		womenPartners[i] = -1
	}

	freeCount := n // Number of free men (initially all men are free).

	// Loop until there are no free men left.
	for freeCount > 0 {
		var man int
		
		// Find the first free man.
		for man = 0; man < n; man++ {
			if !menFree[man] {
				break
			}
		}

		// The man proposes to the next woman on his preference list.
		woman := preferences[man][nextProposals[man]]
		nextProposals[man]++ // Move to the next woman for the next proposal.

		// If the woman is free, they get engaged.
		if womenPartners[woman] == -1 {
			womenPartners[woman] = man // Woman is engaged to this man.
			menFree[man] = true        // The man is no longer free.
			freeCount--                // One less free man.
		} else {
			// The woman is already engaged, so check if she prefers the new man.
			currentPartner := womenPartners[woman]

			// If the woman prefers the new man over her current partner:
			if prefersNewMan(preferences, woman, man, currentPartner, n) {
				womenPartners[woman] = man // She switches to the new man.
				menFree[man] = true        // The new man is now engaged.
				menFree[currentPartner] = false // The previous partner is now free.
			}
		}
	}

	// Prepare the final result array where each index corresponds to a man,
	// and the value at that index is the woman they are matched with.
	result := make([]int, n)
	for woman, man := range womenPartners {
		result[man] = woman
	}
	return result
}

func main() {
	// Example combined preference list.
	preferences := [][]int{
		{7, 5, 6, 4}, // Preferences for Man 0
		{5, 4, 6, 7}, // Preferences for Man 1
		{4, 5, 6, 7}, // Preferences for Man 2
		{4, 5, 6, 7}, // Preferences for Man 3
		{0, 1, 2, 3}, // Preferences for Woman 0
		{0, 1, 2, 3}, // Preferences for Woman 1
		{0, 1, 2, 3}, // Preferences for Woman 2
		{0, 1, 2, 3}, // Preferences for Woman 3
	}

	// Get the stable matching results
	result := stableMarriage(preferences)

	// Print out the stable matchings.
	fmt.Println("The stable matchings are:")
	for man, woman := range result {
		fmt.Printf("Man %d is matched with Woman %d\n", man, woman)
	}
}

