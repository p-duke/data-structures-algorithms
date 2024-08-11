func groupAnagrams(strs []string) [][]string {
	// Problem
	// input is slice of anagrams
	// order the anagrams into groups

	// Solution
	// create a map to hold anagrams (slice)
	// loop over strings
	// split, sort, and join the string
	// check if in map
	// if in map then add to array
	// else add to map
	// loop over map into slice of slices
	result := make([][]string, 0)
	count := make(map[string][]string)
	for _, word := range strs {
		splitWord := strings.Split(word, "")
		sort.Strings(splitWord)
		joined := strings.Join(splitWord, "")
		if vals, ok := count[joined]; !ok {
			count[joined] = []string{word}
		} else {
			count[joined] = append(vals, word)
		}
	}

	for _, v := range count {
		result = append(result, v)
	}

	return result
}
