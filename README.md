Centerm is a terminal tool to interact with sound, power and network. It acts as an interface for a number of tools and unifies what you type as input, though does not do the same with output currently.

## Usage
Specific syntax and commands are shown within the program. Functionalities done through the tools are:
- Get network status and list wifis.
- Connect to wifis, turn networking on/off.
- Make hotspots.
- Get sound level and modify it.
- Get power info.

These tasks are performed using the supported tools:
- networkmanager
- connman (partial)
- pamixer
- amixer
- upower (partial)
- acpi

## Build
Clone the repo and ```go build .``` inside the directory.

## Install
After the build step, copy the binary and completion to their respective locations on your system.
You can also use this presumptious one-liner in the source directory (requires root privileges):
```sh
cp ./completion/centerm /usr/share/bash-completion/completions/ && cp ./centerm /usr/local/bin/
```
