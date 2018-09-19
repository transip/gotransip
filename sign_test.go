package gotransip

import (
	"testing"
)

func TestSignWithKey(t *testing.T) {
	key := `-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`

	params := &soapParams{}
	params.Set("__method", "getHaip")
	params.Set("__service", "HaipService")
	params.Set("__hostname", "api.transip.nl")
	params.Set("__timestamp", "1534839460")
	params.Set("__nonce", "5b7bcab97f1a98.77032926")

	signature, err := signWithKey(params, []byte(key))
	if err != nil {
		t.Error(err.Error())
	}

	fixture := "VBALqck%2BScVciJb%2BN2uBaPT1GaW7uLvKDsC9GfLPZT5sAIh2jhQ9Dc5mDIW9czrxxzOY3Vl1AWQI%2FMDbcIHAT4s8umpBYs8ZH4ORqiMZn4FOcypKRdPZOIdeHsqF%2FMv0Yb5YwBR6lBJrXAdh8DM%2BWt%2Fi8ZfPZV8KtXOyFb1zna9xEVmco6TSDL%2BpjUHDvzRkgJsYLZeZEOrvM7qDOxCVWRz1BKuf7wgsam6VGA6QFMA1mcjWdRd89X55075WwNXm4tEOGOMq%2BN5cf6N%2BERQMVPUF9w3EIv50bJCayWJduuk73sHUJn80tquJ6eVjky%2FS%2FDG1hUhyvngGUhwFaAWcJw%3D%3D"
	if signature != fixture {
		t.Errorf("signature does not match\nexpected: %s\nactual:   %s\n", fixture, signature)
	}
}
