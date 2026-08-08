package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/178619/stube/pkg/app"
)

func main() {
	var fConfig *string = flag.StringP("config", "c", "config.json", "Path to the config file")
	var fHost *string = flag.StringP("host", "H", "", "Hostname / IP")
	var fPort *int = flag.IntP("port", "p", 0, "Port number")
	flag.Parse()
	cfg := app.DefaultConfig()
	err := cfg.ReadFile(*fConfig)
	if err != nil && !(*fConfig == "config.json" && os.IsNotExist(err)) {
		log.Fatal(err)
	}
	if *fHost != "" {
		cfg.Server.Host = *fHost
	}
	if *fPort > 0 {
		cfg.Server.Port = *fPort
	}
	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Local server: http://%s", addr)
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
