#!/bin/bash

# Create the directory if it doesn't exist.
mkdir -p assets/images/thumbnails

for id in \
  "par-3aY6" \
  "bru-qC83" \
  "bru-C27w" \
  "lon-iD95" \
  "flr-t3H0" \
  "flr-h4N0" \
  "flr-fM10" \
  "flr-n71Z" \
  "lon-Vx10" \
  "wnd-a9W2" \
  "wnd-rL20" \
  "wnd-i51N" \
  "lon-44An" \
  "lon-eB17" \
  "lon-0Lz1" \
  "los-2uC8" \
  "los-0x1C" \
  "los-2nC5" \
  "lon-P94z" \
  "lon-9Yn2" \
  "lon-74Nn" \
  "lon-1S5q" \
  "lon-wI31"
do
  filepath="assets/images/thumbnails/$id.jpg"
  if [ ! -f $filepath ]; then
    wget "https://andreimuntean.com/thumbnail/$id" -O $filepath
  fi
done

echo "All thumbnails successfully downloaded."

