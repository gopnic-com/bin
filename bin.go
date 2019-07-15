package bin

import (
	"bytes"
	"encoding/gob"
)

// Register is registering a new data structure.
func Register(v interface{}) {
	gob.Register(v)
}

// Serializer interface.
type Serializer interface {
	Serialize([]byte, error)
}

// Serialize is encoding an object into bytes binary format.
func Serialize(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	if err := gob.NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unserializer interface.
type Unserializer interface {
	Unserialize(interface{}) error
}

// Unserialize is decoding bytes into object.
func Unserialize(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(v)
}
