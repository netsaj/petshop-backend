package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadFileJSON(file string, data interface{}) error {
	var jsonFile, err = os.Open(file)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return err
	}
	return nil
}
