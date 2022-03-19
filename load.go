package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"net/http"
)

func LoadFromReader(file io.Reader, destination interface{}) error {
	decoder := json.NewDecoder(file)
	return decoder.Decode(&destination)
}

func LoadFromPath(path string, destination interface{}) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	err = LoadFromReader(file, destination)
	if err != nil {
		return fmt.Errorf("error mapping json from path %v:\n	==> %v", path, err)
	}

	return nil
}

func LoadFromString(str string, destination interface{}) error {
	return LoadFromReader(strings.NewReader(str), destination)
}

func LoadFromURL(url string, destination interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return LoadFromReader(resp.Body, destination)
}
