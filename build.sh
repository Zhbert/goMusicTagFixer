#!/bin/bash

rm ./builds/gomusictagfixer
go build gomusictagfixer.go scanfunctions.go mp3functions.go help.go
cp gomusictagfixer ./builds/
rm gomusictagfixer
