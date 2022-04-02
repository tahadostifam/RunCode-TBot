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
}

func ReadAndSetConfigs() *Configs {
	cwd, cwd_err := os.Getwd()
	if cwd_err != nil {
		fmt.Println("An error occurred on getting cwd")
		os.Exit(1)
	} else {
		configs_file, configs_file_err := os.ReadFile(cwd + "/configs/configs.yml")
		if configs_file_err != nil {
			fmt.Println("An error occurred on reading configs.yml file")
			os.Exit(1)
		} else {
			parse_err := yaml.Unmarshal(configs_file, &AllConfigs)
			if parse_err != nil {
				fmt.Println("Error in parsing configs.yml file")
				os.Exit(1)
			} else {
				return &AllConfigs
			}
		}
	}
	return &AllConfigs
}
