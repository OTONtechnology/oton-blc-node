#!/bin/sh

for file in *.go;
do
  cpp -dD -fpreprocessed -o output $file;
  sed -i '1d' output
  rm   $file
  mv   output   $file;
done
