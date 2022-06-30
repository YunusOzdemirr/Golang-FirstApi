package utils

import (
	"errors"
	"io/ioutil"
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}
func ReadFile(fileName string) (string, error) {
	if !(IsEmpty(fileName)) {
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return "", errors.New("Dosya yolu uzunluğu 0 olamaz")

}
