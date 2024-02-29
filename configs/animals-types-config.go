package configs

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Animal struct {
	AnimalType string `yaml:"AnimalType"`
}

type AnimalCfg struct {
	AnimalsTypes []*Animal `yaml:"AnimalsTypes"`
}

type AnimalTypeFromIdToStr []string

func ParseAnimalsType(filepath string) *AnimalCfg {
	file, err := os.Open(filepath)
	if err != nil {
		slog.Error("Error during opening stickers config: %v", err.Error())
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	parser := yaml.NewDecoder(file)
	var res AnimalCfg
	if err = parser.Decode(&res); err != nil {
		slog.Error("Error during parsing stickers config: %v", err.Error())
		panic(err)
	}
	return &res
}

type AnimalTypeFromStrToId map[string]int16

func GetAnimalsTYpeListFromIdToStr(cfg *AnimalCfg) AnimalTypeFromIdToStr {
	var res AnimalTypeFromIdToStr

	for _, animal := range cfg.AnimalsTypes {
		res = append(res, animal.AnimalType)
	}
	return res
}

func GetAnimalsTypeListFromStrToId(idToStr *AnimalTypeFromIdToStr) AnimalTypeFromStrToId {
	res := AnimalTypeFromStrToId{}
	for id, str := range *idToStr {
		res[str] = int16(id)
	}
	return res
}
