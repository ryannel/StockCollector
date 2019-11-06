package writer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func New(outputDir string) (JsonWriter, error) {
	if outputDir == "" {
		return JsonWriter{}, errors.New("output directory must be set")
	}

	return JsonWriter{
		outputDir: outputDir,
	}, nil
}

type JsonWriter struct {
	outputDir string
}

func (writer *JsonWriter) WriteJson(companies interface{}, fileName string) error {
	if writer.outputDir == "" {
		return errors.New("output directory must be set")
	}

	jsonByte, err := json.Marshal(companies)
	if err != nil {
		return err
	}

	err = os.MkdirAll(writer.outputDir, os.ModePerm)
	if err != nil {
		return err
	}
	outputFilePath := filepath.Join(writer.outputDir, fileName)

	return ioutil.WriteFile(outputFilePath, jsonByte, os.ModePerm)
}

func (writer *JsonWriter) ReadJson(fileName string, t interface{}) error {
	if writer.outputDir == "" {
		return errors.New("output directory must be set")
	}

	readFilePath := filepath.Join(writer.outputDir, fileName)
	jsonByte, _ := ioutil.ReadFile(readFilePath)

	err := json.Unmarshal(jsonByte, &t)
	return err
}