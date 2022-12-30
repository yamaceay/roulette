package lib

import (
	"fmt"
	"math"
)

type Results struct {
	Bets []*Bet
	Wage float64

	Err error

	WageL1 float64
	WageL2 float64
}

func (r Results) String() string {
	var output string
	for _, bet := range r.Bets {
		output += fmt.Sprintf("%s\n", bet)
	}
	return output
}

func (r Results) Stats() (string, error) {
	if r.Err != nil {
		return "", fmt.Errorf("NoResults: %w", r.Err)
	}
	output := fmt.Sprintf("Mean: %.1f\n", r.Mean())
	if r.Wage != 0 {
		output += fmt.Sprintf("Std: %.1f\n", r.Std())
		output += fmt.Sprintf("RelStd: %.1f\n", r.RelStd())
	} else {
		output += fmt.Sprintf("StdManh: %.1f\n", r.StdL1())
		output += fmt.Sprintf("StdEucl: %.1f\n", r.StdL2())
	}
	return output, nil
}

func (r Results) Mean() float64 {
	var mean float64
	for _, bet := range r.Bets {
		mean += bet.Prob * bet.Wage
	}
	return mean
}

func (r Results) Std() float64 {
	mean := r.Mean()
	var squaredR Results
	for _, bet := range r.Bets {
		squaredBet := Bet{
			Prob: bet.Prob,
			Wage: bet.Wage * bet.Wage,
		}
		squaredR.Bets = append(squaredR.Bets, &squaredBet)
	}
	return math.Sqrt(squaredR.Mean() - mean*mean)
}

func (r Results) RelStd() float64 {
	return r.Std() / r.Wage
}

func (r Results) StdL1() float64 {
	return r.Std() / r.WageL1
}

func (r Results) StdL2() float64 {
	return r.Std() / r.WageL2
}
