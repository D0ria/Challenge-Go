package piscine

func Join(strs []string, sep string) string {
	var mot string
	long := len(strs)
	var cpt int
	for _, i := range strs {
		mot += i
		if cpt < long-1 {
			mot += sep
			cpt += 1
		}
	}
	return mot
}
