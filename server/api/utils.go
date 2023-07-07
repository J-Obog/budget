package api

import (
	"encoding/json"
)

func FromMap[T any](m map[string]any) (T, error) {
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

func FromJSON[T any](body []byte) (T, error) {
	var d T
	err := json.Unmarshal(body, d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func ToJSON(serializable any) ([]byte, error) {
	return json.Marshal(serializable)
}
