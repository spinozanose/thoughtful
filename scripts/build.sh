#!/bin/bash

# Confirm that we are running from the project root
if [ ! -f "scripts/build.sh" ]; then
  echo "Please run this script from the project root directory."
  exit 1
fi

mkdir -p build

pushd src >/dev/null
    GOOS=darwin GOARCH=arm64 go build -o ../build/thoughtful ./...
popd >/dev/null

chmod +x build/thoughtful