package config

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

var NConfig NeptuneConfig

type NeptuneConfig struct {
	Locations LocationsConfig
}

type LocationsConfig struct {
	Spawn        [3]float64
	DmgThreshold float64
	DealerCoords [3]float64
}

func LoadNeptuneConfig() {
	conf, err := readNeptuneConfig()
	if err != nil {
		log.Fatal(err)
	}

	NConfig = conf
}

func readNeptuneConfig() (NeptuneConfig, error) {
	c := NeptuneConfig{}
	var zero NeptuneConfig
	if _, err := os.Stat("neptune.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("neptune.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c, nil
	}
	data, err := os.ReadFile("neptune.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c, nil
}
