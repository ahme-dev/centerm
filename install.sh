#!/bin/sh

if [ "$EUID" -ne 0 ]
then
	echo "Script requires root permissions."
	exit
fi

go build -o centerm .
mv centerm /usr/local/bin/
cp ./completion/centerm /usr/share/bash-completion/completions/
