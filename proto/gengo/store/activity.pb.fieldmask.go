//Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
//versions:
//- protoc-gen-go-fieldmask v1.0.9
//- protoc (unknown)
//source: store/activity.proto

package store

// IActivityMemoCommentPayloadFieldPath is the interface for the field path of ActivityMemoCommentPayload
type IActivityMemoCommentPayloadFieldPath interface {
	String() string
	MemoId() string
	RelatedMemoId() string
}

// activityMemoCommentPayloadFieldPath is the implementation for the field path of ActivityMemoCommentPayload
type activityMemoCommentPayloadFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewActivityMemoCommentPayloadFieldPath creates a new activityMemoCommentPayloadFieldPath
func NewActivityMemoCommentPayloadFieldPath(fieldPath string) IActivityMemoCommentPayloadFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return activityMemoCommentPayloadFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x activityMemoCommentPayloadFieldPath) String() string { return x.fieldPath }

func (x activityMemoCommentPayloadFieldPath) MemoId() string { return x.prefix + "memo_id" }
func (x activityMemoCommentPayloadFieldPath) RelatedMemoId() string {
	return x.prefix + "related_memo_id"
}

// FieldPath returns the field path for ActivityMemoCommentPayload
func (x *ActivityMemoCommentPayload) FieldPath() IActivityMemoCommentPayloadFieldPath {
	return NewActivityMemoCommentPayloadFieldPath("")
}

// IActivityVersionUpdatePayloadFieldPath is the interface for the field path of ActivityVersionUpdatePayload
type IActivityVersionUpdatePayloadFieldPath interface {
	String() string
	Version() string
}

// activityVersionUpdatePayloadFieldPath is the implementation for the field path of ActivityVersionUpdatePayload
type activityVersionUpdatePayloadFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewActivityVersionUpdatePayloadFieldPath creates a new activityVersionUpdatePayloadFieldPath
func NewActivityVersionUpdatePayloadFieldPath(fieldPath string) IActivityVersionUpdatePayloadFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return activityVersionUpdatePayloadFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x activityVersionUpdatePayloadFieldPath) String() string { return x.fieldPath }

func (x activityVersionUpdatePayloadFieldPath) Version() string { return x.prefix + "version" }

// FieldPath returns the field path for ActivityVersionUpdatePayload
func (x *ActivityVersionUpdatePayload) FieldPath() IActivityVersionUpdatePayloadFieldPath {
	return NewActivityVersionUpdatePayloadFieldPath("")
}

// IActivityPayloadFieldPath is the interface for the field path of ActivityPayload
type IActivityPayloadFieldPath interface {
	String() string
	MemoComment() IActivityMemoCommentPayloadFieldPath
	VersionUpdate() IActivityVersionUpdatePayloadFieldPath
}

// activityPayloadFieldPath is the implementation for the field path of ActivityPayload
type activityPayloadFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewActivityPayloadFieldPath creates a new activityPayloadFieldPath
func NewActivityPayloadFieldPath(fieldPath string) IActivityPayloadFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return activityPayloadFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x activityPayloadFieldPath) String() string { return x.fieldPath }

func (x activityPayloadFieldPath) MemoComment() IActivityMemoCommentPayloadFieldPath {
	return NewActivityMemoCommentPayloadFieldPath(x.prefix + "memo_comment")
}
func (x activityPayloadFieldPath) VersionUpdate() IActivityVersionUpdatePayloadFieldPath {
	return NewActivityVersionUpdatePayloadFieldPath(x.prefix + "version_update")
}

// FieldPath returns the field path for ActivityPayload
func (x *ActivityPayload) FieldPath() IActivityPayloadFieldPath {
	return NewActivityPayloadFieldPath("")
}