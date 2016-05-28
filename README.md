[![Build Status](https://travis-ci.org/guygrigsby/cowsay-server.svg?branch=master)](https://travis-ci.org/guygrigsby/cowsay-server)

# cowsay-server
cowsay-server is a Slack slash command server for cowsay

![Sample](/image/sample.png)

##Setup

####Slack Integration

In addition to the server, you'll need to set up your slack integration explained [here](https://api.slack.com/slash-commands). I recommend doing this first so you can add your slack tokens. If you don't add any tokens to the config, the server will accept all requests.

####Server
All you should need to do it complile the server and fill out the sample config with real values. The server looks in it's working directory for a file called `config.json` by default. If you want to put it somewhere else, use the `-c` on server startup and provide the full path and filename, eg `/path/to/yourconfig.json`.

Config
- `tokens` an array of valid tokens from slack. If none are provided, the server will accept all requests
- `certfile` the fullchain cert file for tls
- `keyfile` the private key correspoding to the cert file
- `listenon` the address and port the server will listen on


