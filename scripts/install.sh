#!/bin/sh
BIN_DIR=${XDG_CONFIG_DATA:-$HOME/.local}/bin
cp build/cli $BIN_DIR/crypto-egg-go
sudo cp build/app /usr/bin/crypto-egg-go-server &> /dev/null # ignore busy file error?
sudo cp resources/crypto-egg.service /etc/systemd/system &> /dev/null
