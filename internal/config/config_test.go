package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	t.Run("successfuly read config file", func(t *testing.T) {
		_, err := NewConfig("../../configs/config.toml")
		require.NoError(t, err)
	})

	t.Run("fail read config", func(t *testing.T) {
		_, err := NewConfig("./config/conf.toml")
		require.Error(t, err)
	})
}
