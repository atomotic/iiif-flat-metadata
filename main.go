package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/utahta/go-openuri"
)

// preferred language for multilang properties
const lang = "en"

type Manifest struct {
	Id          string `json:"@id"`
	Label       string `json:"label"`
	Attribution string `json:"attribution"`
	Description string `json:"description"`
	Metadata    []struct {
		Label interface{} `json:"label"`
		Value string      `json:"value"`
	} `json:"metadata"`
}

type Label map[string]interface{}

func getMetadataLabel(label interface{}) (string, error) {
	switch label.(type) {
	case string:
		return label.(string), nil
	case []interface{}:
		for _, labels := range label.([]interface{}) {
			resultLabel := Label(labels.(map[string]interface{}))

			if resultLabel["@language"] == lang {
				return resultLabel["@value"].(string), nil
			} else {
				return "", nil
			}
		}
	}

	return "", errors.New("errors")
}

func main() {
	var manifest Manifest
	metadata := make(map[string]string)

	if len(os.Args) == 1 {
		fmt.Println("USAGE: iiif-flat-metadata {file.json | http://remote.manifest.json}")
		os.Exit(2)
	}

	o, err := openuri.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()
	m, _ := ioutil.ReadAll(o)

	json.Unmarshal(m, &manifest)

	metadata["@id"] = manifest.Id
	metadata["Label"] = manifest.Label
	metadata["Attribution"] = manifest.Attribution
	metadata["Description"] = manifest.Description

	for _, property := range manifest.Metadata {
		label, err := getMetadataLabel(property.Label)
		if err != nil {
			fmt.Println("err")
		}
		if label != "" {
			if metadata[label] != "" {
				metadata[label] = fmt.Sprintf("%s | %s", metadata[label], property.Value)
			} else {
				metadata[label] = property.Value
			}
		}
	}

	output, _ := json.Marshal(metadata)
	fmt.Println(string(output))
}
