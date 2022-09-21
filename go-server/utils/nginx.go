package utils

import (
	"errors"
	"os"
	"os/exec"
	"path"

	"suxin2017.com/server/constants"
)

func StartNginx() error {
	cmd := exec.Command("bash", "start-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))
	return err

}

func StopNginx() error {
	cmd := exec.Command("bash", "stop-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))
	return err
}

func ReloadNginx() error {
	cmd := exec.Command("bash", "reload-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))
	return err
}

func TestNginx() {
	cmd := exec.Command("bash", "test-nginx.sh")
	cmd.Dir = constants.GetScriptDir()
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("error", err.Error())
	}
	println(string(out[:]))

}

func NginxPid() bool {
	dir := constants.GetNginxConfigDir()
	pid := path.Join(dir, "nginx.pid")
	if _, err := os.Stat(pid); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		return false
	}
	return true
}
