package main

import (
	"flag"
	"log"
	"fmt"
	"net/http"
	"os"
	"handlers/admin"
	"handlers/utils"
	"handlers/mobile"
	"handlers/accounts"
)

type Base struct {
	A1 int
}

func (self Base) Print() {
	log.Printf("Log in base: %d", self.A1)
}

func (self Base) Do() {
	log.Printf("Do in base: %#v", self)
	self.Print()
}

type Derived struct {
	Base
}

func (self Derived) Print() {
	log.Printf("Log in Derived: %d", self.A1)
}

func main() {
	//log.SetFlags(log.Lshortfile | log.LstdFlags)
	//a := Derived{
	//}
	//log.Printf("a: %#v", a)
	//a.Print()
	//a.Base.Print()
	//a.Do()
	//return
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	defaultConfig := "/config/config.yaml"
	if tmp, set := os.LookupEnv("CONFIG"); set {
		defaultConfig = tmp
	}
	configFile := flag.String("config", defaultConfig, "Configuration file.")
	flag.Parse()

	opt, err := EnvironmentInitialize(*configFile)
	if err != nil {
		panic(err)
	}
	defer EnvironmentFinish()
	router := InitRouter()

	admin.Register(router)
	utils.Register(router, opt.UGC.Dir, opt.UGC.Prefix)
	mobile.Register(router)
	accounts.Register(router)

	http.Handle("/", router)
	log.Printf("Run on: %d\n", opt.Port)
	server := fmt.Sprintf(":%d", opt.Port)
	log.Fatal(http.ListenAndServe(server, nil))
}
