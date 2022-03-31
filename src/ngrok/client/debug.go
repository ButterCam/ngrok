// +build !release

package client

var (
	defaultRootCrtPaths = []string{}
)

const (
	defaultCrtPath = ""
	defaultKeyPath = ""
)

func useInsecureSkipVerify() bool {
	return true
}
