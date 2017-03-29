package main

import (
	"fmt"

	"github.com/ellcrys/util"
	runtime "github.com/ncodes/cocoon/core/runtime/golang"
)

var log = runtime.GetLogger()

// App is an example cocoon code
type App struct {
}

// OnInit method initializes the app
func (app *App) OnInit(link *runtime.Link) error {
	log.Info("App is initializing!")

	ln := util.RandString(5)
	log.Info("Ledger: ", ln)
	link.CreateLedger(ln, true, false)
	link.PutIn(ln, "account.ken", []byte("10.29"))
	link.PutIn(ln, "account.ben", []byte("3.11"))
	link.PutIn(ln, "account.zen", []byte("4.33"))
	rg := link.NewRangeGetterFrom(ln, "account", "account.z", true)
	for rg.HasNext() {
		util.Printify(rg.Next())
	}

	// log.Info("Connector addrss: ", os.Getenv("CONNECTOR_RPC_ADDR"))
	// log.Info("Cocoon ID: ", runtime.GetCocoonID())
	// log.Info("Linked To: ", runtime.GetID())
	return nil
}

// OnInvoke process invoke transactions
func (app *App) OnInvoke(link *runtime.Link, txID, function string, params []string) (interface{}, error) {
	log.Info("ID: ", txID, function, params)
	return "success", fmt.Errorf("Something happened")
}

func main() {
	runtime.Run(new(App))

	// fmt.Println("Hello Friend")
	// res, err := http.Get("https://google.com.ng")
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// 	return
	// }

	// defer res.Body.Close()
	// fmt.Println(res.StatusCode)
	// _, err = io.Copy(os.Stdout, res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// output, err := exec.Command("bash", "-c", "iptables -S").Output()
	// if err != nil {
	// 	fmt.Println("Err:", err, string(output))
	// 	return
	// }

	// fmt.Println(len(output), string(output))
}
