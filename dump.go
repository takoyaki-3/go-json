package gojson

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

func DumpToWriter(object interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)

	return encoder.Encode(object)
}

func DumpToFile(object interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	err = DumpToWriter(object, file)
	file.Close()

	return err
}

func DumpToString(object interface{}) (string, error) {

	writer := new(bytes.Buffer)

	err := DumpToWriter(object, writer)
	if err != nil {
		return "", err
	}
	return writer.String(), nil
}
