package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlag(t *testing.T) {
	flag := encodeFlag("now", "user", "seed")
	when, user, err := decodeFlag(flag, "seed")
	require.NoError(t, err)
	assert.Equal(t, "now", when)
	assert.Equal(t, "user", user)
}
