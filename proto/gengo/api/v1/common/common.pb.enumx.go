//Code generated by protoc-gen-enumx. DO NOT EDIT.
//versions:
//- protoc-gen-enumx v1.0.9
//- protoc (unknown)
//source: api/v1/common/common.proto

package common

import (
	constraints "golang.org/x/exp/constraints"
)

func (x RowStatus) Int() int       { return int(x) }
func (x RowStatus) Int64() int64   { return int64(x) }
func (x RowStatus) Int32() int32   { return int32(x) }
func (x RowStatus) UInt() uint     { return uint(x) }
func (x RowStatus) UInt64() uint64 { return uint64(x) }
func (x RowStatus) UInt32() uint32 { return uint32(x) }

// RowStatusFrom returns the RowStatus for the given integers, or the zero value if not found.
func RowStatusFrom[T constraints.Integer | constraints.Float](s T) RowStatus {
	return RowStatus(s)
}

// RowStatusFromValid is like RowStatusFrom, but returns an extra boolean value to check if the conversion is valid.
func RowStatusFromValid[T constraints.Integer | constraints.Float](s T) (RowStatus, bool) {
	_, valid := RowStatus_name[int32(s)]
	return RowStatus(s), valid
}

// RowStatusFromStr returns the RowStatus for the given string, or the zero value if not found.
func RowStatusFromStr(s string) RowStatus {
	return RowStatus(RowStatus_value[s])
}

// RowStatusFromValidStr is like RowStatusFromStr, but returns an extra boolean value to check if the conversion is valid.
func RowStatusFromValidStr(s string) (RowStatus, bool) {
	v, valid := RowStatus_value[s]
	return RowStatus(v), valid
}

var _RowStatus_all = []RowStatus{
	RowStatus_ROW_STATUS_UNSPECIFIED,
	RowStatus_ACTIVE,
	RowStatus_ARCHIVED,
}
var _RowStatus_allName = []string{
	"RowStatus_ROW_STATUS_UNSPECIFIED",
	"RowStatus_ACTIVE",
	"RowStatus_ARCHIVED",
}
var _RowStatus_allValue = []int32{
	0,
	1,
	2,
}

// RowStatusAll returns all the values of the RowStatus enum.
func RowStatusAll() []RowStatus {
	return _RowStatus_all[:]
}

// RowStatusAllName returns all the names of the RowStatus enum.
func RowStatusAllName() []string {
	return _RowStatus_allName[:]
}

// RowStatusAllValue returns all the values of the RowStatus enum.
func RowStatusAllValue() []int32 {
	return _RowStatus_allValue[:]
}
