package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("successful open log file", func(t *testing.T) {
		_, err := New("development", "../../logs/logfile.log", "info")
		require.NoError(t, err)
	})

	t.Run("no such file or directory", func(t *testing.T) {
		_, err := New("development", "./logs/log.log", "info")
		require.Error(t, err)
	})

	t.Run("invalid log lovel", func(t *testing.T) {
		_, err := New("development", "../../logs/logfile.log", "bad_level")
		require.Error(t, err)
	})
}
