package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusCodeMatches(t *testing.T) {
	method := PutRestMethod

	assert.Equal(t, false, method.StatusCodeIsCorrect(500))
	assert.Equal(t, true, method.StatusCodeIsCorrect(204))
}
