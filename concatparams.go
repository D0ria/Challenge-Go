package piscine

func ConcatParams(args []string) string {
	var phrase string
	for a, i := range args {
		phrase += i
		if a == len(args)-1 {
			continue
		} else {
			phrase += "\n"
		}
	}
	return phrase
}
