//Code generated by protoc-gen-enumx. DO NOT EDIT.
//versions:
//- protoc-gen-enumx v1.0.7
//- protoc (unknown)
//source: store/user_setting.proto

package store

import (
	constraints "golang.org/x/exp/constraints"
)

func (x UserSettingKey) Int() int       { return int(x) }
func (x UserSettingKey) Int64() int64   { return int64(x) }
func (x UserSettingKey) Int32() int32   { return int32(x) }
func (x UserSettingKey) UInt() uint     { return uint(x) }
func (x UserSettingKey) UInt64() uint64 { return uint64(x) }
func (x UserSettingKey) UInt32() uint32 { return uint32(x) }

// UserSettingKeyFrom returns the UserSettingKey for the given integers, or the zero value if not found.
func UserSettingKeyFrom[T constraints.Integer | constraints.Float](s T) UserSettingKey {
	return UserSettingKey(s)
}

// UserSettingKeyFromValid is like UserSettingKeyFrom, but returns an extra boolean value to check if the conversion is valid.
func UserSettingKeyFromValid[T constraints.Integer | constraints.Float](s T) (UserSettingKey, bool) {
	_, valid := UserSettingKey_name[int32(s)]
	return UserSettingKey(s), valid
}

// UserSettingKeyFromStr returns the UserSettingKey for the given string, or the zero value if not found.
func UserSettingKeyFromStr(s string) UserSettingKey {
	return UserSettingKey(UserSettingKey_value[s])
}

// UserSettingKeyFromValidStr is like UserSettingKeyFromStr, but returns an extra boolean value to check if the conversion is valid.
func UserSettingKeyFromValidStr(s string) (UserSettingKey, bool) {
	v, valid := UserSettingKey_value[s]
	return UserSettingKey(v), valid
}

var _UserSettingKey_all = []UserSettingKey{
	UserSettingKey_USER_SETTING_KEY_UNSPECIFIED,
	UserSettingKey_ACCESS_TOKENS,
	UserSettingKey_LOCALE,
	UserSettingKey_APPEARANCE,
	UserSettingKey_MEMO_VISIBILITY,
}
var _UserSettingKey_allName = []string{
	"UserSettingKey_USER_SETTING_KEY_UNSPECIFIED",
	"UserSettingKey_ACCESS_TOKENS",
	"UserSettingKey_LOCALE",
	"UserSettingKey_APPEARANCE",
	"UserSettingKey_MEMO_VISIBILITY",
}
var _UserSettingKey_allValue = []int32{
	0,
	1,
	2,
	3,
	4,
}

// UserSettingKeyAll returns all the values of the UserSettingKey enum.
func UserSettingKeyAll() []UserSettingKey {
	return _UserSettingKey_all[:]
}

// UserSettingKeyAllName returns all the names of the UserSettingKey enum.
func UserSettingKeyAllName() []string {
	return _UserSettingKey_allName[:]
}

// UserSettingKeyAllValue returns all the values of the UserSettingKey enum.
func UserSettingKeyAllValue() []int32 {
	return _UserSettingKey_allValue[:]
}
