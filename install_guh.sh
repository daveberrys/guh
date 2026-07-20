#!/bin/bash

echo "This script will install the 'guh' binary in the '~/.local/bin' directory."
echo

go mod tidy && echo "Tidied go.mod and go.sum files."
go get ./ && echo "Installed all packages." || echo "Failed to install all packages."
go build -ldflags "-X 'github.com/daveberrys/guh/src/cmd/cli.Version=$(git rev-parse --short HEAD)'" -o guh . && echo "Built 'guh' binary." || echo "Failed to build 'guh' binary."
echo

cp guh /home/$USER/.local/bin/guh && echo "Successfully placed 'guh' in '~/.local/bin' directory." || echo "Failed to place 'guh' in '~/.local/bin' directory."
chmod +x /home/$USER/.local/bin/guh && echo "Successfully made 'guh' executable." || echo "Failed to make 'guh' executable."

guh >> /dev/null && echo "Successfully installed 'guh' binary." || echo "Failed to install 'guh' binary."
echo "Installation complete."