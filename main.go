package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	// TODO: allow to provide this file, maybe templates could have predefined values?
	res, err := os.ReadFile("./thunks.scaffold.keys.json")
	if err != nil {
		log.Fatal(err)
	}

	injectedVars, err := config.NewInjectedVariables(res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", injectedVars)

	if err := scaffolder.Scaffold(cfg, *baseDir, injectedVars); err != nil {
		log.Fatalf("Error occurred during scaffolding: [%s]", err)
	}

	fmt.Printf("Successfully scaffolded the project from '%s' template", *templName)
}
