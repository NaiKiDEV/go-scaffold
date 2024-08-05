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
	var (
		baseDir   = flag.String("baseDir", "./", "Provide relative path to the directory, where the files should be scaffolded")
		templName = flag.String("templName", "", "Provide template name which should be used to scaffold you project")
		templFile = flag.String("templFile", "", "Provide template configuration file which should be used to scaffold you project (prioritized over templName)")
	)
	flag.Parse()

	var cfg *config.Config
	if *templFile == "" {
		if *templName == "" {
			log.Fatal("template file and template was not provided, nothing to scaffold")
		}
		switch *templName {
		case "basic":
			cfg = basic.New()
		default:
			log.Fatalf("template is not supported: %s.", *templName)
		}
	} else {
		cfgFromFile, err := config.ReadConfigFromFile(*templFile)
		if err != nil {
			log.Fatal(err)
		}
		cfg = cfgFromFile
	}

	var injVars *config.InjectedVariables
	if cfg.TemplateVarsPath != "" {
		varsFromFile, err := config.ReadInjectedVariablesFromFile(cfg.TemplateVarsPath)
		if err != nil {
			log.Fatal(err)
		}
		injVars = varsFromFile
	}
	// TODO: Reading from ENV vars when file is not provided?

	if err := scaffolder.Scaffold(cfg, *baseDir, injVars); err != nil {
		log.Fatalf("error occurred during scaffolding: [%s]", err)
	}

	fmt.Print("successfully scaffolded the project")
}
