package math

/*
Average exported comment here
*/
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
