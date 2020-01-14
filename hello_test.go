package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, world."

	assert.Equal(t, expected, Hello())
}
