package invoice

import (
	"encoding/base64"
	"io"
	"strings"
)

// GetIoReader will return a io.Reader containing the decoded base64 contents of the invoice PDF
func (p *Pdf) GetIoReader() io.Reader {
	reader := strings.NewReader(p.Content)

	return base64.NewDecoder(base64.StdEncoding, reader)
}
