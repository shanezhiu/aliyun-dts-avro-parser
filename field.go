// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package dtsavro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Field struct {
	Name string `json:"name"`

	DataTypeNumber int32 `json:"dataTypeNumber"`
}

const FieldAvroCRC64Fingerprint = "\xa4\x82\xda_\x04c\xb1\x00"

func NewField() Field {
	r := Field{}
	return r
}

func DeserializeField(r io.Reader) (Field, error) {
	t := NewField()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFieldFromSchema(r io.Reader, schema string) (Field, error) {
	t := NewField()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeField(r Field, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.DataTypeNumber, w)
	if err != nil {
		return err
	}
	return err
}

func (r Field) Serialize(w io.Writer) error {
	return writeField(r, w)
}

func (r Field) Schema() string {
	return "{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataTypeNumber\",\"type\":\"int\"}],\"name\":\"com.alibaba.dts.formats.avro.Field\",\"type\":\"record\"}"
}

func (r Field) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Field"
}

func (_ Field) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Field) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Field) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Field) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Field) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Field) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Field) SetString(v string)   { panic("Unsupported operation") }
func (_ Field) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Field) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Name}

		return w

	case 1:
		w := types.Int{Target: &r.DataTypeNumber}

		return w

	}
	panic("Unknown field index")
}

func (r *Field) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Field) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Field) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Field) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Field) HintSize(int)                     { panic("Unsupported operation") }
func (_ Field) Finalize()                        {}

func (_ Field) AvroCRC64Fingerprint() []byte {
	return []byte(FieldAvroCRC64Fingerprint)
}

func (r Field) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["dataTypeNumber"], err = json.Marshal(r.DataTypeNumber)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Field) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["dataTypeNumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DataTypeNumber); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for dataTypeNumber")
	}
	return nil
}
