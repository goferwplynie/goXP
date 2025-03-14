package config

import (
	"encoding/json"
	"os"
)

func ConfigLoader() Config {
	var conf Config
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}
