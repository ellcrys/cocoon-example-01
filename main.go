package main

import (
	"fmt"

	"os"

	"github.com/ellcrys/util"
	"github.com/kr/pretty"
	runtime "github.com/ncodes/cocoon/core/runtime/golang"
)

var log = runtime.GetLogger()

// App is an example cocoon code
type App struct {
}

// OnInit method initializes the app
func (app *App) OnInit() error {
	log.Info("App is initializing now")
	log.Infof("Version: %s", os.Getenv("COCOON_CODE_VERSION"))
	return nil
}

func print(err error, v interface{}) {
	if err != nil {
		fmt.Println(err)
		return
	}
	util.Printify(v)
}

// OnInvoke process invoke transactions
func (app *App) OnInvoke(header runtime.Metadata, function string, params []string) ([]byte, error) {

	fmt.Println("Found: ", function)
	fmt.Println("Found Params: ", params)
	fmt.Println("MY_VAR", os.Getenv("MY_VAR"))
	fmt.Println("MY_VAR2", os.Getenv("MY_VAR2"))
	pretty.Println(header)
	a := 0
	panic(4 / a)

	return []byte("success"), nil
}

// OnStop is called when the cocoon code is being stopped
func (app *App) OnStop() {
	log.Info("I am stopping! :(")
}

func main() {
	runtime.Run(new(App))
}
