package main

import "fmt"

// global configuration type
type Config struct {
	SelectedNetTool   string `json:"NetworkTool"`
	SelectedSoundTool string `json:"SoundTool"`
	SelectedLightTool string `json:"LightTool"`
	SelectedPowerTool string `json:"PowerTool"`
}

// global configuration struct
// sets which tools to use
var cfg = Config{}

// define Tools to use
type Tools struct {
	net   []string
	sound []string
	light []string
	power []string
}

var tools = Tools{
	net:   []string{"nmcli", "connmanctl"},
	sound: []string{"amixer", "pamixer"},
	light: []string{"xbacklight"},
	power: []string{"acpi", "upower"},
}

func printTools() {
	// make network string
	var netString = ""
	for _, val := range tools.net {
		if val == cfg.SelectedNetTool {
			netString += " -" + val + "- "
			continue
		}
		netString += " " + val + " "
	}

	// make sound string
	var soundString = ""
	for _, val := range tools.sound {
		if val == cfg.SelectedSoundTool {
			soundString += " -" + val + "- "
			continue
		}
		soundString += " " + val + " "
	}

	// make light string
	var lightString = ""
	for _, val := range tools.light {
		if val == cfg.SelectedLightTool {
			lightString += " -" + val + "- "
			continue
		}
		lightString += " " + val + " "
	}

	// make power string
	var powerString = ""
	for _, val := range tools.power {
		if val == cfg.SelectedPowerTool {
			powerString += " -" + val + "- "
			continue
		}
		powerString += " " + val + " "
	}

	// print out the tools' strings
	fmt.Println("TOOLS \t -selected-")
	fmt.Println()
	fmt.Println("  Net: " + "[" + netString + "]")
	fmt.Println("Sound: " + "[" + soundString + "]")
	fmt.Println("Light: " + "[" + lightString + "]")
	fmt.Println("Power: " + "[" + powerString + "]")
}
