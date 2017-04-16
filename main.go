package main

import (
	"fmt"

	"os"

	"github.com/ellcrys/util"
	runtime "github.com/ncodes/cocoon/core/runtime/golang"
)

var log = runtime.GetLogger()

// App is an example cocoon code
type App struct {
}

// OnInit method initializes the app
func (app *App) OnInit(link *runtime.Link) error {
	log.Info("App is initializing now now")
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
func (app *App) OnInvoke(link *runtime.Link, txID, function string, params []string) (interface{}, error) {
	log.Info("Invoke acknowledged")
	return "success", nil
}

func main() {
	runtime.Run(new(App))
}
