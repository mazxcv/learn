package darts

func Score(x, y float64) int {
	t := x*x + y*y
	switch {
	case t > 100:
		return 0
	case t > 25:
		return 1
	case t > 1:
		return 5
	default:
		return 10
	}
}
