package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func readConfigFile() (map[string]interface{}, error) {
	//open config.json
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileContents string
	for scanner.Scan() {
		fileContents += scanner.Text()
	}

	//parse json
	jsonData := []byte(fileContents)
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
