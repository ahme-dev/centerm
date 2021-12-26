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
	fmt.Println("SELECTED")
	fmt.Printf("\t  Net tool: (%v)\n", cfg.SelectedNetTool)
	fmt.Printf("\tSound tool: (%v)\n", cfg.SelectedSoundTool)
	fmt.Printf("\tLight tool: (%v)\n", cfg.SelectedLightTool)
	fmt.Printf("\tPower tool: (%v)\n", cfg.SelectedPowerTool)
	fmt.Println()
	fmt.Println("SUPPORTED")
	fmt.Printf("\t  Net tools: %v\n", netTools)
	fmt.Printf("\tSound tools: %v\n", soundTools)
	fmt.Printf("\tLight tools: %v\n", lightTools)
	fmt.Printf("\tPower tools: %v\n", powerTools)
}
