package client

import (
	_ "crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"ngrok/client/assets"
)

func LoadCertPool(rootCertPaths []string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	for _, certPath := range rootCertPaths {
		rootCrt, err := assets.Asset(certPath)
		if err != nil {
			return nil, err
		}

		pemBlock, _ := pem.Decode(rootCrt)
		if pemBlock == nil {
			return nil, fmt.Errorf("Bad PEM data")
		}

		certs, err := x509.ParseCertificates(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}

		pool.AddCert(certs[0])
	}

	return pool, nil
}

func LoadTLSConfig(rootCertPaths []string, crtPath string, keyPath string) (tlsConfig *tls.Config, err error) {
	var (
		certs []tls.Certificate
		pool  *x509.CertPool
	)

	if crtPath != "" && keyPath != "" {
		var (
			crt  []byte
			key  []byte
			cert tls.Certificate
		)

		if crt, err = assets.Asset(crtPath); err != nil {
			return
		}

		if key, err = assets.Asset(keyPath); err != nil {
			return
		}

		if cert, err = tls.X509KeyPair(crt, key); err != nil {
			return
		}

		certs = append(certs, cert)
	}

	if pool, err = LoadCertPool(rootCertPaths); err != nil {
		return
	}

	tlsConfig = &tls.Config{
		RootCAs:            pool,
		Certificates:       certs,
		InsecureSkipVerify: useInsecureSkipVerify(),
	}
	return
}
