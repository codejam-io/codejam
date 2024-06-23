#!/bin/bash

# set -x  # enable debug output
set -e  # exit on error

trap ErrorHandler ERR

function ErrorHandler {
  echo -e "\e[91mError Occurred\e[0m"
  exit 1
}

echo "Building Frontend"
cd html || return
npm install
npm run build

mkdir -p ../backend/app/server/static_files
rm -rf ../backend/app/server/static_files/*
cp -r dist/* ../backend/app/server/static_files/
cp -r static/ ../backend/app/server/static_files/

echo "Building Backend"

cd ../backend/app || return
go build

if [ -e codejam.io ]; then
  if [ -e /opt/codejam/codejam.io ]; then
    rm /opt/codejam/codejam.io
  else
    echo "WARNING: Existing deployment not found"
  fi

  cp codejam.io /opt/codejam/codejam.io
else
  echo -e "\e[92mBackend Build Failed: missing binary file found\e[0m"
  exit 1
fi

echo -e "\e[32mSuccess\e[0m"

