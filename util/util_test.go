package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, Contains([]int{1, 2, 3, 4, 5}, 10))
}
