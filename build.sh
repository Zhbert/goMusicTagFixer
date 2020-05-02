#!/bin/bash

rm ./builds/gomusictagfixer
go build gomusictagfixer.go scanfunctions.go
cp gomusictagfixer ./builds/
rm gomusictagfixer
