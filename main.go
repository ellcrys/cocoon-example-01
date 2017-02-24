package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	stub "github.com/ncodes/cocoon/core/stubs/golang"
)

func main() {

	stub.StartServer(func() {
		fmt.Println("Hello Friend")
		res, err := http.Get("http://google.com")
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
	})

	// output, err := exec.Command("bash", "-c", "iptables -S").Output()
	// if err != nil {
	// 	fmt.Println("Err:", err, string(output))
	// 	return
	// }

	// fmt.Println(len(output), string(output))
}
