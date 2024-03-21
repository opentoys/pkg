package encoding

import "github.com/opentoys/pkg/msgpack/def"

func (e *encoder) writeNil() error {
	return e.setByte1Int(def.Nil)
}
