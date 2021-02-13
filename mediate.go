package main

import (
    "fmt"
)

const noTool string = "No tool available"

// print func for following methods
func sanePrint(s string, e error) {
	if s != "" {
		fmt.Println(s)
	}
	if e != nil {
		fmt.Println(e)
	}
}

//general methods handling tool selection and printing
//also future error handling

//net

func (c config) netStatus() {
	var out string
	var err error

	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliGet()
	case "connmanctl":
		out, err = connmanctlGet()
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) netStatusMore() {
	var out string
	var err error

	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliGetMore()
	case "connmanctl":
		out, err = connmanctlGetMore()
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) netSwitch(swi bool) {
	var out string
	var err error

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
		out, err = noTool, nil
	}


	sanePrint(out, err)
}

func (c config) netWifiList() {
	var out string
	var err error

	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliGetWifi()
	case "connmanctl":
		out, err = connmanctlGetWifi()
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) netWifiConnect(ssid string) {
	var out string
	var err error

	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliSetWifi(ssid)
	case "connmanctl":
		out, err = "This tool doesn't support such action yet.", nil
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) netHotspotCreate(ssid, password string) {
	var out string
	var err error

	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliMakeHotspot(ssid, password)
	case "connmanctl":
		out, err = connmanctlMakeTether(ssid, password)
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) netHotspotStop() {
	var out string
	var err error

	switch c.ToolNet {
	case "nmcli":
		out, err = nmcliStopHotspot()
	case "connmanctl":
		out, err = connmanctlStopTether()
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}


//sound

func (c config) soundStatus() {
	var out string
	var err error

	switch c.ToolSound {
	case "pamixer":
		out, err = pamixerGet()
	case "amixer":
		out, err = amixerGet()
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) soundChange(vol string) {
	var out string
	var err error

	switch c.ToolSound {
	case "pamixer":
		out, err = pamixerSet(vol)
	case "amixer":
		out, err = amixerSet(vol)
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) soundStep(pol string) {
	var out string
	var err error

	switch c.ToolSound {
	case "pamixer":
		out, err = "this feature isn't supported", nil
	case "amixer":
		switch pol {
		case "pos":
			out, err = amixerInc()
		case "neg":
			out, err = amixerDec()
		}
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

func (c config) soundSwitch(swi string) {
	var out string
	var err error

	switch c.ToolSound {
	case "pamixer":
		switch swi {
		case "on":
			out, err = pamixerSetOn()
		case "off":
			out, err = pamixerSetOff()
		case "toggle":
			out, err = "this feature isn't supported", nil
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
		out, err = noTool, nil
	}

	sanePrint(out, err)
}

//power

func (c config) powerStatus() {
	var out string
	var err error

	switch c.ToolPower {
	case "acpi":
		out, err = acpiGet()
	case "upower":
		out, err = upowerGet()
	default:
		out, err = noTool, nil
	}

	sanePrint(out, err)
}
