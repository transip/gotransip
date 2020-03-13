package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusCodeMatches(t *testing.T) {
	assert.Equal(t, false, GetMethod.StatusCodeOK(500))
	assert.Equal(t, true, GetMethod.StatusCodeOK(200))

	assert.Equal(t, false, PostMethod.StatusCodeOK(500))
	assert.Equal(t, true, PostMethod.StatusCodeOK(201))

	assert.Equal(t, false, PutMethod.StatusCodeOK(500))
	assert.Equal(t, true, PutMethod.StatusCodeOK(204))

	assert.Equal(t, false, DeleteMethod.StatusCodeOK(500))
	assert.Equal(t, true, DeleteMethod.StatusCodeOK(204))
}

func TestContains(t *testing.T) {
	assert.True(t, contains([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, contains([]int{1, 2, 3, 4, 5}, 10))
}
