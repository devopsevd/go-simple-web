package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/devopsevd/go-simple-web/daemon"
)

var assetsPath string

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "0.0.0.0:3000", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectString, "db-connect", "postgres://postgres:example@postgres_db_1/gowebapp?sslmode=disable", "DB Connect String")
	//flag.StringVar(&cfg.Db.ConnectString, "db-connect", "postgres://postgres:example@localhost/gowebapp?sslmode=disable", "DB Connect String")
	flag.StringVar(&assetsPath, "assets-path", "/assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *daemon.Config) {
	log.Printf("Assets served from %q.", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()

	setupHttpAssets(cfg)

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
