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
Having Git and Golang installed, you can use this one-liner:
```sh
git clone --depth=1 https://gitlab.com/hvf/centerm.git && cd centerm && go build -o ../centerm-bin && cd .. && rm -rf centerm
```
