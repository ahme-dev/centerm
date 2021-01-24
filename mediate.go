package main

import (
    "fmt"
)

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
	}

	sanePrint(out, err)
}

func (c config) soundSwitch(sw bool) {
	var out string
	var err error

	switch c.ToolSound {
	case "pamixer":
		if sw == true {
			out, err = pamixerSetOn()
		}
		if sw == false {
			out, err = pamixerSetOff()
		}
	case "amixer":
		if sw == true {
			out, err = amixerSetOn()
		}
		if sw == false {
			out, err = amixerSetOff()
		}
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
	}

	sanePrint(out, err)
}
