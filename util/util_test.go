package util

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyValueXML(t *testing.T) {
	body := []byte(`<transip><item><item><key xsi:type="xsd:string">startDate</key><value xsi:type="xsd:string">2018-08-01</value></item><item><key xsi:type="xsd:string">endDate</key><value xsi:type="xsd:string">2018-09-01</value></item><item><key xsi:type="xsd:string">usedInBytes</key><value xsi:type="xsd:string">15253974528</value></item><item><key xsi:type="xsd:string">usedTotalBytes</key><value xsi:type="xsd:string">21152084384</value></item><item><key xsi:type="xsd:string">maxBytes</key><value xsi:type="xsd:string">5368709120000</value></item></item></transip>`)

	var kvx KeyValueXML
	err := xml.Unmarshal(body, &kvx)
	require.NoError(t, err)
	assert.Equal(t, 1, len(kvx.Cont))
	assert.Equal(t, 5, len(kvx.Cont[0].Item))

	k := []string{"startDate", "endDate", "usedInBytes", "usedTotalBytes", "maxBytes"}
	v := []string{"2018-08-01", "2018-09-01", "15253974528", "21152084384", "5368709120000"}
	for i := 0; i < len(kvx.Cont[0].Item); i++ {
		assert.Equal(t, k[i], kvx.Cont[0].Item[i].Key)
		assert.Equal(t, v[i], kvx.Cont[0].Item[i].Value)
	}
}

func TestXMLTime(t *testing.T) {
	body := []byte(`<item>
			<date>2018-09-16</date>
			<time>2018-09-06 14:03:24</time>
		</item>`)

	var v struct {
		Date XMLTime `xml:"date"`
		Time XMLTime `xml:"time"`
	}

	err := xml.Unmarshal(body, &v)
	require.NoError(t, err)

	x, _ := time.Parse("2006-01-02 15:04:05", "2018-09-16 00:00:00")
	assert.Equal(t, x, v.Date.Time)
	x, _ = time.Parse("2006-01-02 15:04:05", "2018-09-06 14:03:24")
	assert.Equal(t, x, v.Time.Time)
}
