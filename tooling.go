package main

import "fmt"

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

// define tools to use
var netTools = []string{
	"nmcli",
	"connmanctl",
}
var soundTools = []string{
	"pamixer",
	"amixer",
}
var lightTools = []string{
	"xbacklight",
}
var powerTools = []string{
	"acpi",
	"upower",
}

func printTools() {
	fmt.Println("SELECTED TOOLS")
	fmt.Printf("\t  Net: %v\n", cfg.SelectedNetTool)
	fmt.Printf("\tSound: %v\n", cfg.SelectedSoundTool)
	fmt.Printf("\tLight: %v\n", cfg.SelectedLightTool)
	fmt.Printf("\tPower: %v\n", cfg.SelectedPowerTool)
	fmt.Println()
	fmt.Println("SUPPORTED TOOLS")
	fmt.Printf("\t  Net: %v\n", netTools)
	fmt.Printf("\tSound: %v\n", soundTools)
	fmt.Printf("\tLight: %v\n", lightTools)
	fmt.Printf("\tPower: %v\n", powerTools)
}
