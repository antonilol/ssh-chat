#!/bin/bash

# automatic clone, compile and install for lazy users :)
# with a simple command:
#
#  curl https://raw.githubusercontent.com/antonilol/ssh-chat/master/install.sh | bash
#


# clone to ram, much faster and you don't have to keep it
cd /tmp

# remove ssh-chat dir (if it exists) to stop git from complaining
rm -rf ssh-chat

# clone
git clone https://github.com/antonilol/ssh-chat.git

# go in
cd ssh-chat

# compile
make build &&

# copy built binary to /usr/bin/
sudo cp ssh-chat /bin/ssh-chat &&

# tell the user it's installed
echo "Installed!"
