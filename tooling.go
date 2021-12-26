package main

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
