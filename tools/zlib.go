package tools

import (
	"bytes"
	"compress/zlib"
	"io"
)

func ZlibInflate(src []byte) []byte {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
