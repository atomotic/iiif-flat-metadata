package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/atomotic/go-openuri"
)

const lang = "en"
const separator = "|"

// Manifest ...
type Manifest struct {
	ID          string `json:"@id"`
	Label       string `json:"label"`
	Attribution string `json:"attribution"`
	Description string `json:"description"`
	Metadata    []struct {
		Label interface{} `json:"label"`
		Value interface{} `json:"value"`
	} `json:"metadata"`
	Thumbnail struct {
		ID string `json:"@id"`
	} `json:"thumbnail"`
}

// Label ...
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

func getMetadataValue(value interface{}) (string, error) {
	switch value.(type) {
	case string:
		return value.(string), nil
	case []interface{}:
		ret := []string{}
		for _, values := range value.([]interface{}) {
			ret = append(ret, fmt.Sprintf("%s", values))
		}
		return strings.Join(ret, separator), nil
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

	metadata["@id"] = manifest.ID
	metadata["Label"] = manifest.Label
	metadata["Attribution"] = manifest.Attribution
	metadata["Description"] = manifest.Description
	metadata["Thumbnail"] = manifest.Thumbnail.ID

	for _, property := range manifest.Metadata {
		label, err := getMetadataLabel(property.Label)
		if err != nil {
			fmt.Println("err")
		}

		value, _ := getMetadataValue(property.Value)

		if label != "" {
			if metadata[label] != "" {
				metadata[label] = fmt.Sprintf("%s %s %s", metadata[label], separator, value)
			} else {
				metadata[label] = value
			}
		}
	}

	output, _ := json.Marshal(metadata)
	fmt.Println(string(output))
}
