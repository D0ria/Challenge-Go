package main

func main() {
	Rot14("Hello! How are You?")
}

func Rot14(s string) string {
	res := ""
	for _, i := range s {
		if i > 65 && i < 90 || i > 97 && i > 122 {
			i += 14
		}
		res += string(i)
	}
	return res
}
