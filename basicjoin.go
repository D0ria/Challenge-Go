package piscine

func BasicJoin(elems []string) string {
	var phrase string
	for _, i := range elems {
		phrase += i
	}
	return phrase
}
