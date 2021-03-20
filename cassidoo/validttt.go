package puzzles

func ValidTTTPosition(board []string) bool {
	var xCount, oCount int

	verticalBoard := make([]string, 3)
	diagonals := make([]string, 2)
	for i, v := range board {

		for _, c := range v {
			if c == 'X' {
				xCount++
			} else if c == 'O' {
				oCount++
			}
		}

		//Build vertical
		verticalBoard[0] += string(v[0])
		verticalBoard[1] += string(v[1])
		verticalBoard[2] += string(v[2])

		//Build Diagonals
		diagonals[0] += string(v[i])
		diagonals[1] += string(v[2-i])
	}
	allVariants := append(board, verticalBoard...)
	allVariants = append(allVariants, diagonals...)

	//Game Over situations

	if xCount+oCount == 9 {
		return false
	}

	for _, v := range allVariants {
		if v == "XXX" {
			return false
		} else if v == "OOO" {
			return false
		}
	}

	//If X plays first, then no of X == O or X == O + 1, otherwise it's invalid

	if xCount == oCount || xCount == oCount+1 {
		return true
	}

	return false

}
