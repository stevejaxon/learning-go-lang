package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	newPop := p0
	percent = percent / 100
	for newPop < p {
	  newPop = calculateNewPopulation(newPop, aug, percent)
	  years++
	}
	return years
  }
  
  func calculateNewPopulation(startPop, netMigration int, percentageIncr float64) int {
	return startPop + int(float64(startPop) * percentageIncr) + netMigration
  }