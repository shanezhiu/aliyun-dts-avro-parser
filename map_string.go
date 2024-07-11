// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package dtsavro

import (
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
	"io"
)

func writeMapString(r map[string]string, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapStringWrapper struct {
	Target *map[string]string
	keys   []string
	values []string
}

func (_ *MapStringWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapStringWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapStringWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapStringWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]string, 0, s)
	}
}

func (r *MapStringWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapStringWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapStringWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v string
	r.values = append(r.values, v)
	return &types.String{Target: &r.values[len(r.values)-1]}
}

func (_ *MapStringWrapper) AppendArray() types.Field { panic("Unsupported operation") }
