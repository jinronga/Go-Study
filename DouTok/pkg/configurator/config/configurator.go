package config

import (
	"errors"
	"github.com/go-study/DouTok/pkg/configurator/constants"
	"os"
	"path/filepath"
)

const (
	ErrNotPtrOfStruct = "given config is not a pointer of a struct"
)

func InitConfig(t any, configName string) error {
	configPath, err := getConfigFilesPath(configName)
	if err != nil {
		return err
	}

	file, _ := os.ReadFile(configPath)
	if err := yaml.Unmarshal(file, t); err != nil {
		return err
	}

	return nil
}

func getConfigFilesPath(configName string) (string, error) {
	pathList := [5]string{
		"./config/",
		"../../config/",
		"../../../config/",
		"../../../../config/",
		"../../../../../config/",
	}

	for i := range pathList {
		_, err := os.Stat(pathList[i] + configName)
		if err == nil {
			p, _ := filepath.Abs(pathList[i] + configName)
			return p, nil
		}
	}

	return "", errors.New(constants.ErrConfigFileNotFound + ", file name: " + configName)
}
