package sdcgo

import (
	"bytes"
	"encoding/gob"
)

// Serialize serializes the input into bytes.
func Serialize(input interface{}) ([]byte, error) {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	if err := encoder.Encode(input); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
