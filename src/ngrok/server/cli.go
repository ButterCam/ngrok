package server

import (
	"flag"
)

type Options struct {
	httpAddr         string
	httpsAddr        string
	tunnelAddr       string
	domain           string
	tpcDomain        string
	rootCerts        stringList
	serverCrt        string
	serverKey        string
	managedHttps     bool
	httpsCrt         string
	httpsKey         string
	logto            string
	loglevel         string
	authorizedTokens stringList
}

type stringList []string

func (i *stringList) String() string {
	return "[]string"
}

func (i *stringList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func parseArgs() *Options {
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	httpAddr := flag.String("httpAddr", ":80", "Public address for HTTP connections, empty string to disable")
	httpsAddr := flag.String("httpsAddr", ":443", "Public address listening for HTTPS connections, emptry string to disable")
	tunnelAddr := flag.String("tunnelAddr", ":4443", "Public address listening for ngrok client")
	domain := flag.String("domain", "bybutter.com", "Domain where the tunnels are hosted")
	tpcDomain := flag.String("topDomain", "tcp.bybutter.com", "Domain where the tcp tunnels are hosted")
	managedHttps := flag.Bool("managed-https", false, "Make service as a internal HTTP server, all HTTPS connections will handled by external gateway")
	serverCrt := flag.String("serverCrt", "", "Path to a TLS certificate file which used for tunnel connections")
	serverKey := flag.String("serverKey", "", "Path to a TLS key file which used for tunnel connections")
	httpsCrt := flag.String("httpsCrt", "", "Path to a TLS certificate file which used for HTTPS connections, default same as serverCrt")
	httpsKey := flag.String("httpsKey", "", "Path to a TLS key file which used for HTTPS connections, default same as serverKey")
	var rootCerts stringList
	flag.Var(&rootCerts, "rootCerts", "The accepted root certs used for client auth")
	var authorizedTokens stringList
	flag.Var(&authorizedTokens, "authorizedTokens", "The accepted stringList used for clients")
	flag.Parse()

	if *serverCrt == "" {
		*serverCrt = defaultCrtPath
	}

	if *serverKey == "" {
		*serverKey = defaultKeyPath
	}

	if *httpsCrt == "" {
		*httpsCrt = *serverCrt
	}

	if *httpsKey == "" {
		*httpsKey = *serverKey
	}

	if *tpcDomain == "" {
		*tpcDomain = *domain
	}

	if len(rootCerts) == 0 {
		rootCerts = rootCrtPaths
	}

	return &Options{
		httpAddr:         *httpAddr,
		httpsAddr:        *httpsAddr,
		tunnelAddr:       *tunnelAddr,
		domain:           *domain,
		tpcDomain:        *tpcDomain,
		managedHttps:     *managedHttps,
		serverCrt:        *serverCrt,
		serverKey:        *serverKey,
		httpsCrt:         *httpsCrt,
		httpsKey:         *httpsKey,
		logto:            *logto,
		loglevel:         *loglevel,
		authorizedTokens: authorizedTokens,
		rootCerts:        rootCerts,
	}
}
