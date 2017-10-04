package main

import "os/exec"

func killWithMessage(message string) {
	println(message)
	exec.Command("pkill", "-9", "presento").Output()
}
