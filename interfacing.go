package main

import (
	"fmt"
	"os"
)

//
// 		Error codes
//
// 0	Program ran successfully even if there was no output
//
// 1	Error from limitations of centerm
//		Whether there is no tool to perform action
//		or the tool can't perform the action
//		or there is wrong input
//
// 2	Error on tool's end
//		or due to miscommunication between centerm and tool
//

// centerm errors
const (
	noTool     string = "no tool available"
	wrongInput string = "wrong input was given"
	noSupport  string = "tool doesn't support this action"
)

// print func for following methods
// works as the output pipe for the whole program
func sanePrint(s string, e error) {
	if e != nil {
		// check if the error is of centerm limitation
		switch e.Error() {
		case noTool:
			fallthrough
		case wrongInput:
			fallthrough
		case noSupport:
			fmt.Println("[INFO]", e)
			os.Exit(1)
		}

		// otherwise if error is from program
		fmt.Println("[ERR]", e)
		os.Exit(2)
	} else if s == "" {
		// when there is no output
		fmt.Println("[INF] no output")
		os.Exit(0)
	} else {
		// when output is returned
		fmt.Println(s)
		os.Exit(0)
	}
}

// general methods handling tool selection and printing
// also future error handling

// variables to use rather than be made inside each function
var (
	out string
	err error
)

// net

func (c config) netStatus() {
	switch c.SelectedNetTool {
	case "nmcli":
		out, err = nmcliGet()
	case "connmanctl":
		out, err = connmanctlGet()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netStatusMore() {
	switch c.SelectedNetTool {
	case "nmcli":
		out, err = nmcliGetMore()
	case "connmanctl":
		out, err = connmanctlGetMore()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netSwitch(switchOn bool) {
	switch c.SelectedNetTool {
	case "nmcli":
		if switchOn {
			out, err = nmcliSetOn()
		} else {
			out, err = nmcliSetOff()
		}
	case "connmanctl":
		if switchOn {
			out, err = connmanctlSetOn()
		} else {
			out, err = connmanctlSetOff()
		}
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netWifiList() {
	switch c.SelectedNetTool {
	case "nmcli":
		out, err = nmcliGetWifi()
	case "connmanctl":
		out, err = connmanctlGetWifi()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netWifiConnect(ssid, password string) {
	switch c.SelectedNetTool {
	case "nmcli":
		out, err = nmcliSetWifi(ssid, password)
	case "connmanctl":
		out, err = "", fmt.Errorf(noSupport)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netHotspotCreate(ssid, password string) {
	switch c.SelectedNetTool {
	case "nmcli":
		out, err = nmcliMakeHotspot(ssid, password)
	case "connmanctl":
		out, err = connmanctlMakeTether(ssid, password)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netHotspotStop() {
	switch c.SelectedNetTool {
	case "nmcli":
		out, err = nmcliStopHotspot()
	case "connmanctl":
		out, err = connmanctlStopTether()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

// sound

func (c config) soundStatus() {
	switch c.SelectedSoundTool {
	case "pamixer":
		out, err = pamixerGet()
	case "amixer":
		out, err = amixerGet()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) soundChange(vol string) {
	switch c.SelectedSoundTool {
	case "pamixer":
		out, err = pamixerSet(vol)
	case "amixer":
		out, err = amixerSet(vol)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) soundStep(pol string) {
	switch c.SelectedSoundTool {
	case "pamixer":
		out, err = "", fmt.Errorf(noSupport)
	case "amixer":
		switch pol {
		case "pos":
			out, err = amixerInc()
		case "neg":
			out, err = amixerDec()
		}
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) soundSwitch(swi string) {
	switch c.SelectedSoundTool {
	case "pamixer":
		switch swi {
		case "on":
			out, err = pamixerSetOn()
		case "off":
			out, err = pamixerSetOff()
		case "toggle":
			out, err = "", fmt.Errorf(noSupport)
		}
	case "amixer":
		switch swi {
		case "on":
			out, err = amixerSetOn()
		case "off":
			out, err = amixerSetOff()
		case "toggle":
			out, err = amixerSetOnOff()
		}
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

// light (brightness)

func (c config) lightStatus() {
	switch c.SelectedLightTool {
	case "xbacklight":
		out, err = xbacklightGet()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) lightChange(level string) {
	switch c.SelectedLightTool {
	case "xbacklight":
		out, err = xbacklightSet(level)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) lightStep(pol string) {
	switch c.SelectedLightTool {
	case "xbacklight":
		switch pol {
		case "pos":
			out, err = xbacklightInc()
		case "neg":
			out, err = xbacklightDec()
		}
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

// power

func (c config) powerStatus() {
	switch c.SelectedPowerTool {
	case "acpi":
		out, err = acpiGet()
	case "upower":
		out, err = upowerGet()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) powerStatusMore() {
	switch c.SelectedPowerTool {
	case "acpi":
		out, err = acpiGetMore()
	case "upower":
		out, err = "", fmt.Errorf(noSupport)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}
