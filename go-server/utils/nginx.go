package utils

import (
	"os/exec"

	"suxin2017.com/server/constants"
)

func StartNginx() {
	cmd := exec.Command("bash", "start-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.Output()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))

}

func StopNginx() {
	cmd := exec.Command("bash", "stop-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.Output()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))

}

func ReloadNginx() {
	cmd := exec.Command("bash", "reload-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.Output()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))

}

func TestNginx() {
	cmd := exec.Command("bash", "test-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.Output()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))

}
