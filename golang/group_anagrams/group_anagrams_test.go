package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat","tea","tan","ate","nat","bat"}
	got := groupAnagrams(input)
	want := [][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FAILED! got: %v, want: %v", got, want)
	}
}
