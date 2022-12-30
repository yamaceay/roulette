package roulette

import (
	"fmt"
	"math"
)

type Results []Bet

func (r Results) String() string {
	output := "{\n"
	for i, bet := range r {
		output += fmt.Sprintf("  %s", bet)
		if i < len(r)-1 {
			output += ","
		}
		output += "\n"
	}
	output += "}"
	return output
}

func (r Results) Stats() string {
	return fmt.Sprintf("Mean: %.1f, Std: %.1f", r.Mean(), r.Std())
}

func (r Results) Mean() float64 {
	var mean float64
	for _, bet := range r {
		mean += bet.Prob * bet.Wage
	}
	return mean
}

func (r Results) Std() float64 {
	mean := r.Mean()
	var squaredR Results
	for _, bet := range r {
		squaredBet := Bet{Prob: bet.Prob, Wage: bet.Wage * bet.Wage}
		squaredR = append(squaredR, squaredBet)
	}
	return math.Sqrt(squaredR.Mean() - mean*mean)
}
