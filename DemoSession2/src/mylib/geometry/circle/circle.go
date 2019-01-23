package circle

import "math"
func Dt(r float32) float32 {
	return math.Pi * r * r
}

func Cv(r float32) float32 {
	return 2 * math.Pi *r
}
