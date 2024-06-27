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

type UnionNullStringArrayFieldTypeEnum int

const (
	UnionNullStringArrayFieldTypeEnumString UnionNullStringArrayFieldTypeEnum = 1

	UnionNullStringArrayFieldTypeEnumArrayField UnionNullStringArrayFieldTypeEnum = 2
)

type UnionNullStringArrayField struct {
	Null       *types.NullVal
	String     string
	ArrayField []Field
	UnionType  UnionNullStringArrayFieldTypeEnum
}

func writeUnionNullStringArrayField(r *UnionNullStringArrayField, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullStringArrayFieldTypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionNullStringArrayFieldTypeEnumArrayField:
		return writeArrayField(r.ArrayField, w)
	}
	return fmt.Errorf("invalid value for *UnionNullStringArrayField")
}

func NewUnionNullStringArrayField() *UnionNullStringArrayField {
	return &UnionNullStringArrayField{}
}

func (r *UnionNullStringArrayField) Serialize(w io.Writer) error {
	return writeUnionNullStringArrayField(r, w)
}

func DeserializeUnionNullStringArrayField(r io.Reader) (*UnionNullStringArrayField, error) {
	t := NewUnionNullStringArrayField()
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

func DeserializeUnionNullStringArrayFieldFromSchema(r io.Reader, schema string) (*UnionNullStringArrayField, error) {
	t := NewUnionNullStringArrayField()
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

func (r *UnionNullStringArrayField) Schema() string {
	return "[\"null\",\"string\",{\"items\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataTypeNumber\",\"type\":\"int\"}],\"name\":\"Field\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullStringArrayField) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullStringArrayField) SetLong(v int64) {

	r.UnionType = (UnionNullStringArrayFieldTypeEnum)(v)
}

func (r *UnionNullStringArrayField) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.String{Target: (&r.String)}
	case 2:
		r.ArrayField = make([]Field, 0)
		return &ArrayFieldWrapper{Target: (&r.ArrayField)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullStringArrayField) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullStringArrayField) Finalize()                        {}

func (r *UnionNullStringArrayField) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullStringArrayFieldTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	case UnionNullStringArrayFieldTypeEnumArrayField:
		return json.Marshal(map[string]interface{}{"array": r.ArrayField})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullStringArrayField")
}

func (r *UnionNullStringArrayField) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.String)
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.ArrayField)
	}
	return fmt.Errorf("invalid value for *UnionNullStringArrayField")
}
