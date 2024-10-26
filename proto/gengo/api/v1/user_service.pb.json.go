//Code generated by protoc-gen-go-json. DO NOT EDIT.
//versions:
//- protoc-gen-go-json v1.0.7
//- protoc (unknown)
//source: api/v1/user_service.proto

package apiv1

import (
	bytes "bytes"
	json "encoding/json"
	fmt "fmt"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *User) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *User) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (enum User_Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (enum *User_Role) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	v, err := dec.Token()
	if err != nil {
		return err
	}
	switch v := v.(type) {
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return err
		}
		*enum = User_Role(n)
	case float64:
		*enum = User_Role(v)
	case string:
		*enum = User_Role(User_Role_value[v])
	default:
		return fmt.Errorf("invalid enum value %v", v)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (msg *ListUsersRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListUsersRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListUsersResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListUsersResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SearchUsersRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SearchUsersRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SearchUsersResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SearchUsersResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetUserRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetUserRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetUserAvatarBinaryRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetUserAvatarBinaryRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CreateUserRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CreateUserRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UpdateUserRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UpdateUserRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteUserRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteUserRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UserSetting) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UserSetting) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetUserSettingRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetUserSettingRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UpdateUserSettingRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UpdateUserSettingRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UserAccessToken) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UserAccessToken) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListUserAccessTokensRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListUserAccessTokensRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListUserAccessTokensResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListUserAccessTokensResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CreateUserAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CreateUserAccessTokenRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteUserAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteUserAccessTokenRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}
