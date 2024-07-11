// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package dtsavro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject(r []*UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper struct {
	Target *[]*UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject
}

func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetLong(v int64) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetString(v string) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetUnionElem(v int64) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) Get(i int) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) Finalize() {
}
func (_ ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) SetDefault(i int) {
	panic("Unsupported operation")
}
func (r ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject, 0, s)
	}
}
func (r ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectWrapper) AppendArray() types.Field {
	var v *UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject
	v = NewUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
