package main

import (
	"fmt"
	"math"
	"strings"
)



func parse_card(card string) (winners, numbers []int)  {
	cardPlusData := strings.Split(card, ":")
	winPlusNum := strings.Split(cardPlusData[1], "|")
	return StringToIntArray(winPlusNum[0], " "), StringToIntArray(winPlusNum[1], " ")
}

func day04a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var total = 0
	for _, line := range lines {
		winners, numbers := parse_card(line)

		var winningNumbers = 0
		for _, n := range numbers {
			// no Contains function for array so do by hand
			for _, w := range winners {
				if n == w {
					winningNumbers ++
				}
			}
		}
		if winningNumbers > 0 {
			// no int power function, use the float version
			total += int(math.Pow(2, float64(winningNumbers-1)))
		}
	}
	fmt.Printf("result %d\n", total)
}

func day04b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// calculate the scores of the cards (for later)
	scores := make([]int, len(lines))
	for idx, line := range lines {
		winners, numbers := parse_card(line)

		var winningNumbers = 0
		for _, n := range numbers {
			// no Contains function for array so do by hand
			for _, w := range winners {
				if n == w {
					winningNumbers++
				}
			}
		}
		scores[idx] = winningNumbers
	}
	fmt.Printf("scores %v\n", scores)

	// do the dance of winning more cards
	numCards := make([]int, len(lines))
	for idx := range numCards {
		numCards[idx] = 1
	}
	fmt.Printf("numCards %v\n", numCards)
	// if scores[0]=4 then we increment scores[1,2,3,4]
	ln := len(numCards)
	for idx, score := range scores {
		numCard := numCards[idx]
		for idx2 := idx + 1; idx2 < idx+score+1; idx2++ {
			if idx2 < ln {
				numCards[idx2] += numCard
			}
		}
	}
	fmt.Printf("numCards post %v\n", numCards)

	// sum the total
	var total = 0
	for _, nc := range numCards {
		total += nc
	}
	fmt.Printf("total %d\n", total)
}
