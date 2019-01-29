[![Build Status](https://travis-ci.org/guygrigsby/cowsay-server.svg?branch=master)](https://travis-ci.org/guygrigsby/cowsay-server)

# cowsay-server
cowsay-server is a Slack slash command server for cowsay

![Sample](/image/sample.png)

## Setup

#### Slack Integration

In addition to the server, you'll need to set up your slack integration explained [here](https://api.slack.com/slash-commands). I recommend doing this first so you can add your slack tokens. If you don't add any tokens to the config, the server will accept all requests.

#### Server

The server is just a webserver wrapper for the `cowsay` executable. As such, you must have `cowsay` on present on the machine. If you don't have it, install it with your distro's package manager.

Fedora:
`yum install cowsay`

Debian/Ubuntu:
`apt-get install cowsay`

Mac:
`brew install cowsay`

Windows:
srsly?

After you have `cowsay` proper, you need to compille the server.

```bash
go build ./...
```
or the way I do it is to compile it on a Mac for a Debian server
```bash
env GOOS=linux GOARCH=amd64 go build
```


Next, fill out the sample config with real values. The server looks in it's working directory for a file called `config.json` by default. If you want to put it somewhere else, use the `-c` flag on server startup and provide the full path and filename, eg `/path/to/yourconfig.json`. You'll need tls certs. If you don't have any, I recommend going over to [Let's Encrypt](https://letsencrypt.org/)

Config
- `tokens` an array of valid tokens from slack. If none are provided, the server will accept all requests
- `certfile` the fullchain cert file for tls
- `keyfile` the private key correspoding to the cert file
- `listenon` the address and port the server will listen on
- `cowsayexec` the path the the cowsay executable. If absent, default is `/usr/games/cowsay`


