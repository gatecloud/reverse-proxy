package proxy

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig []Server

func (sc *ServerConfig) LoadFile(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, sc)
	if err != nil {
		return err
	}

	return nil
}
