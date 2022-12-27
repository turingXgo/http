package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Settings domain with application settings variables
type Settings struct {
	DBName string `yaml:"db_name"`
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
}

// ParseYAML parses given yaml file and returns Settings instance
func ParseYAML(fileName string) (Settings, error) {
	body, err := os.ReadFile(fileName)
	if err != nil {
		return Settings{}, err
	}
	settings := Settings{}
	if err := yaml.Unmarshal(body, &settings); err != nil {
		return settings, err
	}
	return settings, nil
}
