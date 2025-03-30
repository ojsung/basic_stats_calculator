package main

import (
	"math/big"
	"fmt"
	"github.com/ojsung/basic_stats_calculator/pkg/calculator"
)

func main() {
	chanceOfSuccess, isRead := new(big.Float).SetString("0.20")
	if !isRead {
		fmt.Sprintln("Failed to convert", chanceOfSuccess, "to string")
		return;
	}
	formattedChance := new(big.Float)
	formattedChance.Mul(chanceOfSuccess, big.NewFloat(100.0))
	var trials int64 = 20
	var successes int64 = 5
	probability, probabilityErr := calculator.CalculateBinomialProbability(chanceOfSuccess, trials, successes)
	if probabilityErr != nil {
		panic(probabilityErr)
	}
	fmt.Sprintln("With a chance of", formattedChance, "the chance of", successes, "successes in", trials, "trials is", probability)
}