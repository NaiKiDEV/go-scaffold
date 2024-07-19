package scaffolder

import (
	"errors"
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

// TODO: custom error handling
func recursivelyGenerateFiles(dir config.Directory, baseDir string) error {
	err := os.Mkdir(path.Join(baseDir, dir.Name), 0777)
	if err != nil {
		return err
	}

	for _, file := range dir.Files {
		if file.Name == "" {
			return errors.New("file must have a name")
		}

		createdFile, err := os.Create(path.Join(baseDir, dir.Name, file.Name))
		if err != nil {
			return err
		}

		// TODO: add custom template variable injection into template
		_, err = createdFile.WriteString(file.Template)
		if err != nil {
			return err
		}
	}

	for _, nestedDir := range dir.SubDirectories {
		if err := recursivelyGenerateFiles(nestedDir, path.Join(baseDir, dir.Name)); err != nil {
			return err
		}
	}

	return nil
}
