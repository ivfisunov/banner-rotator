package ucb

import (
	"math"
)

type BannerID int

type BannerStat struct {
	Trials int
	Reward int
}

func MakeDecision(stat map[BannerID]BannerStat, allTrials int) BannerID {
	decisions := make(map[float64]BannerID)
	for k, v := range stat {
		avReward := float64(v.Reward) / float64(v.Trials)
		decision := avReward + math.Sqrt(2*math.Log(float64(allTrials))/float64(v.Trials))
		decisions[decision] = k
	}
	max := findMax(decisions)

	return decisions[max]
}

func findMax(decisions map[float64]BannerID) float64 {
	max := 0.0
	for k := range decisions {
		if k > max {
			max = k
		}
	}

	return max
}
