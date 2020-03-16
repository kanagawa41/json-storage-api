package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

// Stock is DBのレコード
type Stock struct {
	UUID string
	JSON string
}

var (
	baseURL = "/tmp/stock_api/stocks"
)

func isExistFile(path string) (flag bool, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func selectStock(uuid string) (s Stock, err error) {
	filePath := baseURL + "/" + uuid

	if rel, err := isExistFile(filePath); !rel {
		return s, err
	}

	s.UUID = uuid

	dat, _ := ioutil.ReadFile(filePath)
	s.JSON = string(dat)

	return s, nil
}

func createStock() (s Stock, err error) {
	uuid := uuid.New()
	s.UUID = uuid.String()
	filePath := baseURL + "/" + s.UUID

	// NOTE: Basically This won't be happen.
	if rel, _ := isExistFile(filePath); rel {
		return s, errors.New("already exist specified file")
	}

	if err := os.MkdirAll(baseURL, 0644); err != nil {
		return s, err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return s, err
	}
	defer f.Close()

	return s, nil
}

func updateStock(uuid string, json string) (bool, error) {
	filePath := baseURL + "/" + uuid
	if rel, _ := isExistFile(filePath); !rel {
		return false, errors.New("Don't exist one")
	}

	err := ioutil.WriteFile(filePath, []byte(json), 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

func deleteStock(uuid string) (bool, error) {
	filePath := baseURL + "/" + uuid
	if rel, _ := isExistFile(filePath); !rel {
		return false, errors.New("Don't exist one")
	}

	if err := os.Remove(filePath); err != nil {
		return false, err
	}

	return true, nil
}
