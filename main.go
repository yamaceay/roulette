package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/roulette/roulette"
)

func readGames() (roulette.Games, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("InvalidFilename: %w", err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("InvalidFile: %w", err)
	}
	var options []roulette.GameOptions
	if err := json.Unmarshal(bytes, &options); err != nil {
		return nil, fmt.Errorf("InvalidUnmarshal: %w", err)
	}
	games := roulette.NewGames(options)
	if games == nil {
		return nil, fmt.Errorf("EmptyGames: %w", err)
	}
	return *games, nil
}

func main() {
	games, err := readGames()
	if err != nil {
		panic(err)
	}

	results, err := games.Play()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(results)
		fmt.Println(results.Stats())
	}
}
