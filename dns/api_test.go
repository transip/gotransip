package dns

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip"
)

func TestGetDNSSecEntries(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getdnssecentries.xml")
	assert.NoError(t, err)

	dns, err := GetDnsSecEntries(c, "example.org")
	assert.NoError(t, err)
	assert.IsType(t, DNSSecEntries{}, dns)
	assert.Equal(t, 2, len(dns))
	assert.Equal(t, 1337, dns[0].KeyTag)
	assert.Equal(t, DNSSecFlagKSK, dns[0].Flags)
	assert.Equal(t, DNSSecAlgorithmRSASHA512, dns[0].Algorithm)
	assert.Equal(t, "emFpcmFvcHUzdm9vNGFpczlxdWVpN211dTVvc2g1QWVHaG9pY2kxaWVnaDZmYWk5ZWVtb2hwOWdhaU5nb29DZWlxdWFoZjNhaWdhaDJyYWhRdTFhaGxvaHNlaTh3ZW9zaDRhZXBoYWhzb29raTZFaWNoOGFpdnU5Y2llcGhvYWwK", dns[0].PublicKey)
	assert.Equal(t, 12, dns[1].KeyTag)
	assert.Equal(t, DNSSecFlagZSK, dns[1].Flags)
	assert.Equal(t, DNSSecAlgorithmECDSAP384SHA384, dns[1].Algorithm)
	assert.Equal(t, "dWl4YWl4MHBoZWVtN3lhcGhhaWIwYWhsYWVqMW9odzB1YThYaTFoYUJhaHBvOWhhZXNhaDJBaGQ4b2s4VGhvU2hhaWozc2hhaDluYWljYWljN2lvaG83aW9YZWRvb2w0YWhXYWl0bzNYZWlQaGFlNWVpZ2VpcGVlZzdhZXhpZTAK", dns[1].PublicKey)
}
