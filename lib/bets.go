package lib

import (
	"fmt"
	"sort"
)

func mapToBets(bets map[float64]float64) []*Bet {
	var wages []float64
	for wage := range bets {
		wages = append(wages, wage)
	}
	sort.Slice(wages, func(i, j int) bool {
		return wages[i] < wages[j]
	})
	for i, j := 0, len(wages)-1; i < j; i, j = i+1, j-1 {
		wages[i], wages[j] = wages[j], wages[i]
	}

	var probs []*Bet
	for _, wage := range wages {
		prob := bets[wage]
		bet := Bet{Wage: wage, Prob: prob}
		probs = append(probs, &bet)
	}
	return probs
}

type Bet struct {
	Prob float64 `json:"prob"`
	Wage float64 `json:"wage"`
}

func (b Bet) String() string {
	return fmt.Sprintf("p(%.1f) = %.3f", b.Wage, b.Prob)
}

func (b *Bet) Earn() float64 {
	return (1 - b.Prob) / b.Prob * b.Wage
}
