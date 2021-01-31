package ucb

import "math"

type bannerStat struct {
	trials int
	reward int
}

func MakeDecision(stat map[int]bannerStat) int {
	allTrials := 0
	for _, v := range stat {
		allTrials += v.trials
	}

	var decisions map[float64]int
	for k, v := range stat {
		avReward := float64(v.reward) / float64(v.trials)
		decision := avReward + math.Sqrt(2*math.Log(float64(allTrials))/float64(v.trials))
		decisions[decision] = k
	}
	max := findMax(decisions)

	return decisions[max]
}

func findMax(decisions map[float64]int) float64 {
	max := 0.0
	for k := range decisions {
		if k > max {
			max = k
		}
	}

	return max
}
