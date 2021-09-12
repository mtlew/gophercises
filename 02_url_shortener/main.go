package main

import (
	"flag"
	"log"
)

func main() {
	jsonFilename := flag.String("json", "", "Redirects json filename")
	yamlFilename := flag.String("yaml", "", "Redirects yaml filename")

	flag.Parse()

	pathUrlsMap, errJson, errYaml := getPathUrls(*jsonFilename, *yamlFilename)

	if errJson != nil {
		log.Println(errJson)
	}
	if errYaml != nil {
		log.Println(errYaml)
	}

	serve(pathUrlsMap)
}
