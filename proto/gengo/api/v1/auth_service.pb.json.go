//Code generated by protoc-gen-go-json. DO NOT EDIT.
//versions:
//- protoc-gen-go-json v1.0.7
//- protoc (unknown)
//source: api/v1/auth_service.proto

package apiv1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *GetAuthStatusRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetAuthStatusRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetAuthStatusResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetAuthStatusResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SignInRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SignInRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SignInWithSSORequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SignInWithSSORequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SignUpRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SignUpRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SignOutRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SignOutRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}
