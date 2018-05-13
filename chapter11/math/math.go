package math

// Average returns the mean of an array of floats
func Average(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}
