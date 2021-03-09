package main

import (
	"os"
	"strings"
)

func main() {
	// get arguments slice
	args := getArgs(5)

	//status section
	switch (args[0]) {
	case "net", "n":
		switch (args[1]) {
		case " ":
			cfg.netStatus()
			cfg.netWifiList()
		case "more":
			cfg.netStatusMore()
		case "on":
			cfg.netSwitch(true)
		case "off":
			cfg.netSwitch(false)
		case "con":
			cfg.netWifiConnect(args[2], args[3])
		case "hson" :
			cfg.netHotspotCreate(args[2], args[3])
		case "hsoff" :
			cfg.netHotspotStop()
		}
	case "sound", "s":
		switch (args[1]) {
		case " ":
			cfg.soundStatus()
		case "on":
			cfg.soundSwitch("on")
		case "off":
			cfg.soundSwitch("off")
		case "toggle", "tog":
			cfg.soundSwitch("toggle")
		case "set":
			cfg.soundChange(args[2])
		case "inc", "+":
			cfg.soundStep("pos")
		case "dec", "-":
			cfg.soundStep("neg")
		}
	case "power", "p":
		cfg.powerStatus()

	default:
		//print config
		println("CONFIG")
		println("\tNet tool:", cfg.ToolNet)
		println("\tSound tool:", cfg.ToolSound)
		println("\tPower tool:", cfg.ToolPower)
		println()

		//print help
		println("USAGE")
		println("\tcenterm n|net")
		println("\tcenterm n|net more")
		println("\tcenterm n|net on|off")
		println("\tcenterm n|net con SSID Password")
		println("\tcenterm n|net hson|hsoff SSID Password")
		println()
		println("\tcenterm s|sound")
		println("\tcenterm s|sound +|increase")
		println("\tcenterm s|sound -|decrease")
		println("\tcenterm s|sound on|off|toggle|tog")
		println("\tcenterm s|sound set 1-100")
		println()
		println("\tcenterm p|power")
	}
}

//returns the given input as a slice with specified length
func getArgs(selLen uint8) []string {
	var str string
	var strSlice []string
	var argsNum = uint8(len(os.Args)) - 1

	// get the string
	switch selLen {
	case 0:
		str = ""
	case 1:
		str = os.Args[1]
	default:
		var index uint8 = 1
		for index <= selLen {
			if index <= argsNum {
				if str == "" {
					str = os.Args[index]
				} else {
					str = str + "." + os.Args[index]
				}
			} else {
				str = str + "." + " "
			}
			index = index + 1
		}
	}

	// cut the string
	strSlice = strings.Split(str, ".")

	return strSlice
}
