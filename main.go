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
func (app *App) OnInit() error {
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
func (app *App) OnInvoke(header runtime.Metadata, function string, params []string) ([]byte, error) {

	fmt.Println("Function: ", function)
	fmt.Println("Params: ", params)
	util.Printify(header)

	// link.NewLedger("myledger2", true, true)
	// log.Infof("Hello, I am %s", runtime.GetCocoonID())
	// tx, err := link.Put("myledger2", "account.balance", []byte("10.20"))
	// if err != nil {
	// 	log.Info("failed to create transaction = ", err)
	// } else {
	// 	fmt.Println("Cur: ", tx)
	// 	go func() {
	// 		txRevised, err := link.PutSafe(tx.ID, "myledger2", "account.balance", []byte("20.40"))
	// 		fmt.Println("Revised 1: ", txRevised, err)
	// 	}()
	// 	go func() {
	// 		// time.Sleep(5 * time.Second)
	// 		txRevised2, err := link.PutSafe(tx.ID, "myledger2", "account.balance", []byte("40.40"))
	// 		fmt.Println("Revised 2: ", txRevised2, err)
	// 	}()
	// }

	// start := time.Now()
	// l2 := runtime.NewLink("u1")
	// lock, err := l2.Lock("some_key", 2*time.Second)
	// print(err, lock)
	// err = lock.Acquire()
	// print(err, nil)
	// fmt.Println(time.Since(start))

	// start = time.Now()
	// lock2, err := l2.Lock("some_key", 10*time.Second)
	// fmt.Println(time.Since(start))
	// print(err, lock2)
	// err = lock2.Acquire()
	// print(err, nil)

	// print(err, nil)
	// err = lock.Release()
	// print(err, nil)
	// err = lock.IsAcquirer()
	// print(err, nil)

	// go func() {
	// 	tx, err := link.Put("myledger2", "account.ben", []byte("10.20"))
	// 	fmt.Print("Err 1: ")
	// 	print(err, tx.ID)
	// }()

	// go func() {
	// 	time.Sleep(8 * time.Second)
	// 	fmt.Println("Getting.....")
	// 	tx, err := link.Get("myledger2", "account.ben")
	// 	fmt.Print("Err 2: ")
	// 	print(err, tx)
	// }()

	return []byte("success"), nil
}

func main() {
	runtime.Run(new(App))
}
