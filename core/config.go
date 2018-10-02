package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Loads config file
func LoadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var out Config
	err = json.Unmarshal(raw, &out)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return out
}

// Update config file
func UpdateConfig(new *Config, filename string) error {
	//Turn JSON into bytes for writing
	bytes, err := json.Marshal(new)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//Overwrite old file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	//Put JSON into new file
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = file.Sync()

	return err
}

// Retrieves webhook links from a config
func GetHook(c Config, name string) (string, string, string) {
	for _, n := range c.Webhooks {
		if strings.Compare(name, n.Name) == 0 {
			return n.Url, n.Token, n.Team
		}
	}
	return "", "", ""
}
