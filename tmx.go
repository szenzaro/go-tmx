package tmx

import (
	"encoding/xml"
	"io/ioutil"
)

// Read a tmx file given its path
func Read(filepath string) (*Tmx, error) {
	byteValue, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var tmxData Tmx
	if err := xml.Unmarshal(byteValue, &tmxData); err != nil {
		return nil, err
	}
	return &tmxData, nil
}
