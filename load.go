package gojson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
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
