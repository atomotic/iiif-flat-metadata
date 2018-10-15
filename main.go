package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/utahta/go-openuri"
)

type Manifest struct {
	Description string `json:"description"`
	Id          string `json:"@id"`
	Label       string `json:"label"`
	Metadata    []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
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

	for _, property := range manifest.Metadata {
		if metadata[property.Label] != "" {
			metadata[property.Label] = fmt.Sprintf("%s | %s", metadata[property.Label], property.Value)
		} else {
			metadata[property.Label] = property.Value
		}
	}

	output, _ := json.Marshal(metadata)
	fmt.Println(string(output))
}
