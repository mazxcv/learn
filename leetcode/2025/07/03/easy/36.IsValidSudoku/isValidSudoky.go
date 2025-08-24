package isvalidsudoku

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}

	columnUnique := [9]map[byte]bool{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	quadreUnique := [9]map[byte]bool{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	for i, row := range board {
		if len(row) != 9 {
			return false
		}
		rowUnique := map[byte]bool{}
		for j, v := range row {
			if v > '9' || v < '1' {
				continue
			}
			// check row
			if _, ok := rowUnique[v]; !ok {
				rowUnique[v] = true
			} else {
				return false
			}

			// check column
			if _, ok := columnUnique[j][v]; !ok {
				columnUnique[j][v] = true
			} else {
				return false
			}

			// check quadre
			index := i/3*3 + j/3
			if _, ok := quadreUnique[index][v]; !ok {
				quadreUnique[index][v] = true
			} else {
				return false
			}
		}
	}

	return true
}
