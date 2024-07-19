package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/NaiKiDEV/go-scaffold/config"
	"github.com/NaiKiDEV/go-scaffold/scaffolder"
	"github.com/NaiKiDEV/go-scaffold/templates/basic"
)

func main() {
	var baseDir = flag.String("baseDir", "./", "Provide relative path to the directory, where the files should be scaffolded")
	var templName = flag.String("templName", "basic", "Provide template name which should be used to scaffold you project")
	flag.Parse()

	var cfg config.Config
	switch *templName {
	case "basic":
		cfg = basic.New()
	default:
		fmt.Print("Template is not supported...")
		return
	}

	if err := scaffolder.Scaffold(cfg, *baseDir); err != nil {
		log.Fatalf("Error occurred during scaffolding: [%s]", err)
	}

	fmt.Printf("Successfully scaffolded the project from '%s' template", *templName)
}
