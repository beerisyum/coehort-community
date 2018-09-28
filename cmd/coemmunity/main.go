package main

import (
	"github.com/erei/coemmunity/pkg/config"
	"github.com/erei/coemmunity/pkg/web"

	"flag"
	"fmt"
	"path/filepath"
)

var (
	Port   string
	Static string
	Config string
)

func init() {
	flag.StringVar(&Port, "p", "3000", "Port of your server.")
	flag.StringVar(&Static, "s", "../../public", "Path to your static content")
	flag.StringVar(&Config, "c", "./config.json", "Path to configuration file.")

	flag.Parse()
}

func main() {

	viewPath, err := filepath.Abs(Static + "/views")

	if err != nil {
		fmt.Printf("An error occurred when retrieving the view directory's path: %s\n", err.Error())
	}

	staticPath, err := filepath.Abs(Static + "/static")

	if err != nil {
		fmt.Printf("An error occurred when retrieving the static directory's path: %s\n", err.Error())
	}

	config, err := config.LoadConfig(Config)

	if err != nil {
		fmt.Printf("An error occurred when loading configuration file: %s\n", err.Error())
	}

	server := web.NewServer(staticPath, viewPath, config)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", Port)))
}
