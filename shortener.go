package clip

func Diff(original, shorted string) float64 {
	n := len(original)
	m := len(shorted)

	return float64(n-m) * 100 / float64(n)

}

func Shortit(url string) string {
	return "test"

}
