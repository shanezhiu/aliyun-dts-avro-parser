// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Operation int32

const (
	OperationINSERT     Operation = 0
	OperationUPDATE     Operation = 1
	OperationDELETE     Operation = 2
	OperationDDL        Operation = 3
	OperationBEGIN      Operation = 4
	OperationCOMMIT     Operation = 5
	OperationROLLBACK   Operation = 6
	OperationABORT      Operation = 7
	OperationHEARTBEAT  Operation = 8
	OperationCHECKPOINT Operation = 9
	OperationCOMMAND    Operation = 10
	OperationFILL       Operation = 11
	OperationFINISH     Operation = 12
	OperationCONTROL    Operation = 13
	OperationRDB        Operation = 14
	OperationNOOP       Operation = 15
	OperationINIT       Operation = 16
)

func (e Operation) String() string {
	switch e {
	case OperationINSERT:
		return "INSERT"
	case OperationUPDATE:
		return "UPDATE"
	case OperationDELETE:
		return "DELETE"
	case OperationDDL:
		return "DDL"
	case OperationBEGIN:
		return "BEGIN"
	case OperationCOMMIT:
		return "COMMIT"
	case OperationROLLBACK:
		return "ROLLBACK"
	case OperationABORT:
		return "ABORT"
	case OperationHEARTBEAT:
		return "HEARTBEAT"
	case OperationCHECKPOINT:
		return "CHECKPOINT"
	case OperationCOMMAND:
		return "COMMAND"
	case OperationFILL:
		return "FILL"
	case OperationFINISH:
		return "FINISH"
	case OperationCONTROL:
		return "CONTROL"
	case OperationRDB:
		return "RDB"
	case OperationNOOP:
		return "NOOP"
	case OperationINIT:
		return "INIT"
	}
	return "unknown"
}

func writeOperation(r Operation, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewOperationValue(raw string) (r Operation, err error) {
	switch raw {
	case "INSERT":
		return OperationINSERT, nil
	case "UPDATE":
		return OperationUPDATE, nil
	case "DELETE":
		return OperationDELETE, nil
	case "DDL":
		return OperationDDL, nil
	case "BEGIN":
		return OperationBEGIN, nil
	case "COMMIT":
		return OperationCOMMIT, nil
	case "ROLLBACK":
		return OperationROLLBACK, nil
	case "ABORT":
		return OperationABORT, nil
	case "HEARTBEAT":
		return OperationHEARTBEAT, nil
	case "CHECKPOINT":
		return OperationCHECKPOINT, nil
	case "COMMAND":
		return OperationCOMMAND, nil
	case "FILL":
		return OperationFILL, nil
	case "FINISH":
		return OperationFINISH, nil
	case "CONTROL":
		return OperationCONTROL, nil
	case "RDB":
		return OperationRDB, nil
	case "NOOP":
		return OperationNOOP, nil
	case "INIT":
		return OperationINIT, nil
	}

	return -1, fmt.Errorf("invalid value for Operation: '%s'", raw)

}

func (b Operation) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Operation) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewOperationValue(stringVal)
	*b = val
	return err
}

type OperationWrapper struct {
	Target *Operation
}

func (b OperationWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b OperationWrapper) SetInt(v int32) {
	*(b.Target) = Operation(v)
}

func (b OperationWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b OperationWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b OperationWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b OperationWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b OperationWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b OperationWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b OperationWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b OperationWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b OperationWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b OperationWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b OperationWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b OperationWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b OperationWrapper) Finalize() {}
