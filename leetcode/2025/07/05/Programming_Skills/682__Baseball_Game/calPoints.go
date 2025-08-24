package _82__Baseball_Game

import "strconv"

func calPoints(operations []string) (result int) {
	var res []int
	for _, v := range operations {
		switch v {
		case "D":
			res = append(res, res[len(res)-1]*2)
		case "C":
			res = res[:len(res)-1]
		case "+":
			res = append(res, res[len(res)-1]+res[len(res)-2])
		default:
			c, _ := strconv.Atoi(v)
			res = append(res, c)
		}
	}
	for _, v := range res {
		result += v
	}
	return result
}
