package api

import (
	"encoding/json"
)

func FromMap[T interface{}](m map[string]interface{}) (T, error) {
	var d T
	b, err := json.Marshal(m)

	if err != nil {
		return d, err
	}

	err = json.Unmarshal(b, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func FromJSON[T interface{}](body []byte) (T, error) {
	var d T
	err := json.Unmarshal(body, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func ToJSON(serializable interface{}) ([]byte, error) {
	return json.Marshal(serializable)
}
