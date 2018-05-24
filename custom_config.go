package proxy

import (
	"errors"
	"fmt"
)

type CustomConfig map[string]interface{}

func (cc *CustomConfig) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("key can not be empty")
	}

	v, ok := (*cc)[key]
	if !ok {
		return nil, fmt.Errorf("key %s is not valid", key)
	}
	return v, nil
}
