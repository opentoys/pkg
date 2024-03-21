package decoding

import "github.com/opentoys/pkg/msgpack/def"

func (d *decoder) isCodeNil(v byte) bool {
	return def.Nil == v
}
