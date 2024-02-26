package configs

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type AnimalSticker struct {
	AnimalType string `yaml:"AnimalType"`
	StickerID  string `yaml:"StickerId"`
}

type StickersConfig struct {
	Stickers       []*AnimalSticker `yaml:"AnimalsStickers"`
	StickerExample string           `yaml:"StickerExample"`
}

func ParseStickersConfig(filepath string) StickersConfig {
	file, err := os.Open(filepath)
	if err != nil {
		slog.Error("Error during opening stickers config: %v", err.Error())
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	parser := yaml.NewDecoder(file)
	var res StickersConfig
	if err := parser.Decode(&res); err != nil {
		slog.Error("Error during parsing stickers config: %v", err.Error())
		panic(err)
	}
	return res
}

func (stickersCfg *StickersConfig) GetAnimaTypeByStickerID(stickerID string) string {
	for _, sticker := range stickersCfg.Stickers {
		if sticker.StickerID == stickerID {
			return sticker.AnimalType
		}
	}
	return "notCool:(" // todo: think how handle other stickers
}
