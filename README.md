Centerm is a terminal tool to interact with sound, power, brightness and networking. It acts as an interface for a number of tools and unifies what you type as input, though does not do the same with output.

## Install

- Clone repo and run install.sh with sudo.

- Four-liner to copy-paste into the terminal instead:

  ```sh
  git clone --depth=1 https://github.com/ahmedkabd/centerm &&
  cd centerm && chmod +x ./install.sh &&
  sudo ./install.sh &&
  cd .. && rm -rf centerm;
  ```

- You can also grab a binary from the releases page and run it.

## Usage

Specific syntax and commands are shown within the program. Functionalities done through the tools are:

- Get network status and list wifis, connect to wifis, turn networking on/off, and also make hotspots.
- Control sound, mute and modify it.
- Get brightness level and modify it.
- Get current battery level, or power info.

Tasks are done using the supported tools:

- networkmanager, connman (partial)
- pamixer, amixer
- xbacklight, brightnessctl
- acpi, upower (partial)
