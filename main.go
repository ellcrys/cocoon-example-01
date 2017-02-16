package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello Friend")
	output, err := exec.Command("/bin/sh", "-c", "'curl https://google.com'").Output()
	if err != nil {
		fmt.Println("Err:", err, string(output))
		return
	}

	fmt.Println(string(output))

}
