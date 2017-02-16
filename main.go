package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello Friend")
	output, err := exec.Command("/bin/sh", "-c", "curl http://127.0.0.1:3000").Output()
	if err != nil {
		fmt.Println("Err:", err, string(output))
		return
	}

	fmt.Println(string(output))

}
