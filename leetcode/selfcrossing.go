package puzzles

// https://leetcode.com/problems/self-crossing/

type point struct {
	x int
	y int
}

func IsSelfCrossing(distance []int) bool {
	var pointsVisited []point
	currentPosition := point{0, 0}
	pointsVisited = append(pointsVisited, currentPosition)
	for i, v := range distance {
		switch i % 4 {
		case 0: // North
			for x := 0; x < v; x++ {
				currentPosition = point{
					currentPosition.x + 1,
					currentPosition.y,
				}
				if in(pointsVisited, currentPosition) {
					return true
				}
				pointsVisited = append(pointsVisited, currentPosition)
			}
		case 1: // West
			for x := 0; x < v; x++ {
				currentPosition = point{
					currentPosition.x,
					currentPosition.y - 1,
				}
				if in(pointsVisited, currentPosition) {
					return true
				}
				pointsVisited = append(pointsVisited, currentPosition)
			}
		case 2: // South
			for x := 0; x < v; x++ {
				currentPosition = point{
					currentPosition.x - 1,
					currentPosition.y,
				}
				if in(pointsVisited, currentPosition) {
					return true
				}
				pointsVisited = append(pointsVisited, currentPosition)
			}
		case 3: // East
			for x := 0; x < v; x++ {
				currentPosition = point{
					currentPosition.x,
					currentPosition.y + 1,
				}
				if in(pointsVisited, currentPosition) {
					return true
				}
				pointsVisited = append(pointsVisited, currentPosition)
			}
		}
	}
	return false
}

func in(pointsVisited []point, currentPosition point) bool {
	for _, point := range pointsVisited {
		if point.x == currentPosition.x && point.y == currentPosition.y {
			return true
		}
	}
	return false
}
