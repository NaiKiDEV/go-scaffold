package scaffolder

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/NaiKiDEV/go-scaffold/config"
)

func Scaffold(cfg *config.Config, baseDir string, injectedVars *config.InjectedVariables) error {
	dirConfig := cfg.DirectoryConfig

	if injectedVars == nil {
		injectedVars = &map[string]string{}
	}

	for _, rootConfigDir := range dirConfig {
		err := recursivelyScaffoldFiles(cfg, rootConfigDir, baseDir, injectedVars)
		if err != nil {
			return err
		}
	}

	return nil
}

func recursivelyScaffoldFiles(cfg *config.Config, dir config.Directory, baseDirPath string, injectedVars *config.InjectedVariables) error {
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

		if file.TemplateFile != "" {
			templateFileContent, err := os.ReadFile(path.Join(cfg.TemplateRootPath, file.TemplateFile))
			if err != nil {
				return err
			}

			if err = writeToFile(createdFile, string(templateFileContent), injectedVars); err != nil {
				return err
			}
		}

		if file.TemplateFile == "" {
			if file.Template == "" {
				return fmt.Errorf("template file and template were not provided, not able to scaffold the file with name: %s", file.Name)
			}
			if err = writeToFile(createdFile, file.Template, injectedVars); err != nil {
				return err
			}
		}

		createdFile.Close()
	}

	for _, nestedDir := range dir.SubDirectories {
		if err := recursivelyScaffoldFiles(cfg, nestedDir, dirPath, injectedVars); err != nil {
			return err
		}
	}

	return nil
}

func writeToFile(file *os.File, template string, injectedVars *config.InjectedVariables) error {
	_, err := file.WriteString(replaceTemplateVarsInString(template, injectedVars))
	if err != nil {
		return err
	}
	return nil
}

func replaceTemplateVarsInString(str string, injectedVars *config.InjectedVariables) string {
	for key, value := range *injectedVars {
		str = strings.ReplaceAll(str, fmt.Sprintf("{{%s}}", key), value)
	}
	return str
}
