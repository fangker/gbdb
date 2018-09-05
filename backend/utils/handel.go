package utils

func IndexOfStringArray(sa []string, s string) int {
	for i, v := range sa {
		if v == s {
			return i
		}
	}
	return -1
}
