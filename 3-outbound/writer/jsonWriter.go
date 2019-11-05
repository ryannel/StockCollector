package writer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"stockCollector/2-core/outboundProviders"
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

func (writer *JsonWriter) WriteCompany(company outboundProviders.Company, fileName string) error {
	if writer.outputDir == "" {
		return errors.New("output directory must be set")
	}

	err := os.MkdirAll(writer.outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	jsonByte, err := json.Marshal(company)
	return writer.writeFile(jsonByte, fileName)
}

func (writer *JsonWriter) WriteCompanies(companies []outboundProviders.Company, fileName string) error {
	if writer.outputDir == "" {
		return errors.New("output directory must be set")
	}

	err := os.MkdirAll(writer.outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	jsonByte, err := json.Marshal(companies)
	if err != nil {
		return err
	}

	return writer.writeFile(jsonByte, fileName)
}

func (writer *JsonWriter) writeFile(json []byte, fileName string) error {
	outputFilePath := filepath.Join(writer.outputDir, fileName)
	return ioutil.WriteFile(outputFilePath, json, os.ModePerm)
}