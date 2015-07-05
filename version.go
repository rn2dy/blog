package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const versionSRC = "app.yaml"

func version() string {
	bytes, err := ioutil.ReadFile(versionSRC)
	if err != nil {
		log.Fatal(err)
	}
	v := struct {
		Version string
	}{}
	err = yaml.Unmarshal(bytes, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v.Version
}
