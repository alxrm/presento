#!/usr/bin/env bash

rm -rf ./dist/

go generate
goxc -bc="windows,darwin" -d=./dist

find ./dist/ -type d -depth -empty -exec rmdir "{}" \;

if [ ! -d "./dist/" ]; then
  echo "Could build nothing"
  exit 1
fi

echo "Done building, compressing"
cd ./dist/snapshot/

for i in */; do
 zip -r "${i%/}.zip" "$i";
done