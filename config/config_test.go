package config

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func clear() {
	configOnce = &sync.Once{}
}

func TestConfigIntegration(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		clear()
		assert.Equal(t, NewConfig(), NewConfig())
	})
}
