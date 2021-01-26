package main

import (
	"os/exec"
)

//networkmanager
func nmcliGet() (output string, err error) {
	o, e := exec.Command("nmcli", "general", "status").CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliGetMore() (output string, err error) {
	o, e := exec.Command("nmcli", "device", "show").CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliSetOn() (output string, err error) {
	o, e := exec.Command("nmcli", "networking", "on").CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliSetOff() (output string, err error) {
	o, e := exec.Command("nmcli", "networking", "off").CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliGetWifi() (output string, err error) {
	exec.Command("nmcli", "device", "wifi", "rescan").Run()
	o, e := exec.Command("nmcli", "device", "wifi", "list").CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliSetWifi(n string) (output string, err error) {
	o, e := exec.Command("nmcli", "device", "wifi", "connect", n).CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliMakeHotspot(ssid, password string) (output string, err error) {
	o, e := exec.Command("nmcli", "device", "wifi", "hotspot", "ssid", ssid, "password", password).CombinedOutput()
	output, err = string(o), e
	return
}

func nmcliStopHotspot() (output string, err error) {
	o, e := exec.Command("nmcli", "device", "disconnect", "wlan0").CombinedOutput()
	output, err = string(o), e
	return
}

//connman
func connmanctlGet() (output string, err error) {
	o, e := exec.Command("connmanctl", "state").CombinedOutput()
	output, err = string(o), e
	return
}

func connmanctlGetMore() (output string, err error) {
	o, e := exec.Command("connmanctl", "technologies").CombinedOutput()
	output, err = string(o), e
	return
}

func connmanctlSetOn() (output string, err error) {
	o, e := exec.Command("connmanctl", "disable", "offline").CombinedOutput()
	output, err = string(o), e
	return
}

func connmanctlSetOff() (output string, err error) {
	o, e := exec.Command("connmanctl", "enable", "offline").CombinedOutput()
	output, err = string(o), e
	return
}

func connmanctlGetWifi() (output string, err error) {
	exec.Command("connmanctl", "scan", "wifi").Run()
	o, e := exec.Command("connmanctl", "services").CombinedOutput()
	output, err = string(o), e
	return
}

func connmanctlMakeTether(ssid, password string) (output string, err error) {
	o, e := exec.Command("connmanctl", "tether", "wifi", "on", ssid, password).CombinedOutput()
	output, err = string(o), e
	return
}

func connmanctlStopTether() (output string, err error) {
	o, e := exec.Command("connmanctl", "tether", "wifi", "off").CombinedOutput()
	output, err = string(o), e
	return
}

//pamixer
func pamixerGet() (output string, err error) {
	o, e := exec.Command("pamixer", "--get-volume-human").CombinedOutput()
	output, err = string(o), e
	return
}

func pamixerSet(n string) (output string, err error) {
	o, e := exec.Command("pamixer", "--set-volume", n).CombinedOutput()
	output, err = string(o), e
	return
}

func pamixerSetOn() (output string, err error) {
	o, e := exec.Command("pamixer", "-u").CombinedOutput()
	output, err = string(o), e
	return
}

func pamixerSetOff() (output string, err error) {
	o, e := exec.Command("pamixer", "-m").CombinedOutput()
	output, err = string(o), e
	return
}

//alsa-utils
func amixerGet() (output string, err error) {
	o, e := exec.Command("amixer", "sget", "Master").CombinedOutput()
	output, err = string(o), e
	return
}

func amixerSet(n string) (output string, err error) {
	o, e := exec.Command("amixer", "sset", "Master", n+"%").CombinedOutput()
	output, err = string(o), e
	return
}

func amixerInc() (output string, err error) {
	o, e := exec.Command("amixer", "sset", "Master", "5+").CombinedOutput()
	output, err = string(o), e
	return
}

func amixerDec() (output string, err error) {
	o, e := exec.Command("amixer", "sset", "Master", "5-").CombinedOutput()
	output, err = string(o), e
	return
}

func amixerSetOn() (output string, err error) {
	o, e := exec.Command("amixer", "sset", "Master", "on").CombinedOutput()
	output, err = string(o), e
	return
}

func amixerSetOff() (output string, err error) {
	o, e := exec.Command("amixer", "sset", "Master", "off").CombinedOutput()
	output, err = string(o), e
	return
}

func amixerSetOnOff() (output string, err error) {
	o, e := exec.Command("amixer", "sset", "Master", "toggle").CombinedOutput()
	output, err = string(o), e
	return
}

//acpi
func acpiGet() (output string, err error) {
	o, e := exec.Command("acpi").CombinedOutput()
	output, err = string(o), e
	return
}

func upowerGet() (output string, err error) {
	o, e := exec.Command("upower", "-d").CombinedOutput()
	output, err = string(o), e
	return
}
