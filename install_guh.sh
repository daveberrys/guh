#!/bin/bash

echo "This will install the compiled local 'guh' binary to '~/.local/bin' directory."
echo "Installing..."
echo

go mod tidy && echo "Tidied go.mod and go.sum files."
go get ./ && echo "Installed all packages." || echo "Failed to install all packages."
go build && echo "Built 'guh' binary." || echo "Failed to build 'guh' binary."
echo

cp guh ~/.local/bin/guh && echo "Successfully placed 'guh' in '~/.local/bin' directory." || echo "Failed to place 'guh' in '~/.local/bin' directory."
chmod +x ~/.local/bin/guh && echo "Successfully made 'guh' executable." || echo "Failed to make 'guh' executable."

guh >> /dev/null && echo "Successfully installed 'guh' binary." || echo "Failed to install 'guh' binary."
echo "Installation complete."