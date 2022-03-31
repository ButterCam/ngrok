.PHONY: default server client deps fmt clean all release-all assets client-assets server-assets contributors
export GOPATH:=$(shell pwd)

BUILDTAGS=debug
default: all

deps: assets
	go get -tags '$(BUILDTAGS)' -d -v ngrok/...

server: deps
	go install -tags '$(BUILDTAGS)' ngrok/main/ngrokd

fmt:
	go fmt ngrok/...

client: deps
	go install -tags '$(BUILDTAGS)' ngrok/main/ngrok

assets: client-assets server-assets

bin/cfssl:
	GOOS="" GOARCH="" go get github.com/cloudflare/cfssl/cmd/cfssl
	GOOS="" GOARCH="" go get github.com/cloudflare/cfssl/cmd/cfssljson

gencert: bin/cfssl
ifeq ("$(wildcard assets/tls/ca.pem)","")
	bin/cfssl gencert -initca assets/tls/ca-csr.json | bin/cfssljson -bare assets/tls/ca -
	bin/cfssl gencert -ca=assets/tls/ca.pem -ca-key=assets/tls/ca-key.pem -config=assets/tls/ca-config.json -profile=server assets/tls/server.json | bin/cfssljson -bare assets/server/tls/server
	bin/cfssl gencert -ca=assets/tls/ca.pem -ca-key=assets/tls/ca-key.pem -config=assets/tls/ca-config.json -profile=server assets/tls/client.json | bin/cfssljson -bare assets/client/tls/client
	cp assets/tls/ca.pem assets/server/tls
	cp assets/tls/ca.pem assets/client/tls
endif

bin/go-bindata:
	GOOS="" GOARCH="" go get github.com/jteeuwen/go-bindata/go-bindata

client-assets: gencert bin/go-bindata
	bin/go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=src/ngrok/client/assets/assets_$(BUILDTAGS).go \
		assets/client/...

server-assets: bin/go-bindata
	bin/go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=src/ngrok/server/assets/assets_$(BUILDTAGS).go \
		assets/server/...

release-client: BUILDTAGS=release
release-client: client

release-server: BUILDTAGS=release
release-server: server

release-all: fmt release-client release-server

all: fmt client server

clean:
	go clean -i -r ngrok/...
	rm -rf src/ngrok/client/assets/ src/ngrok/server/assets/
	rm -rf assets/tls/*.pem

contributors:
	echo "Contributors to ngrok, both large and small:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS

