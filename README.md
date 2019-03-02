[![Build Status](https://travis-ci.org/guygrigsby/cowsay-server.svg?branch=master)](https://travis-ci.org/guygrigsby/cowsay-server)

# cowsay-server
cowsay-server is a Slack slash command server for cowsay

## TL;DR;
I just want somewhere to point my Slack app!

`https://cowsay.grigsby.dev`

![Sample](/image/sample.png)

## Setup

### Slack Integration

In addition to the server, you'll need to set up your slack integration explained [here](https://api.slack.com/slash-commands). I recommend doing this first so you can add your slack tokens. If you don't add any tokens to the config, the server will accept all requests.

### Server

  To run the server you need only build the Docker image and run it, or use the one here `guygrigsby/cowsay`. It will automatically get a TLS cert from Let's Encrypt via the go `autotls` package that works with the gin webserver framework as long as you have the envvar `COWSAY_TLS_DOMAIN` set to the proper domain.

### Config

#### EnvVars
 - `COWSAY_TLS_DOMAIN` The domain where the server is hosted. REQUIRED
 - `COWSAY_TOKENS` The Slack tokens to verify incoming requests against. If this is blank, all requests will be accepted.

