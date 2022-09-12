package piscine

func IsNumeric(s string) bool {
	bool := true
	for _, val := range s {
		if val >= '0' && val <= '9' {
			bool = true
		} else {
			bool = false
		}
	}
	return bool
}
