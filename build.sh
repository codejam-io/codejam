#!/bin/bash


cd html
npm install
npm run build

mkdir -p ../backend/server/static_files
cp -r dist/* ../backend/server/static_files/
