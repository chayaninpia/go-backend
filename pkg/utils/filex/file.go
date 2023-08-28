package filex

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSONFile[T any](path string) (*T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fileData := new(T)
	if err := json.Unmarshal(bytes, fileData); err != nil {
		return nil, err
	}
	return fileData, err
}

func ReadJSONFileArray[T any](path string) ([]T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fileData := make([]T, 0)
	if err := json.Unmarshal(bytes, &fileData); err != nil {
		return nil, err
	}
	return fileData, err
}

func WriteFile(fileData any, pathWithFileName string) error {
	newBytes, err := json.Marshal(&fileData)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(pathWithFileName, newBytes, 0777); err != nil {
		return err
	}
	return nil
}
