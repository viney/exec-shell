package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "./exec.sh")

	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}

	fmt.Println(string(bytes))
}
