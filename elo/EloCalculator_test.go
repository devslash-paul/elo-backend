package elo

import "testing"

func TestEloFromStartingPosition(t *testing.T) {
	res := CalculateEloFor(EloCalculationRequest{
		winnerElo: 1000,
		loserElo:  1000,
		kvalue:    32,
	})

	if res.change != 16 {
		t.Errorf("Expected res.change to be 16, but it is %d", res.change)
	}
}

func TestLosingAtZeroEloDoesGoNegative(t *testing.T) {
	res := CalculateEloFor(EloCalculationRequest{
		winnerElo: 100,
		loserElo:  0,
		kvalue:    32,
	})

	if res.loserElo != -16 {
		t.Errorf("Elo should be -16 but was %d", res.loserElo)
	}
}

func TestCannotLoseMoreThanKValue(t *testing.T) {
	const kval = 32
	res := CalculateEloFor(EloCalculationRequest{
		winnerElo: 0,
		loserElo:  10000,
		kvalue:    kval,
	})

	if res.change != kval {
		t.Errorf("K value should be at max (%d) but was (%d)", kval, res.change)
	}
}

func TestBothNegativeWorks(t *testing.T) {
	res := CalculateEloFor(EloCalculationRequest{
		winnerElo: -100,
		loserElo:  -100,
		kvalue:    32,
	})

	if res.change != 16 {
		t.Errorf("Change should be 16 still for negative values, was %d", res.change)
	}
}

func TestChangeAppliedCorrectlyToBoth(t *testing.T) {
	res := CalculateEloFor(EloCalculationRequest{
		winnerElo: 1000,
		loserElo:  1000,
		kvalue:    32,
	})

	if res.loserElo != 984 || res.winnerElo != 1016 {
		t.Errorf("Elo should have added to 984 and 1016, instead was (%d)/(%d)",
			res.loserElo, res.winnerElo)
	}
}

func TestWinnerReallyShouldHaveWin(t *testing.T) {
	res := CalculateEloFor(EloCalculationRequest{
		winnerElo: 1000000000,
		loserElo:  -100000000,
		kvalue:    32,
	})

	if res.change != 0 {
		t.Errorf("Elo should not change for laaaarge differences, instead was %d", res.change)
	}
}
