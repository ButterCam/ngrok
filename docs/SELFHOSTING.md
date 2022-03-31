# How to run your own ngrokd server

Running your own ngrok server is really easy! The instructions below will guide you along your way!

## 1. Change server const for your company

There are some const should be changed for easily use. Change the const in [const.go](../src/ngrok/selfhosting/const.go)
in selfhosting package.

```go
package selfhosting

const (
	Domain       = "bybutter.com"     // <- change to your company domain
	TcpDomain    = "tcp." + Domain    // domain for tcp tunnels
	NgrokdDomain = "ngrokd." + Domain // domain for ngrokd server
	NgrokdPort   = ":4443"
)
```

Change the SSL certificate config in [ca-config.json](../assets/tls/ca-config.json)

```json5
{
  // change to your company domain
  "CN": "bybutter.com",
  "hosts": [
    // also this
    "bybutter.com",
    // and also this
    "*.bybutter.com"
  ]
}
```

Change the SSL certificate config in [server.json](../assets/tls/server.json)

```json5
{
  "CN": "Server",
  "hosts": [
    // change to your ngrokd domain
    "ngrokd.bybutter.com",
    // add this line if you need run https tunnel with self-signed
    "*.bybutter.com"
  ],
}
```

## 2. Generate self-signed SSL certificates for auth

ngrok use self-signed cert for mutual authentication, you should generate a new cert for it.
If your company has owned general cert for https, don't use it for this sense.

```shell
make gencert # run this command to generate all cert we needed
             # root cert will store in assets/tls
             # server cert will store in assets/server/tls
             # client cert will store in assets/client/tls
```

## 3. Modify your DNS

You need to use the DNS management tools given to you by your provider to create an A record which points *.example.com
to the IP address of the server where you will run ngrokd.

## 4. Compile it

You can compile an ngrokd server with the following command:

```shell
make release-server # this command will auto gen certs
```

Make sure you compile it with the GOOS/GOARCH environment variables set to the platform of your target server. Then copy
the binary over to your server.

## 5. Run the server

You'll run the server with ont of following commands.

**Use self-signed cert for https connections.**
```shell
./bin/ngrokd
```

**Use general cert for https connections.**
```shell
./bin/ngrokd -httpsKey="/path/to/tls.key" -httpsCrt="/path/to/tls.crt"
```

**Use managed https connections.**
Managed https means ngrokd will not listen the https port, you should use a gateway for proxying the https requests to http port like cloudflare, nginx or something.
```shell
./bin/ngrokd -managed-https
```

## 6. Compile the client

You can compile auto authed clients with the following command:

```shell
make release-client # this command will auto gen certs
GOOS=windows GOARCH=amd64 make release-client # compile windows client
GOOS=darwin GOARCH=amd64 make release-client # compile macos client
GOOS=darwin GOARCH=arm64 make release-client # compile macos client for M1
GOOS=linux  GOARCH=amd64 make release-client # compile linux client
# all golang GOOS/GOARCH also be supported
```

## 7. Connect with a client

Then, just run ngrok as usual to connect securely to your own ngrokd server!

```shell
ngrok 80
```
