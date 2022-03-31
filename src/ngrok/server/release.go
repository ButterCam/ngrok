// +build release

package server

import "crypto/tls"

var (
	rootCrtPaths = []string{"assets/server/tls/ca.pem"}
)

const (
	defaultCrtPath = "assets/server/tls/server.pem"
	defaultKeyPath = "assets/server/tls/server-key.pem"
)

func clientAuth() tls.ClientAuthType {
	return tls.RequireAndVerifyClientCert
}
