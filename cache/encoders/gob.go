package encoders

import (
	"bytes"
	"encoding/gob"
)

type GobEncoder struct{}

func NewGobEncoder() *GobEncoder {
	return &GobEncoder{}
}

func (e *GobEncoder) Encode(object any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(object); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (e *GobEncoder) Decode(rawObject []byte, returnObject any) error {
	buffer := bytes.NewBuffer(rawObject)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(returnObject)
	if err != nil {
		return err
	}
	return nil
}
