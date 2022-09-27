package main

func main() {
	s := "Viande"
	firstrune(s)
}

func firstrune(s string) rune {
	prem_lettre := []rune(s)
	return prem_lettre[0]

}
