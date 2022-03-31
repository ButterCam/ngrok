// +build !release

package server

import "crypto/tls"

var (
	rootCrtPaths = []string{}
)

const (
	defaultCrtPath = "assets/server/tls/server.pem"
	defaultKeyPath = "assets/server/tls/server-key.pem"
)

func clientAuth() tls.ClientAuthType {
	return tls.NoClientCert
}
