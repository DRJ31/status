package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func getConfig() Config {
	configFile, err := os.Open("etc/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	data, _ := ioutil.ReadAll(configFile)

	var result Config
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
