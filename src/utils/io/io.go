package talkerutils

import (
	"encoding/json"
)

type FileReader interface {
	ReadFile(filePath string) ([]byte, error)
}

type FileWriter interface {
	WriteFile(filePath string, data []byte)
}

type IoUtils struct {
	fileReader FileReader
	fileWriter FileWriter
}

func (i *IoUtils) ReadFile(filePath string, conversor interface{}) {
	data, err := i.fileReader.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, conversor); err != nil {
		panic(err.Error())
	}
}

func (i *IoUtils) WriteFile(filePath string, data []byte) {
	i.fileWriter.WriteFile(filePath, data)
}

func New(fileReader FileReader, fileWriter FileWriter) *IoUtils {
	return &IoUtils{
		fileReader: fileReader,
		fileWriter: fileWriter,
	}
}
