package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func embedImage(id string, username string, filename string) ([]byte, error) {
	//open templates/embed.html
	file, err := ioutil.ReadFile("templates/embed.html")
	//println(string(file))

	tmpl, err := template.New("template").Parse(string(file))
	if err != nil {
		//println(err.Error())
		return nil, err
	}

	var config, _ = readConfigFile()

	var URL = config["URL"].(string)

	data := map[string]interface{}{
		"id":       id,
		"username": username,
		"filename": filename,
		"URL":      URL,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		//println(err.Error())
		return nil, err
	}

	//println(buf.String())

	return buf.Bytes(), nil
}

func renderIndexHTML() []byte {
	file, _ := ioutil.ReadFile("templates/index.html")
	tmpl, _ := template.New("template").Parse(string(file))

	var buf bytes.Buffer
	tmpl.Execute(&buf, nil)

	return buf.Bytes()
}
