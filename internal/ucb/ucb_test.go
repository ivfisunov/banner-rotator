package ucb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeDesicion(t *testing.T) {
	stats := make(map[BannerID]BannerStat)
	stats[1] = BannerStat{Trials: 10, Reward: 0}
	stats[2] = BannerStat{Trials: 10, Reward: 0}
	stats[3] = BannerStat{Trials: 10, Reward: 0}

	// all banners have equal stats
	t.Run("maybe any banner", func(t *testing.T) {
		decision := MakeDecision(stats, 30)
		require.Contains(t, []BannerID{1, 2, 3}, decision)
	})

	// banner #1 has lots of clicks
	stats[1] = BannerStat{Trials: 100, Reward: 9999}
	t.Run("should be banner #1", func(t *testing.T) {
		decision := MakeDecision(stats, 55)
		require.Equal(t, 1, int(decision))
	})

	// banner #2 also has lots of clicks
	stats[2] = BannerStat{Trials: 100, Reward: 9999}
	t.Run("should be banner #1 or #2", func(t *testing.T) {
		decision := MakeDecision(stats, 55)
		require.Contains(t, []BannerID{1, 2}, decision)
	})

	// banner #3 has the largest number of clicks
	stats[3] = BannerStat{Trials: 100, Reward: 99999}
	t.Run("should be banner #1 or #2", func(t *testing.T) {
		decision := MakeDecision(stats, 55)
		require.Equal(t, 3, int(decision))
	})

	// if stats are not valid should be Zero banner
	stats[1] = BannerStat{Trials: 0, Reward: 0}
	stats[2] = BannerStat{Trials: 0, Reward: 0}
	stats[3] = BannerStat{Trials: 0, Reward: 0}
	t.Run("should be zero", func(t *testing.T) {
		decision := MakeDecision(stats, 0)
		require.Equal(t, 0, int(decision))
	})
}
