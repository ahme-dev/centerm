package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

// global configuration type
type config struct {
	SelectedNetTool   string `json:"NetworkTool"`
	SelectedSoundTool string `json:"SoundTool"`
	SelectedLightTool string `json:"LightTool"`
	SelectedPowerTool string `json:"PowerTool"`
}

// global configuration struct
// sets which tools to use
var cfg = config{}

// initialize and set/get configuration
func init() {
	//make vars
	var home, _ = os.UserHomeDir()
	var configDir = home + "/.config/centerm"
	var configFile = configDir + "/toolscfg.json"

	//stat config file
	var _, err = os.Stat(configFile)

	//if config file does not exist
	if os.IsNotExist(err) {
		// define tools to use
		var netTools = []string{"nmcli", "connmanctl"}
		var soundTools = []string{"pamixer", "amixer"}
		var lightTools = []string{"xbacklight"}
		var powerTools = []string{"acpi", "upower"}

		// check for the available network tools
		for _, toolName := range netTools {
			if err = exec.Command(toolName).Run(); err == nil {
				cfg.SelectedNetTool = toolName
				break
			}
		}

		// check for the available sound tools
		for _, toolName := range soundTools {
			if err = exec.Command(toolName).Run(); err == nil {
				cfg.SelectedSoundTool = toolName
				break
			}
		}

		// check for the available brightness tools
		for _, toolName := range lightTools {
			if err = exec.Command(toolName).Run(); err == nil {
				cfg.SelectedLightTool = toolName
				break
			}
		}

		// check for the available power tools
		for _, toolName := range powerTools {
			// exception for upower which has no output without -d option
			if toolName == "upower" {
				if err = exec.Command(toolName, "-d").Run(); err == nil {
					cfg.SelectedPowerTool = toolName
					break
				}
				continue
			}

			// other tools
			if err = exec.Command(toolName).Run(); err == nil {
				cfg.SelectedPowerTool = toolName
				break
			}
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
