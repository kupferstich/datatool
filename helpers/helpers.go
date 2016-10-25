package helpers

// StrInSlice checks if a string is in slice
func StrInSlice(str string, sl []string) bool {
	i := FindStrInSlice(str, sl)
	if i == -1 {
		return false
	}
	return true
}

// FindStrInSlice returns the index, where the string is in slice
// returns -1
func FindStrInSlice(str string, sl []string) int {
	for i, s := range sl {
		if str == s {
			return i
		}
	}
	return -1
}

// AddStrToSlice adds the string, just when the string is not inside
// the slice.
func AddStrToSlice(str string, sl []string) []string {
	if !StrInSlice(str, sl) {
		sl = append(sl, str)
	}
	return sl
}

// RemoveStrFromSlice removes the str from the sl
func RemoveStrFromSlice(str string, sl []string) []string {
	i := FindStrInSlice(str, sl)
	if i == -1 {
		return sl
	}
	sl = append(sl[:i], sl[i+1:]...)
	return sl
}
