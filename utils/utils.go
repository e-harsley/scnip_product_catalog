package utils

import (
	"encoding/json"
	"io"
)

func BindDataOperationStruct(reader io.Reader, output interface{}) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(output)
}
