package invoice

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPdf_GetIoReader(t *testing.T) {
	pdf := Pdf{Content: "dGVzdDEyMw=="}
	reader := pdf.GetReader()

	bytes, err := io.ReadAll(reader)
	require.NoError(t, err)

	assert.Equal(t, []byte("test123"), bytes)
}
