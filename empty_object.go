// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package dtsavro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EmptyObject int32

const (
	EmptyObjectNULL EmptyObject = 0
	EmptyObjectNONE EmptyObject = 1
)

func (e EmptyObject) String() string {
	switch e {
	case EmptyObjectNULL:
		return "NULL"
	case EmptyObjectNONE:
		return "NONE"
	}
	return "unknown"
}

func writeEmptyObject(r EmptyObject, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewEmptyObjectValue(raw string) (r EmptyObject, err error) {
	switch raw {
	case "NULL":
		return EmptyObjectNULL, nil
	case "NONE":
		return EmptyObjectNONE, nil
	}

	return -1, fmt.Errorf("invalid value for EmptyObject: '%s'", raw)

}

func (b EmptyObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *EmptyObject) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewEmptyObjectValue(stringVal)
	*b = val
	return err
}

type EmptyObjectWrapper struct {
	Target *EmptyObject
}

func (b EmptyObjectWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b EmptyObjectWrapper) SetInt(v int32) {
	*(b.Target) = EmptyObject(v)
}

func (b EmptyObjectWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b EmptyObjectWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b EmptyObjectWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b EmptyObjectWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b EmptyObjectWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b EmptyObjectWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b EmptyObjectWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b EmptyObjectWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b EmptyObjectWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b EmptyObjectWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b EmptyObjectWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b EmptyObjectWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b EmptyObjectWrapper) Finalize() {}
