package kata

import "math"

// Solution to the Kata https://www.codewars.com/kata/going-to-the-cinema/train/go - see for detailed explination
// John wants to know how many times he must go to the cinema so that the final result of System B, when rounded up to the next dollar, will be cheaper than System A.
func Movie(card, ticket int, perc float64) (cinemaVisits int) {  
	systemA := calculateTicketPrice(ticket)
	systemB := calculateCardPrice(card, ticket, perc)

	for roundUpToNearestDollar(systemB()) >= systemA() {
		cinemaVisits++
	}
	cinemaVisits++

  	return
}

func calculateCardPrice(card, ticket int, perc float64) func() float64 {
	upfrontCost := float64(card)
	incrementalCost := float64(ticket)
	var totalIncrementalCost float64 
	

	return func() float64 {
		nextTicketCost := incrementalCost * perc
		totalIncrementalCost += nextTicketCost
		incrementalCost = nextTicketCost
		return upfrontCost + totalIncrementalCost
	}
}

func calculateTicketPrice(ticket int) func() int {
	ticketPrice := ticket
	var incrementalCost int

	return func() int {
		incrementalCost += ticketPrice
		return incrementalCost
	}
}

func roundUpToNearestDollar(cost float64) int {
	return int(math.Ceil(cost))
}