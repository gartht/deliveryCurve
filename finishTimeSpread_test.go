package finishCurve

import "testing"

var cases = map[int]int{
	4:  1,
	5:  1,
	6:  2,
	7:  4,
	8:  10,
	9:  10,
	11: 12,
	13: 9,
	14: 8,
	15: 6,
	16: 7,
	17: 7,
	18: 5,
	19: 4,
	21: 4,
	22: 2,
	24: 2,
	25: 2,
	26: 1,
	30: 1,
	32: 1,
	35: 1,
}

func TestCompletedWorkCount(t *testing.T) {

	workDone := GetCompletedWorkCount(cases)
	if workDone != 100 {
		t.Errorf("GetCompletedWork returned %q expected %q", workDone, 100)
	}
}

func TestGetPercentageChanceOfCompletionBasedOnDays(t *testing.T) {
	chanceComplete := GetPercentChanceOfCompletion(11, cases)
	if chanceComplete != 40.0 {
		t.Errorf("GetChanceOfCompletion(%d, workHistory) = %b expected %d \n", 11, chanceComplete, 40)
	}

	chanceComplete = GetPercentChanceOfCompletion(16, cases)

	if chanceComplete != 70.0 {
		t.Errorf("GetChanceOfCompletion(%d, workHistory) = %b expected %d \n", 16, chanceComplete, 70)
	}

	chanceComplete = GetPercentChanceOfCompletion(2, cases)

	if chanceComplete != 0.0 {
		t.Errorf("Snowballs chance failed")
	}
}
