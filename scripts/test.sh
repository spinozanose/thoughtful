#!/bin/bash

# Confirm that we are running from the project root
if [ ! -f "scripts/build.sh" ]; then
  echo "Please run this script from the project root directory."
  exit 1
fi

go test ./src/...