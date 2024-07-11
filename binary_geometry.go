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

type BinaryGeometry struct {
	Type string `json:"type"`

	Value Bytes `json:"value"`
}

const BinaryGeometryAvroCRC64Fingerprint = "\x92\xc6=\xfac$\x8e\x80"

func NewBinaryGeometry() BinaryGeometry {
	r := BinaryGeometry{}
	return r
}

func DeserializeBinaryGeometry(r io.Reader) (BinaryGeometry, error) {
	t := NewBinaryGeometry()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBinaryGeometryFromSchema(r io.Reader, schema string) (BinaryGeometry, error) {
	t := NewBinaryGeometry()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBinaryGeometry(r BinaryGeometry, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r BinaryGeometry) Serialize(w io.Writer) error {
	return writeBinaryGeometry(r, w)
}

func (r BinaryGeometry) Schema() string {
	return "{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"com.alibaba.dts.formats.avro.BinaryGeometry\",\"type\":\"record\"}"
}

func (r BinaryGeometry) SchemaName() string {
	return "com.alibaba.dts.formats.avro.BinaryGeometry"
}

func (_ BinaryGeometry) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BinaryGeometry) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BinaryGeometry) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BinaryGeometry) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BinaryGeometry) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BinaryGeometry) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BinaryGeometry) SetString(v string)   { panic("Unsupported operation") }
func (_ BinaryGeometry) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BinaryGeometry) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Type}

		return w

	case 1:
		w := BytesWrapper{Target: &r.Value}

		return w

	}
	panic("Unknown field index")
}

func (r *BinaryGeometry) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BinaryGeometry) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BinaryGeometry) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BinaryGeometry) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BinaryGeometry) HintSize(int)                     { panic("Unsupported operation") }
func (_ BinaryGeometry) Finalize()                        {}

func (_ BinaryGeometry) AvroCRC64Fingerprint() []byte {
	return []byte(BinaryGeometryAvroCRC64Fingerprint)
}

func (r BinaryGeometry) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BinaryGeometry) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for value")
	}
	return nil
}
