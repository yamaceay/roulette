package roulette

import (
	"fmt"
)

type GameOptions struct {
	Prob          float64 `json:"prob"`
	Wage          float64 `json:"wage"`
	StepFuncProxy string  `json:"stepFunc"`
	StopLoss      float64 `json:"stopLoss"`
	WinRound      int     `json:"winRound"`
}

func NewGame(options GameOptions) *Game {
	var stepFunc StepFunc
	switch options.StepFuncProxy {
	case "two":
		stepFunc = Two
	case "fib":
		stepFunc = Fib
	default:
		return nil
	}
	return &Game{
		Bet: Bet{
			Prob: options.Prob,
			Wage: options.Wage,
		},
		History: History{
			options.Wage,
		},
		WinRound:  options.WinRound,
		StartWage: options.Wage,
		StepFunc:  stepFunc,
		StopLoss:  options.StopLoss,
	}
}

type Game struct {
	Bet       `json:"bet"`
	History   `json:"history"`
	WinRound  int     `json:"winRound"`
	StartWage float64 `json:"startWage"`
	StepFunc  `json:"stepFunc"`
	StopLoss  float64 `json:"stopLoss"`
}

func (g *Game) Play() (Results, error) {
	p, invP := g.Prob, 1-g.Prob
	tempFailProb := 1.
	bets := make(map[float64]float64)
	for i := 0; i < g.WinRound; i++ {
		tempFailProb *= p
		saldoAfterWin := g.Earn() + g.Wage + g.Saldo()
		bets[saldoAfterWin] += tempFailProb

		tempFailProb *= invP / p
		if err := g.Step(); err != nil {
			return nil, fmt.Errorf("GameStopped: %w", err)
		}
	}
	if err := g.Unstep(); err != nil {
		return nil, fmt.Errorf("InvalidWinRound: %w", err)
	}
	bets[g.Saldo()] += tempFailProb
	return mapToBets(bets), nil
}

func (g *Game) Step() error {
	g.Wage = g.StepFunc(g.History)
	g.History = append(g.History, g.Wage)
	if g.Saldo()+g.StopLoss >= 0 {
		return nil
	}
	g.Unstep()
	return fmt.Errorf("Bankruptcy")
}

func (g *Game) Unstep() error {
	if length := len(g.History); length > 0 {
		g.History = g.History[:length-1]
		g.Wage = g.History[length-2]
		return nil
	}
	return fmt.Errorf("EmptyHistory")
}
