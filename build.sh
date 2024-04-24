#!/bin/bash


cd html || return
npm install
npm run build

mkdir -p ../backend/app/server/static_files
cp -r dist/* ../backend/app/server/static_files/
