# go-ssh-agent-locker

go-ssh-agent-locker is a simple application that kills ssh-agent whenever the keychain is locked

## Installation

```
go get github.com/MDrollette/go-ssh-agent-locker
cp $GOPATH/bin/go-ssh-agent-locker /usr/local/bin/go-ssh-agent-locker
cp $GOPATH/src/github.com/MDrollette/go-ssh-agent-locker/com.drollette.matt.go-ssh-agent-locker.plist ~/Library/LaunchAgents/com.drollette.matt.go-ssh-agent-locker.plist
launchctl load ~/Library/LaunchAgents/com.drollette.matt.go-ssh-agent-locker.plist
```

## Basic Usage

ssh-agent is used to manage and securely store ssh private keys. On OSX, when a keychain is locked the ssh keys remain in ssh-agent. This application will listen for keychain lock events and shut down the ssh-agent daemon, removing any keys.
