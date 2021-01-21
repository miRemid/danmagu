package tools

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(data []byte, src interface{}) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(src)
}
