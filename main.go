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

	results, resultsEach, err := games.Play()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("------------------------------------")
		fmt.Println("Possible Events")
		fmt.Println("------------------------------------")
		fmt.Println("p(<wage>) = <prob>")
		fmt.Println(results)
		fmt.Println("------------------------------------")
		fmt.Println("Descriptive Stats (Summary)")
		fmt.Println("------------------------------------")
		fmt.Println(results.Stats())
		fmt.Println("------------------------------------")
		fmt.Println("Descriptive Stats (Individual)")
		fmt.Println("------------------------------------")
		for i, eachResult := range resultsEach {
			stats, err := eachResult.Stats()
			if err != nil {
				fmt.Printf("[%d] - %s\n\n", i, err)
			} else {
				fmt.Printf("[%d] \n%s\n", i, stats)
			}
		}
	}
}
