package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type Config struct {
	ProjectName string
	KeyPair     string
}

func LoadConfig(file string) (Config, error) {
	d, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("Error loading file: %s", file)
		return Config{}, err
	}

	var out map[string]interface{}
	hcl.Decode(&out, string(d))

	return Config{out["project-name"].(string), out["key-pair"].(string)}, nil
}
