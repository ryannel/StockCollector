package writer

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"stockCollector/2-core/outboundProviders"
)

func NewJsonLinesWriter(outputFilePath string) (JsonLinesWriter, error) {
	if outputFilePath == "" {
		return JsonLinesWriter{}, errors.New("output file path must be set")
	}

	return JsonLinesWriter{
		outputFilePath: outputFilePath,
	}, nil
}

type JsonLinesWriter struct {
	outputFilePath string
}

func (writer *JsonLinesWriter) AppendLine(obj interface{}) error {
	if writer.outputFilePath == "" {
		return errors.New("output file path must be set")
	}

	err := os.MkdirAll(filepath.Dir(writer.outputFilePath), os.ModePerm)
	if err != nil {
		return err
	}

	file, _ := os.OpenFile(writer.outputFilePath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(obj)
}

func (writer *JsonLinesWriter) ReadCompanies() ([]outboundProviders.Company, error) {
	if writer.outputFilePath == "" {
		return nil, errors.New("output file path must be set")
	}

	file, err := os.Open(writer.outputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var results []outboundProviders.Company
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			err = nil // End of file
			break
		}
		if err != nil {
			break
		}

		company := outboundProviders.Company{}
		err = json.Unmarshal(line, &company)
		if err != nil {
			return nil, err
		}
		results = append(results, company)
	}

	return results, nil
}