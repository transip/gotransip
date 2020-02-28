package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusCodeMatches(t *testing.T) {
	assert.Equal(t, false, GetRestMethod.StatusCodeIsCorrect(500))
	assert.Equal(t, true, GetRestMethod.StatusCodeIsCorrect(200))

	assert.Equal(t, false, PostRestMethod.StatusCodeIsCorrect(500))
	assert.Equal(t, true, PostRestMethod.StatusCodeIsCorrect(201))

	assert.Equal(t, false, PutRestMethod.StatusCodeIsCorrect(500))
	assert.Equal(t, true, PutRestMethod.StatusCodeIsCorrect(204))

	assert.Equal(t, false, DeleteRestMethod.StatusCodeIsCorrect(500))
	assert.Equal(t, true, DeleteRestMethod.StatusCodeIsCorrect(204))
}
