package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"net/http"
)

func main() {
	// fmt.Println("Hello Friend")
	// output, err := exec.Command("wget", `http://www.nairaland.com/`).Output()
	// if err != nil {
	// 	fmt.Println("Err:", err, string(output))
	// 	return
	// }

	// fmt.Println(len(output), output)

	res, err := http.Get("http://127.0.0.1:3000/hello")
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
