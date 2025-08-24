package _4__Spiral_Matrix

// Идея: делаем втупую, бегаем по периметру.
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	n := len(matrix)
	m := len(matrix[0])

	res := make([]int, len(matrix)*len(matrix[0]))
	way := "RDLU"
	iWay := 0
	x, y := 0, 0
	indent := 0
	for i := 0; i < len(res); i++ {
		res[i] = matrix[x][y]
		switch way[iWay] {
		case 'R':
			switch {
			case y == m-indent-1:
				iWay = (iWay + 1) % 4
				x++
			default:
				y++
			}
		case 'D':
			switch {
			case x == n-indent-1:
				iWay = (iWay + 1) % 4
				y--
			default:
				x++
			}
		case 'L':
			switch {
			case y == 0+indent:
				iWay = (iWay + 1) % 4
				x--
			default:
				y--
			}
		case 'U':
			switch {
			case x == 0+indent+1:
				iWay = (iWay + 1) % 4
				y++
				indent++
			default:
				x--
			}
		}
	}
	return res
}
