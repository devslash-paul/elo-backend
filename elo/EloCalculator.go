package elo

import (
	"math"
)

type EloCalculationRequest struct {
	winnerElo int
	loserElo  int
	kvalue    int
}

type EloCalculationResult struct {
	winnerElo int
	loserElo  int
	change    int
}

func CalculateEloFor(req EloCalculationRequest) EloCalculationResult {
	exponent := math.Pow(10, float64((req.winnerElo-req.loserElo)/400))
	ex := 1.0 / (1 + exponent)

	change := int(float64(req.kvalue) * ex)

	return EloCalculationResult{
		change:    change,
		winnerElo: req.winnerElo + change,
		loserElo:  req.loserElo - change,
	}
}
