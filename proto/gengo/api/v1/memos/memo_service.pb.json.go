//Code generated by protoc-gen-go-json. DO NOT EDIT.
//versions:
//- protoc-gen-go-json v1.0.9
//- protoc (unknown)
//source: api/v1/memos/memo_service.proto

package memos

import (
	bytes "bytes"
	json "encoding/json"
	fmt "fmt"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (enum Visibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (enum *Visibility) UnmarshalJSON(b []byte) error {
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
		*enum = Visibility(n)
	case float64:
		*enum = Visibility(v)
	case string:
		*enum = Visibility(Visibility_value[v])
	default:
		return fmt.Errorf("invalid enum value %v", v)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (enum MemoView) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (enum *MemoView) UnmarshalJSON(b []byte) error {
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
		*enum = MemoView(n)
	case float64:
		*enum = MemoView(v)
	case string:
		*enum = MemoView(MemoView_value[v])
	default:
		return fmt.Errorf("invalid enum value %v", v)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (msg *Memo) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Memo) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *MemoProperty) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *MemoProperty) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Location) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Location) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CreateMemoRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CreateMemoRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemosRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemosRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemosResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemosResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetMemoRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetMemoRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetMemoByUidRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetMemoByUidRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UpdateMemoRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UpdateMemoRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteMemoRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteMemoRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *RenameMemoTagRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *RenameMemoTagRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteMemoTagRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteMemoTagRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SetMemoResourcesRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SetMemoResourcesRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoResourcesRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoResourcesRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoResourcesResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoResourcesResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SetMemoRelationsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SetMemoRelationsRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoRelationsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoRelationsRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoRelationsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoRelationsResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CreateMemoCommentRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CreateMemoCommentRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoCommentsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoCommentsRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoCommentsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoCommentsResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoReactionsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoReactionsRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListMemoReactionsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListMemoReactionsResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UpsertMemoReactionRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UpsertMemoReactionRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteMemoReactionRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteMemoReactionRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(b, msg)
}
