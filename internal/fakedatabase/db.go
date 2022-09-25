package fakedatabase

import "errors"

var db = make(map[string]string)

func Insert(k string, v string) error {
	db[k] = v
	return nil
}

func Query(k string) (*string, error) {
	value, ok := db[k]
	if !ok {
		return nil, errors.New("key not found in database")
	}
	return &value, nil
}
