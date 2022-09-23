package piscine

func ConcatParams(args []string) string {
	var phrase string
	for _, i := range args {
		phrase += i + "\n"
	}
	return phrase + "\n"
}
