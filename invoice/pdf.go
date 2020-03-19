package invoice

import (
	"encoding/base64"
	"io"
	"strings"
)

// Pdf struct for a invoice as Pdf
type Pdf struct {
	// Content contains a base64 encoded string containing the pdf content
	// We provide a io.Reader on this struct called GetReader
	Content string `json:"pdf"`
}

// GetReader will return a io.Reader containing the decoded base64 contents of the invoice PDF
func (p *Pdf) GetReader() io.Reader {
	reader := strings.NewReader(p.Content)

	return base64.NewDecoder(base64.StdEncoding, reader)
}
