package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {

	for true {
		cmd := exec.Command("curl", "--user", "admin:WelcometoCX01", "10.0.1.17:8083/ZAutomation/api/v1/devices/ZWayVDev_zway_2-0-37/command/on")
		cmd2 := exec.Command("curl", "--user", "admin:WelcometoCX01", "10.0.1.17:8083/ZAutomation/api/v1/devices/ZWayVDev_zway_2-0-37/command/off")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
		err = cmd2.Run()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}
