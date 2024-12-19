package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*

Title: Caesar Cipher Encryptor

---

Question:

  Given a non-empty string of lowercase letters and a non-negative integer
  representing a key, write a function that returns a new string obtained by
  shifting every letter in the input string by k positions in the alphabet,
  where k is the key.


  Note that letters should "wrap" around the alphabet; in other words, the
  letter z shifted by one returns the letter a.

Sample Input
string = "xyz"
key = 2

Sample Output
"zab"

---

Hint: HintsHint 1
Most languages have built-in functions that give you the Unicode value of a character as well as the character corresponding to a Unicode value. Consider using such functions to determine which letters the input string's letters should be mapped to.

Hint 2

Try creating your own mapping of letters to codes. In other words, try associating each letter in the alphabet with a specific number - its position in the alphabet, for instance - and using that to determine which letters the input string's letters should be mapped to.

Hint 3

How do you handle cases where a letter gets shifted to a position that requires wrapping around the alphabet? What about cases where the key is very large and causes multiple wrappings around the alphabet? The modulo operator should be your friend here.
Optimal Space & Time ComplexityO(n) time | O(n) space - where n is the length of the input string

*/

func caesarCipherEncryptor(str string, key int) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	var result string
	for _, v := range str {
		currPos := strings.Index(alpha, string(v))
		newPos := (currPos + key) % 26
		result += string(alpha[newPos])
	}

	return result
}

func main() {
	tests := []struct {
		str  string
		key  int
		want string
	}{
		{str: "xyz", key: 2, want: "zab"},
	}

	for _, tc := range tests {
		got := caesarCipherEncryptor(tc.str, tc.key)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAILED! got: %s, want: %s", got, tc.want)
		} else {
			fmt.Printf("PASSED! got: %s, want: %s", got, tc.want)
		}
	}
}
