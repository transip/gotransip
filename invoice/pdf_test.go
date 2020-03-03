package invoice

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestPdf_GetIoReader(t *testing.T) {
	pdf := Pdf{Content: "dGVzdDEyMw=="}
	reader := pdf.GetIoReader()

	bytes, err := ioutil.ReadAll(reader)
	require.NoError(t, err)

	assert.Equal(t, []byte("test123"), bytes)
}
