package kata

func Points(games []string) int {
	scores := 0
	for _, game := range games {
		scores += calculatePoints(game)
	}
	return scores
}

/*
 * Business rules: 
 * The result of each match look like "x:y"
 * there are 10 matches in the championship
 * 0 <= x <= 4
 * 0 <= y <= 4
 */
func calculatePoints(gameScore string) (score int) {
	x := int(gameScore[0])
	y := int(gameScore[2])
	switch {
		case x == y:
			score = 1
		case x > y:
			score = 3
		case x < y:
			score = 0
	}
	return
}