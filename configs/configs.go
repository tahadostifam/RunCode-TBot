package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var AllConfigs Configs

type Configs struct {
	BOT_TOKEN string `yaml:"BOT_TOKEN"`
	SSH_HOST  string `yaml:"SSH_HOST"`
	SSH_PORT  string `yaml:"SSH_PORT"`
	SSH_USER  string `yaml:"SSH_USER"`
	SSH_PASS  string `yaml:"SSH_PASS"`
	ENV       string `yaml:"ENV"`
}

func ReadAndSetConfigs() *Configs {
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		fmt.Println("An error occurred on getting cwd")
		os.Exit(1)
	} else {
		configsFile, configsFileErr := os.ReadFile(cwd + "/configs/configs.yml")
		if configsFileErr != nil {
			fmt.Println("An error occurred on reading configs.yml file")
			os.Exit(1)
		} else {
			parseErr := yaml.Unmarshal(configsFile, &AllConfigs)
			if parseErr != nil {
				fmt.Println("Error in parsing configs.yml file")
				os.Exit(1)
			} else {
				return &AllConfigs
			}
		}
	}
	return &AllConfigs
}
