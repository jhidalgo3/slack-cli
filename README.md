

It is adaptation of [https://github.com/takecy/slack-cli](https://github.com/takecy/slack-cli). Thanks @takecy for the original code.

![](https://img.shields.io/badge/golang-1.10.0-blue.svg?style=flat)
[![GoDoc](https://godoc.org/github.com/jhidalgo3/slack-cli?status.svg)](https://godoc.org/github.com/jhidalgo3/slack-cli)

# slack-cli

Simple command-line client for slack by golang

## Features
* Post message only
* Use Incoming Web hook service on slack

## Install
### via Go
```shell
$ go get github.com/jhidalgo3/slack-cli
```

### via Binary
[Download](https://github.com/jhidalgo3/slack-cli/releases) and copy to your `$PATH`.

## Usage
```shell
$ slack-cli config <incoming web hook URL>
$ slack-cli post <message>
```
post to specific channel (not prefix [#])
```shell
$ slack-cli post -c general <message>
```

## Remove
### via Go
```shell
$ rm $GOPATH/bin/slack-cli
$ rm -r $HOME/.slack_cli
```
### via Binary
Remove your slac-cli binary in `$PATH`.  
and
```shell
$ rm -r $HOME/.slack_cli
```

## License
MIT