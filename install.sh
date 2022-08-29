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

# check if there are binary files
shouldBuild=true
for FILE in centerm-amd64 centerm-arm centerm-i386
do
	if [ -f "$FILE" ]; then
		mv "$FILE" "$binName" && shouldBuild=false
	fi
done
if [ -f "$binName" ]; then
	shouldBuild=false
fi

# proceed with install
if $shouldBuild; then
	go build -o "$binName" . && echo "Built binary."
else
	echo "Not building, as a binary was found."
fi

mv "$binName" "$binPath" && echo "Moving binary."
cp ./completion/centerm "$completionPath" && echo "Moving bash completion file."