# ngrok - Introspected tunnels to localhost ([homepage](https://ngrok.com))
### ”I want to expose a local server behind a NAT or firewall to the internet.”
![](https://ngrok.com/static/img/overview.png)

## What is ngrok?
ngrok is a reverse proxy that creates a secure tunnel from a public endpoint to a locally running web service.
ngrok captures and analyzes all traffic over the tunnel for later inspection and replay.

## Based on ngrok 1.x

This fork is based on ngrok 1.x which is opensource and free.
You must be noticed that ngrok 1.x is no longer supported by ngrok officially.  
We create this fork for our development env, **you should not use this version for PRODUCTION.**  
ngrok 2.x is the successor to 1.x. But it no longer opensource and free with self-hosting.

## Changes from ngrok 1.x main branch

We add some features for more easy and safety self-hosting ngrokd server.
1. Mutual authentication for ngrok server and client.
2. Managed https tunnel support.
3. Easy config and build for your own ngrok server and client.
4. Simple token by using `-authorizedTokens` opt with ngrokd.

See [ngrok self-hosting guide](docs/SELFHOSTING.md) for more information.

## Production Use

**DO NOT RUN THIS VERSION OF NGROK (1.X) IN PRODUCTION**.   
Both the client and server are known to have serious reliability issues including memory and file descriptor leaks as well as crashes. There is also no HA story as the server is a SPOF. You are advised to run 2.0 for any production quality system. 

## What can I do with ngrok?
- Expose any http service behind a NAT or firewall to the internet on a subdomain of ngrok.com
- Expose any tcp service behind a NAT or firewall to the internet on a random port of ngrok.com
- Inspect all http requests/responses that are transmitted over the tunnel
- Replay any request that was transmitted over the tunnel

## What is ngrok useful for?
- Temporarily sharing a website that is only running on your development machine
- Demoing an app at a hackathon without deploying
- Developing any services which consume webhooks (HTTP callbacks) by allowing you to replay those requests
- Debugging and understanding any web service by inspecting the HTTP traffic
- Running networked services on machines that are firewalled off from the internet

## Developing on ngrok
[ngrok developer's guide](docs/DEVELOPMENT.md)
