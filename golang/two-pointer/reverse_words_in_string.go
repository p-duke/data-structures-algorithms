package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// EDUCATIVE SOLUTION ---------------------------------------------------------------------------------------
func reverseWordsEducative(sentence string) string {
	// remove leading, trailing and multiple spaces
	trimedSentence := trimString(sentence)

	// We need to convert the input strings to lists of characters as strings are immutable in Go
	sentenceBytes := []byte(trimedSentence)

	strLen := len(sentenceBytes)

	// We will first reverse the entire string.
	sentenceBytes = strRev(sentenceBytes, 0, strLen-1)

	// Now all the words are in the desired location, but
	// in reverse order: "Hello World" -> "dlroW olleH".

	start, end := 0, 0

	for start < strLen {

		for end < strLen && sentenceBytes[end] != ' ' {
			end += 1
		}

		// Let's call our helper function to reverse the word in-place.
		strRev(sentenceBytes, start, end-1)
		start = end + 1
		end += 1
	}
	return string(sentenceBytes)
}

func trimString(str string) string {
	// Trim extra spaces at the beginning and end of the string
	str = strings.TrimSpace(str)

	// Replace multiple spaces with a single space
	regex := regexp.MustCompile("\\s+")
	str = regex.ReplaceAllString(str, " ")

	return str
}

// strRev function that reverses a whole sentence character by character
func strRev(str []byte, startRev int, endRev int) []byte {
	// Starting from the two ends of the list, and moving
	// in towards the centre of the string, swap the characters
	for startRev < endRev {
		temp := str[startRev]       // temp store for swapping
		str[startRev] = str[endRev] // swap step 1
		str[endRev] = temp          // swap step 2

		startRev += 1 // move forwards towards the middle
		endRev -= 1   // move backwards towards the middle
	}
	return str
}

// MY SOLUTION ---------------------------------------------------------------------------------------
// The time complexity of the function “reverseWords” is O(n), where n is the length of the input string
// The time complexity of the built-in Go functions invoked by the code are as follows:
// strings.Split: O(n), where n is the length of the input string
// regexp.MatchString: O(n), where n is the length of the input string
// strings.Join: O(n), where n is the length of the input string

// The space complexity of the “reverseWords” function is O(n), where n is the length of the input sentence
// This is because the function creates a new array of strings to store the words in the reversed order,
// which can be up to the same length as the input sentence
func reverseWords(sentence string) string {
	splitString := strings.Split(sentence, " ")

	left := 0
	right := len(splitString) - 1

	for left < right {
		var rightVal string = splitString[right]
		var leftVal string = splitString[left]

		if leftVal == " " {
			left++
		} else if rightVal == " " {
			right--
		} else {
			splitString[left] = rightVal
			splitString[right] = leftVal
			left++
			right--
		}
	}

	var newSentence []string
	for _, v := range splitString {
		if match, _ := regexp.MatchString(`\w`, v); match {
			newSentence = append(newSentence, v)
		}
	}

	return strings.Join(newSentence, " ")
}

func main() {
	tests := []struct {
		sentence string
		want     string
	}{
		{sentence: "We love GoLang", want: "GoLang love We"},
		{sentence: "To be or not to be", want: "be to not or be To"},
		{sentence: "You are amazing", want: "amazing are You"},
		{sentence: "Hello     World", want: "World Hello"},
		{sentence: "Hey", want: "Hey"},
	}

	for _, tc := range tests {
		got := reverseWords(tc.sentence)
		if !reflect.DeepEqual(tc.want, got) {
			fmt.Printf("FAIL! Expected: %s --- Got: %s\n", tc.want, got)
		} else {
			fmt.Printf("PASS Expected: %s --- Got: %s\n", tc.want, got)
		}
	}
}
