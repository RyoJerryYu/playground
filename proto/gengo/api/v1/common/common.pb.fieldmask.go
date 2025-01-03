//Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
//versions:
//- protoc-gen-go-fieldmask v1.0.9
//- protoc (unknown)
//source: api/v1/common/common.proto

package common

// IPageTokenFieldPath is the interface for the field path of PageToken
type IPageTokenFieldPath interface {
	String() string
	Limit() string
	Offset() string
}

// pageTokenFieldPath is the implementation for the field path of PageToken
type pageTokenFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewPageTokenFieldPath creates a new pageTokenFieldPath
func NewPageTokenFieldPath(fieldPath string) IPageTokenFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return pageTokenFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x pageTokenFieldPath) String() string { return x.fieldPath }

func (x pageTokenFieldPath) Limit() string  { return x.prefix + "limit" }
func (x pageTokenFieldPath) Offset() string { return x.prefix + "offset" }

// FieldPath returns the field path for PageToken
func (x *PageToken) FieldPath() IPageTokenFieldPath {
	return NewPageTokenFieldPath("")
}
