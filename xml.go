package common

import (
	"encoding/xml"
	"io/ioutil"
)

func ReadXmlToStruct(path string, config interface{}) error {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(buffer, &config)
	if err != nil {
		return err
	}
	return nil
}
