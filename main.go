package main

import (
	"fmt"

	"os"

	stub "github.com/ellcrys/cocoon/core/stub"
	"github.com/ellcrys/util"
)

var log = stub.GetLogger()

// App is an example cocoon code
type App struct {
}

// OnInit method initializes the app
func (app *App) OnInit() error {
	log.Info("App is initializing now ")
	log.Info("Source Dir: ", os.Getenv("SOURCE_DIR"))
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
func (app *App) OnInvoke(header stub.Metadata, function string, params []string) ([]byte, error) {
	switch function {
	case "create-account":
		return stub.Render("myview", nil)
	case "about":
		return stub.RenderString(`<view style="color: green; background-color: gray;" minHeight="102px" minWidth=300px>About Me!</view>`, nil)
	case "others":
		return stub.RenderString(`<view>Others like me</view>`, nil)
	default:
		return nil, fmt.Errorf("unexpected function '%s'", function)
	}
}

// OnStop is called when the cocoon code is being stopped
func (app *App) OnStop() {
	log.Info("Contract is stopping")
}

func main() {
	stub.Run(new(App))
}
