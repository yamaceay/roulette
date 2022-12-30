package lib

import (
	"fmt"
	"math"
)

type Games []Game

func wageL1(wages []float64) float64 {
	var wageSum float64
	for _, wage := range wages {
		wageSum += wage
	}
	return wageSum
}

func wageL2(wages []float64) float64 {
	var wageSum float64
	for _, wage := range wages {
		wageSum += wage * wage
	}
	return math.Sqrt(wageSum)
}

func (games *Games) Play() (*Results, []*Results, error) {
	if games == nil {
		return nil, nil, fmt.Errorf("NoGame")
	}
	var wages []float64
	var betsList [][]*Bet
	var resultsList []*Results
	for _, game := range *games {
		wages = append(wages, game.Wage)
		results, err := game.Play()
		var bets []*Bet
		if err == nil {
			bets = results.Bets
		} else {
			results = &Results{Err: err}
		}
		betsList = append(betsList, bets)
		resultsList = append(resultsList, results)
	}

	betsMap := make(map[float64]float64)
	for _, bets := range Prod(betsList...) {
		prob := 1.
		wage := 0.
		for _, bet := range bets {
			if bet == nil {
				continue
			}
			wage += bet.Wage
			prob *= bet.Prob
		}
		betsMap[wage] += prob
	}

	return &Results{
		Bets:   mapToBets(betsMap),
		WageL1: wageL1(wages),
		WageL2: wageL2(wages),
	}, resultsList, nil
}

func NewGames(options []GameOptions) *Games {
	var games Games
	for _, option := range options {
		game := NewGame(option)
		if game != nil {
			games = append(games, *game)
		}
	}
	return &games
}
