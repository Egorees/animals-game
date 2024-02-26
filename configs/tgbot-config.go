package configs

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type BotConfig struct {
	Timeout int `yaml:"Timeout"`
	Offset  int `yaml:"Offset"`
}

func ParseBotConfig(filepath string) *BotConfig {
	file, err := os.Open(filepath)
	if err != nil {
		slog.Error("Error during opening bot config: %v", err.Error())
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	parser := yaml.NewDecoder(file)
	var res BotConfig
	if err := parser.Decode(&res); err != nil {
		slog.Error("Error during parsing bot config: %v", err.Error())
		panic(err)
	}
	return &res
}
