package main

import (
	"fmt"

	"github.com/franela/goreq"
)

func main() {
	// fmt.Println("Hello Friend")
	// output, err := exec.Command("wget", `http://www.nairaland.com/`).Output()
	// if err != nil {
	// 	fmt.Println("Err:", err, string(output))
	// 	return
	// }

	// fmt.Println(len(output), output)

	res, err := goreq.Request{
		Uri: "http://google.com",
	}.Do()

	if err != nil {
		fmt.Println("Err: ", err)
	}

	fmt.Println(res.Body.ToString())
	fmt.Println(res.StatusCode)

}
