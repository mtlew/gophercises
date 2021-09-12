package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type pathUrl struct {
	Path string `json:"path" yaml:"path"`
	Url  string `json:"url" yaml:"url"`
}

type unmarshalFunction func(data []byte, v interface{}) error

func getPathUrls(jsonFilename string, yamlFilename string) (map[string]string, error, error) {
	pathsToUrls := make(map[string]string)
	pathsToUrls, errJson := appendMap(pathsToUrls, jsonFilename, json.Unmarshal)
	pathsToUrls, errYaml := appendMap(pathsToUrls, yamlFilename, yaml.Unmarshal)

	return pathsToUrls, errJson, errYaml
}

func appendMap(pathsToUrls map[string]string, filename string, unmarshalFunction unmarshalFunction) (map[string]string, error) {
	if filename == "" {
		return pathsToUrls, nil
	}
	pathUrls, err := fileToMap(filename, unmarshalFunction)
	if err != nil {
		return pathsToUrls, err
	}
	for path, url := range pathUrls {
		pathsToUrls[path] = url
	}
	return pathsToUrls, nil
}

func fileToMap(filename string, unmarshalFunction unmarshalFunction) (map[string]string, error) {
	bytes, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	pathUrls, err := unmarshal(bytes, unmarshalFunction)
	if err != nil {
		return nil, err
	}
	return structsToMap(pathUrls), nil
}

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func unmarshal(bytes []byte, unmarshalFunction unmarshalFunction) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := unmarshalFunction(bytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func structsToMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}
