package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// get arguments slice
	args := getArgs(5)

	//status section
	switch args[0] {
	case "net", "n":
		switch args[1] {
		case " ":
			cfg.netStatus()
		case "list":
			cfg.netWifiList()
		case "more":
			cfg.netStatusMore()
		case "on":
			cfg.netSwitch(true)
		case "off":
			cfg.netSwitch(false)
		case "con":
			cfg.netWifiConnect(args[2], args[3])
		case "hson":
			cfg.netHotspotCreate(args[2], args[3])
		case "hsoff":
			cfg.netHotspotStop()
		default:
			sanePrint("", fmt.Errorf(wrongInput))
		}

	case "sound", "s":
		switch args[1] {
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
		default:
			sanePrint("", fmt.Errorf(wrongInput))
		}

	case "light", "l":
		switch args[1] {
		case " ":
			cfg.lightStatus()
		case "set":
			cfg.lightChange(args[2])
		case "inc", "+":
			cfg.lightStep("pos")
		case "dec", "-":
			cfg.lightStep("neg")
		default:
			sanePrint("", fmt.Errorf(wrongInput))
		}

	case "power", "p":
		switch args[1] {
		case " ":
			cfg.powerStatus()
		case "more":
			cfg.powerStatusMore()
		default:
			sanePrint("", fmt.Errorf(wrongInput))
		}

	case "tools":
		switch args[1] {
		case " ":
			printTools()
		case "recheck":
			modifyConfig(true)
		default:
			sanePrint("", fmt.Errorf(wrongInput))
		}

	case "help":
		fallthrough

	default:
		printUsage()
	}
}

func printUsage() {
	println("USAGE")
	println("\tcenterm n|net")
	println("\tcenterm n|net list")
	println("\tcenterm n|net more")
	println("\tcenterm n|net on|off")
	println("\tcenterm n|net con SSID Password")
	println("\tcenterm n|net hson|hsoff SSID Password")
	println()
	println("\tcenterm s|sound")
	println("\tcenterm s|sound +|inc|-|dec")
	println("\tcenterm s|sound on|off|toggle|tog")
	println("\tcenterm s|sound set 1-100")
	println()
	println("\tcenterm l|light")
	println("\tcenterm l|light +|inc|-|dec")
	println("\tcenterm l|light set 1-100")
	println()
	println("\tcenterm p|power")
	println("\tcenterm p|power more")
	println()
	println("\tcenterm tools")
	println("\tcenterm tools recheck")
	println("\tcenterm help")
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
