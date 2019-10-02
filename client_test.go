package gotransip

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSOAPClient(t *testing.T) {
	var cc ClientConfig
	var err error

	// empty ClientConfig should raise error about missing AccountName
	_, err = NewSOAPClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("AccountName is required"), err)

	cc.AccountName = "foobar"
	// ClientConfig with only AccountName set should raise error about private keys
	_, err = NewSOAPClient(cc)
	require.Error(t, err)
	assert.Equal(t, errors.New("PrivateKeyPath or PrivateKeyBody is required"), err)

	cc.PrivateKeyPath = "/file/not/found"
	// ClientConfig with PrivateKeyPath set to nonexisting file should raise error
	_, err = NewSOAPClient(cc)
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("^could not open private key"), err.Error())

	// ClientConfig with PrivateKeyPath that does exist but is unreadable should raise
	// error
	// prepare tmpfile
	tmpFile, err := ioutil.TempFile("", "gotransip")
	assert.NoError(t, err)
	err = os.Chmod(tmpFile.Name(), 0000)
	assert.NoError(t, err)

	cc.PrivateKeyPath = tmpFile.Name()
	_, err = NewSOAPClient(cc)
	require.Error(t, err)
	assert.Regexp(t, regexp.MustCompile("permission denied"), err.Error())

	os.Remove(tmpFile.Name())
	cc.PrivateKeyPath = ""

	// ClientConfig with PrivateKeyBody set but no PrivateKeyPath should have
	// PrivateKeyBody as private key body
	cc.PrivateKeyBody = []byte{1, 2, 3, 4}
	c, err := NewSOAPClient(cc)
	assert.NoError(t, err)
	assert.Equal(t, cc.PrivateKeyBody, c.soapClient.PrivateKey)

	// Also, with no mode set, it should default to APIModeReadWrite
	assert.Equal(t, APIModeReadWrite, c.soapClient.Mode)

	// Override PrivateKeyBody with PrivateKeyPath
	pkBody := []byte{2, 3, 4, 5}
	// prepare tmpfile
	tmpFile, err = ioutil.TempFile("", "gotransip")
	assert.NoError(t, err)
	err = ioutil.WriteFile(tmpFile.Name(), []byte(pkBody), 0)
	assert.NoError(t, err)

	cc.PrivateKeyPath = tmpFile.Name()
	c, err = NewSOAPClient(cc)
	assert.NoError(t, err)
	assert.Equal(t, pkBody, c.soapClient.PrivateKey)

	os.Remove(tmpFile.Name())
	cc.PrivateKeyPath = ""

	// override API mode to APIModeReadOnly
	cc.Mode = APIModeReadOnly

	c, err = NewSOAPClient(cc)
	assert.Equal(t, APIModeReadOnly, c.soapClient.Mode)
}

func TestFakeSOAPClientCall(t *testing.T) {

	c := FakeSOAPClient{
		fixture: []byte(`<SOAP-ENV:Envelope>
	<SOAP-ENV:Body>
		<ns1:testResponse>
			<return>
				<item>
					<key>foo</key>
				</item>
			</return>
		</ns1:testResponse>
	</SOAP-ENV:Body>
</SOAP-ENV:Envelope>`),
	}

	var v struct {
		Item struct {
			Key string `xml:"key"`
		} `xml:"item"`
	}

	err := c.Call(SoapRequest{Method: "test"}, &v)
	require.NoError(t, err)
	assert.Equal(t, "foo", v.Item.Key)
}

func TestFakeSOAPClientFixtureFromFile(t *testing.T) {
	var err error
	c := FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/thisfiledoesnotexist")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, []byte(nil), c.fixture)

	err = c.FixtureFromFile("testdata/fakesoapclientfixturefromfile")
	require.NoError(t, err)
	assert.Equal(t, []byte("testfoobar\n"), c.fixture)
}
