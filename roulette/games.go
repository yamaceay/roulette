package roulette

import (
	"fmt"
)

type Games []Game

func (games *Games) Play() (Results, error) {
	if games == nil {
		return nil, fmt.Errorf("NoGame")
	}
	var resultsList [][]Bet
	for _, game := range *games {
		results, err := game.Play()
		if err == nil {
			resultsList = append(resultsList, results)
		}
	}
	resultsMap := make(map[float64]float64)
	for _, results := range Prod(resultsList...) {
		prob := 1.
		wage := 0.
		for _, bet := range results {
			wage += bet.Wage
			prob *= bet.Prob
		}
		resultsMap[wage] += prob
	}
	return mapToBets(resultsMap), nil
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
