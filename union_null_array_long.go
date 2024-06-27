// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayLongTypeEnum int

const (
	UnionNullArrayLongTypeEnumArrayLong UnionNullArrayLongTypeEnum = 1
)

type UnionNullArrayLong struct {
	Null      *types.NullVal
	ArrayLong []int64
	UnionType UnionNullArrayLongTypeEnum
}

func writeUnionNullArrayLong(r *UnionNullArrayLong, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayLongTypeEnumArrayLong:
		return writeArrayLong(r.ArrayLong, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayLong")
}

func NewUnionNullArrayLong() *UnionNullArrayLong {
	return &UnionNullArrayLong{}
}

func (r *UnionNullArrayLong) Serialize(w io.Writer) error {
	return writeUnionNullArrayLong(r, w)
}

func DeserializeUnionNullArrayLong(r io.Reader) (*UnionNullArrayLong, error) {
	t := NewUnionNullArrayLong()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullArrayLongFromSchema(r io.Reader, schema string) (*UnionNullArrayLong, error) {
	t := NewUnionNullArrayLong()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullArrayLong) Schema() string {
	return "[\"null\",{\"items\":\"long\",\"type\":\"array\"}]"
}

func (_ *UnionNullArrayLong) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayLong) SetLong(v int64) {

	r.UnionType = (UnionNullArrayLongTypeEnum)(v)
}

func (r *UnionNullArrayLong) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayLong = make([]int64, 0)
		return &ArrayLongWrapper{Target: (&r.ArrayLong)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayLong) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullArrayLong) Finalize()                        {}

func (r *UnionNullArrayLong) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayLongTypeEnumArrayLong:
		return json.Marshal(map[string]interface{}{"array": r.ArrayLong})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayLong")
}

func (r *UnionNullArrayLong) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayLong)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayLong")
}