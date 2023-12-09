package main

import "testing"

func Test_rank_card(t *testing.T) {
	tests := []struct {
		name string
		hand string
		want int
	}{
		{"pair", "32T3K", 2 },
		{"full", "T55T5", 5 },
		{"two pair", "KK677", 3 },
		{"two pair", "KTJJT", 3 },
		{"three", "QQQJA", 4 },
		{"odds", "12345", 1 },
		{"four", "33332", 6 },
		{"four", "2AAAA", 6 },
		{"five", "AAAAA", 7 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rank_card(card_hand{Cards: tt.hand}); got != tt.want {
				t.Errorf("rank_card() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rank_card_joker(t *testing.T) {
	type args struct {
		hand card_hand
	}
	tests := []struct {
		name string
		hand string
		want int
	}{
		{"no joker", "32T3K", 2},
		{"2 pair", "KK677", 3},
		{"4 kind", "T55J5", 6},
		{"4 kind", "KTJJT", 6},
		{"4 kind", "QQQJA", 6},
		{"jacks", "JJ8JJ", 7},
		{"alljacks", "JJJJJ", 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rank_card_joker(card_hand{Cards: tt.hand}); got != tt.want {
				t.Errorf("rank_card_joker() = %v, want %v", got, tt.want)
			}
		})
	}
}