#!/bin/bash

# automatic clone, compile and install for lazy users :)
cd /tmp
rm -rf ssh-chat # git will complain otherwise
git clone https://github.com/antonilol/ssh-chat.git
cd ssh-chat
make build
sudo cp ssh-chat /bin/ssh-chat
echo "Installed!"
