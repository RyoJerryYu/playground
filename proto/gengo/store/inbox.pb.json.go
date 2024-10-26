//Code generated by protoc-gen-go-json. DO NOT EDIT.
//versions:
//- protoc-gen-go-json v1.0.7
//- protoc (unknown)
//source: store/inbox.proto

package store

import (
	bytes "bytes"
	json "encoding/json"
	fmt "fmt"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *InboxMessage) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *InboxMessage) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (enum InboxMessage_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (enum *InboxMessage_Type) UnmarshalJSON(b []byte) error {
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
		*enum = InboxMessage_Type(n)
	case float64:
		*enum = InboxMessage_Type(v)
	case string:
		*enum = InboxMessage_Type(InboxMessage_Type_value[v])
	default:
		return fmt.Errorf("invalid enum value %v", v)
	}
	return nil
}