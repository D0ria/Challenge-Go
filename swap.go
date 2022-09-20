package piscine

func Swap(a *int, b *int) {
	stock := *a
	*a = *b
	*b = stock

}
