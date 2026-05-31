#!/bin/sh

if [ -e passphrase ]; then
  rm passphrase
fi

GOOS=darwin GOARCH=amd64 go build -o passphrase_amd64 .
GOOS=darwin GOARCH=arm64 go build -o passphrase_arm64 .
lipo -create -output passphrase passphrase_amd64 passphrase_arm64
rm passphrase_amd64 passphrase_arm64

echo "built universal binary: passphrase"
