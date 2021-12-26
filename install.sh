#!/bin/sh

# used values
binName=centerm
binPath=/usr/local/bin
completionPath=/usr/share/bash-completion/completions

# if not root exit
if [ "$USER" != "root" ]
then
	echo "Script requires root permissions."
	exit
fi

# actions
go build -o "$binName" .
mv centerm "$binPath"
cp ./completion/centerm "$completionPath"
