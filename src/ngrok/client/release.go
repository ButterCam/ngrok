// +build release

package client

var (
	defaultRootCrtPaths = []string{"assets/client/tls/ca.pem"}
)

const (
	defaultCrtPath = "assets/client/tls/client.pem"
	defaultKeyPath = "assets/client/tls/client-key.pem"
)

func useInsecureSkipVerify() bool {
	return false
}
