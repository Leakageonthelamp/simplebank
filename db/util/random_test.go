package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	min := int64(10)
	max := int64(20)
	for i := 0; i < 100; i++ {
		n := RandomInt(min, max)
		require.True(t, n >= min && n <= max)
	}
}

func TestRandomString(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := RandomString(10)
		require.Len(t, s, 10)
	}
}
