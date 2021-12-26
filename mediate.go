package main

import (
	"fmt"
	"os"
)

const (
	noTool    string = "[INF] no tool available"
	noSupport string = "[INF] tool doesn't support this action"
)

var (
	out string
	err error
)

// print func for following methods
func sanePrint(s string, e error) {
	if e != nil {
		fmt.Println("[ERR]", e)
		os.Exit(2)
	} else if s == "" {
		fmt.Println("[INF] no output")
		os.Exit(1)
	} else {
		fmt.Println(s)
		os.Exit(0)
	}
}

//general methods handling tool selection and printing
//also future error handling

//net

func (c config) netStatus() {
	switch c.ToolNet {
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
	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliGetMore()
	case "connmanctl":
		out, err = connmanctlGetMore()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netSwitch(swi bool) {
	switch c.ToolNet {
	case "nmcli":
		if swi == true {
			out, err = nmcliSetOn()
		}
		if swi == false {
			out, err = nmcliSetOff()
		}
	case "connmanctl":
		if swi == true {
			out, err = connmanctlSetOn()
		}
		if swi == false {
			out, err = connmanctlSetOff()
		}
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) netWifiList() {
	switch c.ToolNet {
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
	switch c.ToolNet {
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
	switch c.ToolNet {
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
	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliStopHotspot()
	case "connmanctl":
		out, err = connmanctlStopTether()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

//sound

func (c config) soundStatus() {
	switch c.ToolSound {
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
	switch c.ToolSound {
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
	switch c.ToolSound {
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
	switch c.ToolSound {
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

//light (brightness)

func (c config) lightStatus() {
	switch c.ToolLight {
	case "xbacklight":
		out, err = xbacklightGet()
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) lightChange(level string) {
	switch c.ToolLight {
	case "xbacklight":
		out, err = xbacklightSet(level)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}

func (c config) lightStep(pol string) {
	switch c.ToolLight {
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

//power

func (c config) powerStatus() {
	switch c.ToolPower {
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
	switch c.ToolPower {
	case "acpi":
		out, err = acpiGetMore()
	case "upower":
		out, err = "", fmt.Errorf(noSupport)
	default:
		out, err = "", fmt.Errorf(noTool)
	}

	sanePrint(out, err)
}
