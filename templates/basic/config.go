package basic

import (
	_ "embed"

	"github.com/NaiKiDEV/go-scaffold/config"
)

//go:embed templates/basic.ts
var basicFileBytes []byte

func New() config.Config {
	return config.Config{
		DirectoryConfig: []config.Directory{
			{
				Name: "basic-example",
				Files: []config.File{
					{
						Name:     "{{featureName}}basic.ts",
						Template: string(basicFileBytes),
					},
				},
				SubDirectories: []config.Directory{
					{
						Name: "{{projectName}}nested-folder",
						Files: []config.File{
							{
								Name:     "basic.ts",
								Template: string(basicFileBytes),
							},
						},
					},
				},
			},
		},
	}
}
