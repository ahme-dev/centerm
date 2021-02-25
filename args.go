package main

import (
	"os"
	"regexp"
	"strings"
)

//returns one arg or nothing if non-existent, rather than error
func getArg(sel uint8) string {
	var str string
	var argsNum = uint8(len(os.Args)) - 1

	//wrong
	switch {
	case sel > argsNum:
		str = ""
		return str
	case sel == 0:
		str = ""
		return str
	}

	//right
	str = os.Args[sel]
	return str
}

//returns the given input as one string according to specified length
func getArgsOneString(selLen uint8) string {
	var str string
	var argsNum = uint8(len(os.Args)) - 1

	if selLen > argsNum {
		selLen = argsNum
	}

	switch selLen {
	case 0:
		str = ""
	case 1:
		str = os.Args[1]
	default:
		var index uint8 = 1
		for index <= selLen {
			if str == "" {
				str = os.Args[index]
			} else {
				str = str + " " + os.Args[index]
			}
			index = index + 1
		}
	}

	return str
}

//returns the given hyphen input as map
func getArgsHyphen() (flagsArg []string) {
	var flagPattern = regexp.MustCompile(`(-\w{1,}) (\w{1,})`)
	var argsSlice = flagPattern.FindAllString(getArgsOneString(5), 2)
	var counter int

	for _,n := range argsSlice {
		cutSlice := strings.Split(n, " ")
		flagsArg = append(flagsArg, cutSlice[0])
		flagsArg = append(flagsArg, cutSlice[1])
		counter = counter + 2
	}
	return
}
