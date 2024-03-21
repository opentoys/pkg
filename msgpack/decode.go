package msgpack

import (
	"io"

	"github.com/opentoys/pkg/msgpack/internal/decoding"
	streamdecoding "github.com/opentoys/pkg/msgpack/internal/stream/decoding"
)

// UnmarshalAsMap decodes data that is encoded as map format.
// This is the same thing that StructAsArray sets false.
func UnmarshalAsMap(data []byte, v interface{}) error {
	return decoding.Decode(data, v, false)
}

// UnmarshalAsArray decodes data that is encoded as array format.
// This is the same thing that StructAsArray sets true.
func UnmarshalAsArray(data []byte, v interface{}) error {
	return decoding.Decode(data, v, true)
}

// UnmarshalReadAsMap decodes from stream. stream data expects map format.
// This is the same thing that StructAsArray sets false.
func UnmarshalReadAsMap(r io.Reader, v interface{}) error {
	return streamdecoding.Decode(r, v, false)
}

// UnmarshalReadAsArray decodes from stream. stream data expects array format.
// This is the same thing that StructAsArray sets true.
func UnmarshalReadAsArray(r io.Reader, v interface{}) error {
	return streamdecoding.Decode(r, v, true)
}
