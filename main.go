package main

func main() {
	//get first two args as one string
	switch getArgsOneString(2) {

	//status section

	case "net", "n":
		cfg.netStatus()
		cfg.netWifiList()
	case "net more", "n more":
		cfg.netStatusMore()
	case "sound", "s":
		cfg.soundStatus()
	case "power", "p":
		cfg.powerStatus()

	//modification section

	case "net on", "n on":
		cfg.netSwitch(true)
	case "net off", "n off":
		cfg.netSwitch(false)
	case "net con", "n con":
		cfg.netWifiConnect(getArg(3), getArg(4))

	case "net hson", "n hson":
		cfg.netHotspotCreate(getArg(3), getArg(4))
	case "net hsoff", "n hsoff":
		cfg.netHotspotStop()

	case "sound on", "s on":
		cfg.soundSwitch("on")
	case "sound off", "s off":
		cfg.soundSwitch("off")
	case "sound toggle", "s toggle":
		cfg.soundSwitch("toggle")
	case "sound set", "s set":
		cfg.soundChange(getArg(3))
	case "sound increase", "s +":
		cfg.soundStep("pos")
	case "sound decrease", "s -":
		cfg.soundStep("neg")

	//overview section

	default:
		printDefault()
	}
}

func printDefault() {
	//print program
	println("CENTERM")
	println("\tVersion: v0.3")
	println()

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
	println("\tcenterm s|sound on|off|toggle")
	println("\tcenterm s|sound set 1-100")
	println()
	println("\tcenterm p|power")
	println()
}
