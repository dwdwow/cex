package cex

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	apiKeyFileRelativePath = "/cex/key/apikey.yml"
)

/**
JUST FOR TEST
*/

func ReadApiKey() (map[Name]Api, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := dirname + apiKeyFileRelativePath
	fileByte, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	data := map[Name]Api{}
	err = yaml.Unmarshal(fileByte, &data)
	return data, err
}

func MustReadApiKey() map[Name]Api {
	apiKey, err := ReadApiKey()
	if err != nil {
		panic(err)
	}
	return apiKey
}
