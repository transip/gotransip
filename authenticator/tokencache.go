package authenticator

// TokenCache asks for two methods,
// one to save a token by name as byte array
// and one to get a previously acquired token by name returned as byte array
type TokenCache interface {
	// Set will save a token by name as byte array
	Set(key string, data []byte) error
	// Get a previously acquired token by name returned as byte array
	Get(key string) ([]byte, error)
}
