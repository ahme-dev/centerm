package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

type config struct {
	ToolNet string `json:"Network Tool"`
	ToolSound string `json:"Sound Tool"`
	ToolPower string `json:"Power Tool"`
}

var cfg = config{}

func init() {
	//make vars
	var home, _ = os.UserHomeDir()
	var configDir = home + "/.config/centerm"
	var configFile = configDir + "/toolscfg.json"

	//stat config file
	var _, err = os.Stat(configFile)

	//if config file does not exist
	if os.IsNotExist(err) {
		//set config
		if err = exec.Command("nmcli").Run(); err == nil {
			cfg.ToolNet = "nmcli"
		} else if err = exec.Command("connmanctl").Run(); err == nil {
			cfg.ToolNet = "connmanctl"
		}

		if err = exec.Command("pamixer").Run(); err == nil {
			cfg.ToolSound = "pamixer"
		} else if err = exec.Command("amixer").Run(); err == nil {
			cfg.ToolSound = "amixer"
		}

		if err = exec.Command("acpi").Run(); err == nil {
			cfg.ToolPower = "acpi"
		} else if err = exec.Command("upower").Run(); err == nil {
			cfg.ToolPower = "upower"
		}

		//write config
		jsoncfg, _ := json.Marshal(cfg)
		os.MkdirAll(configDir, os.ModePerm)
		ioutil.WriteFile(configFile, jsoncfg, 0664)
	}

	//read config
	f, _ := ioutil.ReadFile(configFile)
	json.Unmarshal(f, &cfg)
}
