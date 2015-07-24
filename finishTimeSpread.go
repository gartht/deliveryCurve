package finishCurve

func GetCompletedWorkCount(finishTimes map[int]int) int {
	//first get the count of all completed work

	workDone := 0
	for _, doneWork := range finishTimes {
		workDone += doneWork
	}
	return workDone
}

func GetPercentChanceOfCompletion(daysOfWork int, finishTimeHistory map[int]int) float64 {
	workCount := float64(GetCompletedWorkCount(finishTimeHistory))
	subset := map[int]int{}
	for i := 1; i <= daysOfWork; i++ {
		subset[i] = finishTimeHistory[i]
	}
	workCompletedInSubset := float64(GetCompletedWorkCount(subset))

	percentChance := workCompletedInSubset / workCount

	return percentChance * 100

}
