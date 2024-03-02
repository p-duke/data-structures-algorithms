func (s *Solution) IsAnagram(str1 string, str2 string) bool {
	// Convert strings to sorted slices of characters
	ss := strings.Split(str1, "")
	tt := strings.Split(str2, "")
	sort.Strings(ss)
	sort.Strings(tt)

	// Check if the sorted slices are equal
	return strings.Join(ss, "") == strings.Join(tt, "")
}

func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }

  count := make(map[string]int)
  for i, v := range s {
    sLetter := string(v)
    tLetter := string(t[i])
    if _, ok := count[sLetter]; ok {
      count[sLetter] += 1
    } else {
      count[sLetter] = 1
    }

    if _, ok := count[tLetter]; ok {
      count[tLetter] -= 1
    } else {
      count[tLetter] = -1
    }
  }

  for _, v := range count {
    if v != 0 {
      return false
    }
  }

  return true
}
