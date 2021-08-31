#!/bin/bash

# automatic clone, compile and install for lazy users :)
cd /tmp
go get github.com/antonilol/ssh-chat
make build
sudo cp ssh-chat /bin/ssh-chat
