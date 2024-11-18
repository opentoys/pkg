package jsonx

import (
	"bytes"
	"encoding/json"
)

var marshal = func(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	e := enc.Encode(v)
	return buf.Bytes(), e
}

var unmarshal = json.Unmarshal

func ReplaceMarshal(fn func(any) ([]byte, error)) {
	marshal = fn
}

func ReplaceUnmarshal(fn func([]byte, any) error) {
	unmarshal = fn
}

func Marshal(v any) ([]byte, error) {
	return marshal(v)
}

func Copy(dst, src any) (e error) {
	buf, e := marshal(src)
	if e != nil {
		return
	}
	return unmarshal(buf, dst)
}

func Unmarshal(buf []byte, v any) (e error) {
	return unmarshal(buf, v)
}

func Stringify(v any) string {
	return string(BytesNoError(v))
}

func BytesNoError(v any) (buf []byte) {
	buf, _ = marshal(v)
	return
}
