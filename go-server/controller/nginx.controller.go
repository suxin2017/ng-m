package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
	"suxin2017.com/server/utils"
)

func Reload() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func GetNginxStatus(c *gin.Context) {
	isRunning := utils.NginxPid()
	if isRunning {
		c.JSON(200, utils.Ok(string("running")))
	} else {
		c.JSON(200, utils.Ok("stop"))
	}

}

func StartNginx(c *gin.Context) {
	if err := utils.StartNginx(); err == nil {
		c.JSON(200, utils.OkMessage())

	} else {
		c.Error(fmt.Errorf("操作失败 %v", err))
	}

}
func ReloadNginx(c *gin.Context) {
	if err := utils.ReloadNginx(); err == nil {
		c.JSON(200, utils.OkMessage())

	} else {
		c.Error(fmt.Errorf("操作失败 %v", err))
	}

}
func StopNginx(c *gin.Context) {
	if err := utils.StopNginx(); err == nil {
		c.JSON(200, utils.OkMessage())

	} else {
		c.Error(fmt.Errorf("操作失败 %v", err))
	}

}
