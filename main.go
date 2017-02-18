package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// output, err := exec.Command("bash", "-c", "iptables -S").Output()
	// if err != nil {
	// 	fmt.Println("Err:", err, string(output))
	// 	return
	// }

	// fmt.Println(len(output), string(output))

	res, err := http.Get("https://google.com.ng")
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
