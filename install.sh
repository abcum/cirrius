#!/bin/bash

# bash <(curl -s https://pkg.cirrius.io/install.sh)

mkdir -p $HOME/.cirrius/bin

curl -o $HOME/.cirrius/bin/cirrius https://pkg.cirrius.io/osx && chmod +x $HOME/.cirrius/bin/cirrius

echo 'export PATH=$PATH:$HOME/.cirrius/bin' >> ~/.bash_profile
