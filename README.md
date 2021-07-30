# About
A terminal tool to interact with sound, power and network. CenTerm abstracts the common tasks done in the mentioned categories, making things less of a headache.

# Usage
Specific syntax and commands are shown within the program. Functionalities done through the tools are:
- Get network status and list wifis.
- Connect to wifis, turn networking on/off.
- Make hotspots.
- Get sound level and modify it.
- Get power info.

# Build
Clone the repo and ```go build .``` inside the directory.

# Install
After the build step, copy the binary and completion to their respective locations on your system.
You can also use this presumptious one-liner in the source directory (requires root privileges):
```sh
cp ./completion/centerm /usr/share/bash-completion/completions/ && cp ./centerm /usr/local/bin/
```
