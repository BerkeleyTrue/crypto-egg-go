# Crypto Egg-Go

A crypto token aggregater built in GO

## Install

Requires go binaries to build

* Run ./scripts/build-release.sh to build binaries
* Run ./scripts/install.sh to:
  * move server binary to global bin directory to be used with systemd service
  * client binary is moved to local bin directory.
* To start service run `$ systemctl start crypto-egg`
* To have service start automatically on startup run `$ systemctl enable crypto-egg`
