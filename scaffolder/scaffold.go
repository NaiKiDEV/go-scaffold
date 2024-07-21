package scaffolder

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/NaiKiDEV/go-scaffold/config"
)

func Scaffold(c config.Config, baseDir string, injectedVars config.InjectedVariables) error {
	dirConfig := c.DirectoryConfig

	if injectedVars == nil {
		injectedVars = make(map[string]string, 0)
	}

	for _, rootConfigDir := range dirConfig {
		err := recursivelyScaffoldFiles(rootConfigDir, baseDir, injectedVars)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: add custom error handling for readable errors
func recursivelyScaffoldFiles(dir config.Directory, baseDirPath string, injectedVars config.InjectedVariables) error {
	dirPath := path.Join(baseDirPath, replaceTemplateVarsInString(dir.Name, injectedVars))
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	for _, file := range dir.Files {
		if file.Name == "" {
			return fmt.Errorf("file must have a name")
		}

		createdFile, err := os.Create(path.Join(dirPath, replaceTemplateVarsInString(file.Name, injectedVars)))
		if err != nil {
			return err
		}

		_, err = createdFile.WriteString(replaceTemplateVarsInString(file.Template, injectedVars))
		if err != nil {
			return err
		}

		createdFile.Close()
	}

	for _, nestedDir := range dir.SubDirectories {
		if err := recursivelyScaffoldFiles(nestedDir, dirPath, injectedVars); err != nil {
			return err
		}
	}

	return nil
}

func replaceTemplateVarsInString(str string, injectedVars config.InjectedVariables) string {
	for key, value := range injectedVars {
		str = strings.ReplaceAll(str, fmt.Sprintf("{{%s}}", key), value)
	}
	return str
}
