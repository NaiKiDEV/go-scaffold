package scaffolder

import (
	"fmt"
	"os"
	"path"

	"github.com/NaiKiDEV/go-scaffold/config"
)

func Scaffold(c config.Config, baseDir string) error {
	dirConfig := c.DirectoryConfig

	for _, rootConfigDir := range dirConfig {
		err := recursivelyGenerateFiles(rootConfigDir, baseDir)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: add custom error handling for readable errors
func recursivelyGenerateFiles(dir config.Directory, baseDir string) error {
	// TODO: dir.Name should be injectable {{projectName}}
	err := os.Mkdir(path.Join(baseDir, dir.Name), 0777)
	if err != nil {
		return err
	}

	for _, file := range dir.Files {
		// TODO: file.Name should be injectable {{projectName}}
		if file.Name == "" {
			return fmt.Errorf("file must have a name")
		}

		createdFile, err := os.Create(path.Join(baseDir, dir.Name, file.Name))
		if err != nil {
			return err
		}

		// TODO: add custom variable injection into template
		_, err = createdFile.WriteString(file.Template)
		if err != nil {
			return err
		}

		createdFile.Close()
	}

	for _, nestedDir := range dir.SubDirectories {
		if err := recursivelyGenerateFiles(nestedDir, path.Join(baseDir, dir.Name)); err != nil {
			return err
		}
	}

	return nil
}
