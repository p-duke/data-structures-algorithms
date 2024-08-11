import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
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
