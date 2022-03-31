package server

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"ngrok/server/assets"
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
		crt  []byte
		key  []byte
		cert tls.Certificate
		pool *x509.CertPool
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

	if pool, err = LoadCertPool(rootCertPaths); err != nil {
		return
	}

	tlsConfig = &tls.Config{
		RootCAs:      pool,
		ClientAuth:   clientAuth(),
		ClientCAs:    pool,
		Certificates: []tls.Certificate{cert},
	}

	return
}

func LoadHttpsConfig(crtPath string, keyPath string) (tlsConfig *tls.Config, err error) {
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

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return
}
