# cowsay-server
cowsay-server is a Slack slash command server for cowsay

![Sample](/image/sample.png)

## TL;DR;
I just want somewhere to point my Slack app!
https://us-central1-cowsay.cloudfunctions.net/cowsay`

I want a docker image!
https://cloud.docker.com/u/guygrigsby/repository/docker/guygrigsby/cowsay



## Manual Setup

### Slack Integration

In addition to the server, you'll need to set up your slack integration explained [here](https://api.slack.com/slash-commands). I recommend doing this first so you can add your slack tokens. If you don't add any tokens to the config, the server will accept all requests.

### Server

  To run the server you need only build the Docker image and run it, or use the one here `guygrigsby/cowsay`. It can automatically get a TLS cert from Let's Encrypt via the go `autotls` package that works with the gin webserver framework as long as you have the envvar `COWSAY_TLS_DOMAIN` set to the proper domain and `COWSAY_AUTO_TLS` set to `TRUE`. If auto TLS is used, on the first connect for a image/pod the request may timeout in Slack. This is expected because the server has to obtain a cert/key pair.

### Config

#### EnvVars
 - `COWSAY_TLS_DOMAIN` The domain where the server is hosted.
 - `COWSAY_AUTO_TLS` set to `TRUE` for the server to automatically get a TLS cert for the domain in the envvar above.
 - `COWSAY_TOKENS` The Slack tokens to verify incoming requests against. If this is blank, all requests will be accepted.

