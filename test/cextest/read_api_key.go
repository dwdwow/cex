package cextest

import (
	"github.com/dwdwow/cex"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	apiKeyFileRelativePath = "/cex/key/apikey.yml"
)

func ReadApiKey() (map[cex.Cex]cex.Api, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := dirname + apiKeyFileRelativePath
	fileByte, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	data := map[cex.Cex]cex.Api{}
	err = yaml.Unmarshal(fileByte, &data)
	return data, err
}

func MustReadApiKey() map[cex.Cex]cex.Api {
	apiKey, err := ReadApiKey()
	if err != nil {
		panic(err)
	}
	return apiKey
}
