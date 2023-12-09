package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type card_hand struct {
	Cards string
	Bid int
}

func parse_card_hand(filename string) []card_hand  {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}

	var cards []card_hand = nil
	for _, line := range lines {
		arr:=strings.Split(line, " ")
		bid, _ := strconv.Atoi(arr[1])
		cards = append(cards, card_hand{
			Cards: arr[0],
			Bid:   bid,
		})
	}

	return cards
}

// returns the cards rank
// 5 of kind is best (7)
// high card is worst (1)
func rank_card(hand card_hand) int  {
	// scores
	score_5 := 7
	score_4 := 6
	score_full := 5
	score_3 := 4
	score_pairs := 3
	score_2 := 2
	score_1 := 1

	// count card freq
	freq := make(map[rune]int)
	for _,c := range hand.Cards {
		freq[c] = strings.Count(hand.Cards, string(c))
	}
	//fmt.Printf("hand %s freq %v\n", hand.Cards, freq)

	current :=0
	// count card freq
	for _,f := range freq {
		switch f {
		case 5:
			return score_5
		case 4:
			return score_4
		case 3:
			// if there is a pair, then this is a full house
			if current == score_2 {
				return score_full
			}
			// looks like a 3 of a kind
			current = score_3
		case 2:
			// if there is a 3, then its a full house
			if current == score_3 {
				return score_full
			}
			// if already a pair, then its two pairs
			if current == score_2 {
				return score_pairs
			}
			current = score_2
		case 1:
			if current == 0 {
				current = score_1
			}
		}
	}
	return current
}

// returns true is hand1 is better than hand2
func compare_cards(hand1, hand2 card_hand) bool  {
	rank1:= rank_card(hand1)
	rank2:= rank_card(hand2)
	if rank1 != rank2 {
		return rank1 < rank2
	}
	// compare by card
	card_value := "23456789TJQKA"
	for i := range hand1.Cards {
		rank1 := strings.IndexByte(card_value, hand1.Cards[i])
		rank2 := strings.IndexByte(card_value, hand2.Cards[i])
		if rank1 != rank2 {
			return rank1 < rank2
		}
	}
	// seems they are the same
	return false
}

func day07a(filename string)  {
	hands := parse_card_hand(filename)

	fmt.Printf("hands %+v\n", hands)
	// test the go-sort functions
	sort.Slice(hands, func(i, j int) bool {
		return compare_cards(hands[i],hands[j])
	})
	fmt.Printf("sorted hands %+v\n", hands)
	score :=0
	for i,hand := range hands {
		score += hand.Bid * (i+1)
	}
	fmt.Printf("final score %d", score)
}

// returns the cards rank with joker rules
// 5 of kind is best (7)
// high card is worst (1)
func rank_card_joker(hand card_hand) int {
	// scores
	score_5 := 7
	score_4 := 6
	score_full := 5
	score_3 := 4
	score_pairs := 3
	score_2 := 2
	score_1 := 1

	// count card freq
	freq := make(map[rune]int)
	for _, c := range hand.Cards {
		freq[c] = strings.Count(hand.Cards, string(c))
	}
	joker, found := freq['J']
	if found == false {
		joker = 0
	}
	// remove the jokers from the calculations
	delete(freq, 'J')
	//fmt.Printf("hand %s freq %v jokers %d\n", hand.Cards, freq, joker)

	current := 0
	// count card freq
	for _, f := range freq {
		switch f {
		case 5:
			current = Max(current,score_5)
		case 4:
			current = Max(current,score_4)
		case 3:
			// if there is a pair, then this is a full house
			if current == score_2 {
				current = Max(current,score_full)
			}
			// looks like a 3 of a kind
			current = Max(current,score_3)
		case 2:
			// if there is a 3, then its a full house
			if current == score_3 {
				current = Max(current,score_full)
			}
			// if already a pair, then its two pairs
			if current == score_2 {
				current = Max(current,score_pairs)
			}
			current = Max(current, score_2)
		case 1:
			current = Max(current, score_1)
		}
	}
	//fmt.Printf("hand %s base score %d jokers %d\n", hand.Cards, current, joker)
	// now consider jokers
	switch current {
	case score_4:
		if joker == 1 {
			current = score_5
		}
	case score_3:
		if joker == 2 {
			current = score_5
		} else if joker == 1 {
			current = score_4
		}
	case score_pairs:
		if joker == 1 {
			current = score_full
		}
	case score_2:
		if joker == 3 {
			current = score_5
		} else if joker == 2 {
			current = score_4
		} else if joker == 1 {
			current = score_3
		}
	case score_1, 0:
		if joker >= 4 {
			current = score_5
		} else if joker == 3 {
			current = score_4
		} else if joker == 2 {
			current = score_3
		} else if joker == 1 {
			current = score_2
		}
	}
	//fmt.Printf("hand %s final score %d\n", hand.Cards, current)
	return current
}

// returns true is hand1 is better than hand2
func compare_cards_joker(hand1, hand2 card_hand) bool  {
	rank1:= rank_card_joker(hand1)
	rank2:= rank_card_joker(hand2)
	if rank1 != rank2 {
		return rank1 < rank2
	}
	// compare by card
	card_value := "J23456789TQKA"
	for i := range hand1.Cards {
		rank1 := strings.IndexByte(card_value, hand1.Cards[i])
		rank2 := strings.IndexByte(card_value, hand2.Cards[i])
		if rank1 != rank2 {
			return rank1 < rank2
		}
	}
	// seems they are the same
	return false
}

func day07b(filename string)  {
	hands := parse_card_hand(filename)

	fmt.Printf("hands %+v\n\n", hands)
	// test the go-sort functions
	sort.Slice(hands, func(i, j int) bool {
		return compare_cards_joker(hands[i],hands[j])
	})
	fmt.Printf("sorted hands %+v\n\n", hands)
	score :=0
	for i,hand := range hands {
		score += hand.Bid * (i+1)
	}
	fmt.Printf("final score %d", score)

	// algol failed: 249096602 is too low
	// 249137934 too low
	// 249390065 too low (i missed pairs=>full house)
	// final bug: I didn't consider single card+4 jokers or 5 jokers
}
