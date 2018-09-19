package gotransip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSoapResponse(t *testing.T) {
	var data []byte
	var err error
	err = parseSoapResponse([]byte("foo"), nil, 0, nil)
	assert.Error(t, err)

	// try parsing soap fault
	data = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"><SOAP-ENV:Body><SOAP-ENV:Fault><faultcode>123</faultcode><faultstring>Test Soap Fault</faultstring></SOAP-ENV:Fault></SOAP-ENV:Body></SOAP-ENV:Envelope>`)
	err = parseSoapResponse(data, nil, 0, nil)
	assert.Error(t, err)
	assert.Equal(t, "SOAP Fault 123: Test Soap Fault", err.Error())

	// try parsing empty result with HTTP 200
	data = []byte(`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://www.transip.nl/soap" xmlns:ns2="http://xml.apache.org/xml-soap" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><SOAP-ENV:Body><ns1:test><return></return></ns1:test></SOAP-ENV:Body></SOAP-ENV:Envelope>`)
	err = parseSoapResponse(data, nil, 200, nil)
	assert.NoError(t, err)

	// try parsing simple struct
	data = []byte(`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://www.transip.nl/soap" xmlns:ns2="http://xml.apache.org/xml-soap" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><SOAP-ENV:Body><ns1:test><return><key>foo</key><value>bar</value></return></ns1:test></SOAP-ENV:Body></SOAP-ENV:Envelope>`)
	var v struct {
		K string `xml:"key"`
		V string `xml:"value"`
	}

	err = parseSoapResponse(data, nil, 200, &v)
	assert.NoError(t, err)
	assert.Equal(t, "foo", v.K)
	assert.Equal(t, "bar", v.V)

	// try parsing simple struct that requires padding
	data = []byte(`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://www.transip.nl/soap" xmlns:ns2="http://xml.apache.org/xml-soap" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><SOAP-ENV:Body><ns1:test><return>foobar</return></ns1:test></SOAP-ENV:Body></SOAP-ENV:Envelope>`)
	var w struct {
		Item string `xml:"item"`
	}
	err = parseSoapResponse(data, []string{"item"}, 0, &w)
	assert.NoError(t, err)
	assert.Equal(t, "foobar", w.Item)
}

func TestSoapClientHttpRequest(t *testing.T) {
	sc := soapClient{
		Login: "test",
		Mode:  APIModeReadOnly,
		PrivateKey: []byte(`-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCZT5Eh9PmQ3flx
uFJyVG55A/RuxbYb5qv/1SBcPSZajBogtgEKvw7lcKLhkXLDSCN0pQGABRl6vTgP
aSi/s3wrKA3n9tpVa0VAQi9QGP7oVQeq3UxJ0L+yEX5HsuqYRw+mFEqxkXcdYxeV
3xGF8eB01cBOles2H5JUOMeKTyT4VQxNe+gqyG09Ia8aPDocvERBmCAdrZBSLEdH
fxGxTVzhwFh81qpR9CD8q7Q8nX7Tk0a8s42WozXu8pHQhTMHPdRkxLFxGab0p/AT
jNzG2nLi5LXS3rCSkQeHJjbPZP0T3m2OehNk40uXthH9BHgOMfyGXbX8BA8EgOeG
BCp6TNHTAgMBAAECggEANQ/4AJPEiSJ7AqQ0TQPyFIqM4IYnyLJnF64RfDth+fcB
2A6Gf8yvADSi+4WW/gYK14WA5mldb0DslVDlXKxnrpw3a/Dhkq0FE/+UVpnAKHO9
qqLbk7TflGc/mNtRHRGDVg0x6RGa853nfOTvMLgN4wJUhB6ZgWsd/26DidhoyFZG
P3Poz1u4VCsupfvxr2wo/u1vAQzAG6fVIFDTVWYtAp6nRCSg5kORIALpeNEoSMj0
JWaUoA0LiUjV6JnagMQmQtkb5ScoFMpQOoNmRdJHsJTF62lWwKSarMWHq6eEpdXG
O+Zl9CcE1/wmnY0+rVMeTJxbWn7gVXoxKG75C5jSQQKBgQDGlYvCiy6Ap/WjtrEg
Jr0aVpSSf0LkKYnvWvW7/kZJBz5KC6mEG/jW16CnA9ZnMAOTs+HKOQCYd+vPMhvz
XBNoyNxF8MrPn/lTD+3EH4sV6ZXxnLM4kpJiAEgoRSB49W4nRI6F72nUT+KakUVq
LQurujv69mxowF+W8QpnluncYQKBgQDFow6mZ6DKILqQ4425NpxZqCDXUnb71UyS
iq0dg+Fd0BsstqGqUIVeiC3I0guqMwmQ8l/o6xvXr4UVA0m26MqF5OPix6KDRfu4
Dqy1qWRFNQcj0wrFen0gAmbIEX6C6EW5wmQQ46YIym1yGOXAU7owV4pF5cZL32h3
OrhahTH6swKBgGW8YZB2S4mgAqkvxEirb//ZUV5IElXfrgnQ+Mmp+Aobyt6WYO8M
gYxXhbdqsOHGaF64LjmywEpcTZOloUoo5sys8qRmOxDpbQsPwwjR/ChqteXFGNAn
zxSj/lObLoqpehhl9/pH8FjT4Ey9lelSUINW8rmcm2eC/rXOoTz2xLKhAoGARYLG
AkzsRmsgcxk1nXDRqM7zTggZBRXOKrRPktPxjddF14IcdhR/8/GdeMY3iBMPSEWW
6grW7hMzkWJoqMZThKguZnKke9s/X0r5/6KmO5kc+8KcRTyBiaKOl8tfXZdn/p+a
Jj6LBQh9WeXb2LsZ/yqq3U6lYcYfrd+fO2chXvUCgYA0RxpJpa8TxwCPjS7u8mZm
enKsThRtC63+IWfaoGv3JkNiNbchpWcMDi7k0aAwT5aS1kbHMoGiCmnVTPmFYshS
fXxzeTpm7N5Oa2kABzA6KZot1ckfA9x5Cv/6HMToNiuKmwScx4x7ASwiLSNwuxA5
AeN9hjadhpK2ql+X9qnmkw==
-----END PRIVATE KEY-----`),
	}

	bodyFixt := `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="api.transip.nl">
	<SOAP-ENV:Body><ns1:getIpForVps><vpsName xsi:type="xsd:string">test</vpsName></ns1:getIpForVps></SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

	sr := SoapRequest{
		Service: "VpsService",
		Method:  "getIpForVps",
	}
	sr.AddArgument("vpsName", "test")

	http, err := sc.httpReqForSoapRequest(sr)
	assert.NoError(t, err)
	assert.Equal(t, "https://api.transip.nl/soap/?service=VpsService", http.URL.String())
	assert.Equal(t, "POST", http.Method)
	body, err := ioutil.ReadAll(http.Body)
	assert.NoError(t, err)
	assert.Equal(t, bodyFixt, string(body))
	assert.Equal(t, 5, len(http.Cookies()))
	for _, c := range http.Cookies() {
		switch c.Name {
		case "login":
			assert.Equal(t, "test", c.Value)
		case "mode":
			assert.Equal(t, "readonly", c.Value)
		case "timestamp":
			i, err := strconv.ParseInt(c.Value, 10, 64)
			assert.NoError(t, err)
			if i < 1536053383 {
				t.Errorf("timestamp makes no sense: %d", i)
			}
		case "nonce":
			fallthrough
		case "signature":
			assert.NotZero(t, c.Value)
		default:
			t.Errorf("unexpected cookie %s", c.Name)
		}
	}
}

func TestSoapParamsSet(t *testing.T) {
	p := soapParams{}
	// empty soapParams
	if p.Len() != 0 {
		t.Errorf("expected empty params, got length %d", p.Len())
	}

	// set first pair
	p.Set("foo", "bar")
	if p.Len() != 1 {
		t.Errorf("expected 1 entry, got length %d", p.Len())
	} else if p.keys[0] != "foo" {
		t.Errorf("expected foo, got %s", p.keys[0])
	} else if p.values[0] != "bar" {
		t.Errorf("expected bar, got %s", p.values[0])
	}

	// set second pair
	p.Set("foo2", "bar2")
	if p.Len() != 2 {
		t.Errorf("expected 2 entries, got length %d", p.Len())
	} else if p.keys[1] != "foo2" {
		t.Errorf("expected foo2, got %s", p.keys[1])
	} else if p.values[1] != "bar2" {
		t.Errorf("expected bar2, got %s", p.values[1])
	}

	// override first pair
	p.Set("foo", "bar3")
	if p.Len() != 2 {
		t.Errorf("expected 2 entries, got length %d", p.Len())
	} else if p.keys[0] != "foo" {
		t.Errorf("expected foo, got %s", p.keys[0])
	} else if p.values[0] != "bar3" {
		t.Errorf("expected bar, got %s", p.values[0])
	}
}

func TestSoapParamsSetMulti(t *testing.T) {
	p := soapParams{}

	p.SetMulti("foo", []string{
		"foo", "bar", "baz",
	})
	p.Set("bar", "baz")

	fixt := "foo[0]=foo&foo[1]=bar&foo[2]=baz&bar=baz"
	if p.Len() != 4 {
		t.Errorf("expected 4 entries, got length %d", p.Len())
	} else if result := p.Encode(); result != fixt {
		t.Errorf("encoded params do not match\nexpected: %s\nactual:   %s\n", fixt, result)
	}
}

func TestSoapParamsGet(t *testing.T) {
	p := soapParams{}
	if v, err := p.Get("foo"); err == nil {
		t.Errorf("expected error, got %s", v)
	}

	p.Set("foo", "bar")
	if v, err := p.Get("foo"); err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	} else if v != "bar" {
		t.Errorf("expected bar string, got %s", v)
	}

	p.Set("foo", "bar2")
	if v, err := p.Get("foo"); err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	} else if v != "bar2" {
		t.Errorf("expected bar2 string, got %s", v)
	}
}

func TestSoapParamsEncode(t *testing.T) {
	p := soapParams{}
	p.Set("foo", "bar+bar")
	p.Set("bar", "bar bar")
	p.Set("baz", "YmFyCg==")

	fixt := "foo=bar%2Bbar&bar=bar+bar&baz=YmFyCg%3D%3D"
	if enc := p.Encode(); enc != fixt {
		t.Errorf("encoded params do not match\nexpected: %s\nactual:   %s\n", fixt, enc)
	}
}

func TestGetSOAPArgs(t *testing.T) {
	fixture := []byte("<ns1:getStuff>foo</ns1:getStuff>")
	if o := getSOAPArgs("getStuff", "foo"); !bytes.Equal(o, fixture) {
		t.Errorf("\nexpected: %s\nactual:   %s\n", fixture, o)
	}

	fixture = []byte("<ns1:getStuff>foobar</ns1:getStuff>")
	if o := getSOAPArgs("getStuff", "foo", "bar"); !bytes.Equal(o, fixture) {
		t.Errorf("\nexpected: %s\nactual:   %s\n", fixture, o)
	}

	if o := getSOAPArgs("getStuff", []string{"foo", "bar"}...); !bytes.Equal(o, fixture) {
		t.Errorf("\nexpected: %s\nactual:   %s\n", fixture, o)
	}
}

func TestGetSOAPArg(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		input    interface{}
		fixture  string
	}{
		{"string", "fooBar", "barFoo", `<fooBar xsi:type="xsd:string">barFoo</fooBar>`},
		{"int", "BooFar", int(1), `<BooFar xsi:type="xsd:integer">1</BooFar>`},
		{"int32", "BarFoo", int32(1), `<BarFoo xsi:type="xsd:integer">1</BarFoo>`},
		{"int64", "BaoFor", int64(1), `<BaoFor xsi:type="xsd:integer">1</BaoFor>`},
		{"array", "barFoo", []string{"bar", "Foo"}, `<barFoo SOAP-ENC:arrayType="xsd:string[2]" xsi:type="ns1:ArrayOfString"><item xsi:type="xsd:string">bar</item><item xsi:type="xsd:string">Foo</item></barFoo>`},
	}

	for _, test := range tests {
		if output := getSOAPArg(test.name, test.input); output != test.fixture {
			t.Errorf("getSOAPArg test '%s' failed\nexpected: %s\nactual:   %s\n", test.testName, test.fixture, output)
		}
	}
}

func TestPadXMLData(t *testing.T) {
	// this is our initial XML
	data := []byte(`<fooBar xsi:type="xsd:string">barFoo</fooBar>`)
	// we would like to pad it with these elements
	padding := []string{
		"foo",
		"<foo>",
		`<foo bar="baz">`,
	}

	output := string(padXMLData(data, padding))
	fixture := `<foo><foo><foo bar="baz"><fooBar xsi:type="xsd:string">barFoo</fooBar></foo></foo></foo>`
	if output != fixture {
		t.Errorf("padding not same\nexpected: %s\nactual:   %s\n", fixture, output)
	}
}

func TestSoapRequestAddArgument(t *testing.T) {
	sr := SoapRequest{
		Service: "VpsService",
		Method:  "handoverVps",
	}

	sr.AddArgument("vpsName", "test")
	sr.AddArgument("handoverVps", "test2")

	assert.Equal(t, "0=test&1=test2", sr.params.Encode())
	assert.Equal(t, []string{"<vpsName xsi:type=\"xsd:string\">test</vpsName>", "<handoverVps xsi:type=\"xsd:string\">test2</handoverVps>"}, sr.args)

	sr = SoapRequest{
		Service: "HaipService",
		Method:  "setHaipVpses",
	}

	sr.AddArgument("haipName", "test")
	sr.AddArgument("vpsNames", []string{"test-vps", "test-vps2"})
	sr.AddArgument("ipAddress", net.ParseIP("1.2.3.4"))
	sr.AddArgument("sourcePort", 86400)

	assert.Equal(t, "0=test&1[0]=test-vps&1[1]=test-vps2&3=1.2.3.4&4=86400", sr.params.Encode())
	assert.Equal(t, []string{
		"<haipName xsi:type=\"xsd:string\">test</haipName>",
		"<vpsNames SOAP-ENC:arrayType=\"xsd:string[2]\" xsi:type=\"ns1:ArrayOfString\"><item xsi:type=\"xsd:string\">test-vps</item><item xsi:type=\"xsd:string\">test-vps2</item></vpsNames>",
		"<ipAddress xsi:type=\"xsd:string\">1.2.3.4</ipAddress>",
		"<sourcePort xsi:type=\"xsd:string\">86400</sourcePort>",
	}, sr.args)
}

type TestParamsEncoder struct {
	key   string
	value string
}

func (t TestParamsEncoder) EncodeParams(prm ParamsContainer) {
	prm.Set(fmt.Sprintf("%d[key]", prm.Len()), t.key)
	prm.Set(fmt.Sprintf("%d[value]", prm.Len()), t.value)
}

func (t TestParamsEncoder) EncodeArgs(key string) string {
	return fmt.Sprintf("<%s><item><key>%s</key><value>%s</value></item></%s>",
		key, t.key, t.value, key)
}

func TestSoapRequestAddArgumentParamsEncoder(t *testing.T) {
	sr := SoapRequest{
		Service: "VpsService",
		Method:  "handoverVps",
	}

	enc := TestParamsEncoder{
		key:   "foo",
		value: "bar",
	}
	sr.AddArgument("encoder", enc)

	fixtEnv := `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="api.transip.nl">
	<SOAP-ENV:Body><ns1:handoverVps><encoder><item><key>foo</key><value>bar</value></item></encoder></ns1:handoverVps></SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
	assert.Equal(t, fixtEnv, sr.getEnvelope())
	assert.Equal(t, "0[key]=foo&1[value]=bar", sr.params.Encode())
}

func TestSoapRequestGetEnvelope(t *testing.T) {
	sr := SoapRequest{
		Service: "VpsService",
		Method:  "handoverVps",
	}

	sr.AddArgument("vpsName", "test")
	sr.AddArgument("handoverVps", "test2")

	fixture := `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="api.transip.nl">
	<SOAP-ENV:Body><ns1:handoverVps><vpsName xsi:type="xsd:string">test</vpsName><handoverVps xsi:type="xsd:string">test2</handoverVps></ns1:handoverVps></SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
	assert.Equal(t, fixture, sr.getEnvelope())
}

func TestTestParamsContainer(t *testing.T) {
	prm := TestParamsContainer{}

	prm.Set("foo", "bar")
	prm.SetMulti("bar", []string{"boo", "far"})

	assert.Equal(t, "0foo=bar&8bar[0]=boo&20bar[1]=far", prm.Prm)
}
