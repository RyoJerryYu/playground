//Code generated by protoc-gen-enumx. DO NOT EDIT.
//versions:
//- protoc-gen-enumx v1.0.7
//- protoc (unknown)
//source: api/v1/markdown_service.proto

package apiv1

import (
	constraints "golang.org/x/exp/constraints"
)

func (x NodeType) Int() int       { return int(x) }
func (x NodeType) Int64() int64   { return int64(x) }
func (x NodeType) Int32() int32   { return int32(x) }
func (x NodeType) UInt() uint     { return uint(x) }
func (x NodeType) UInt64() uint64 { return uint64(x) }
func (x NodeType) UInt32() uint32 { return uint32(x) }

// NodeTypeFrom returns the NodeType for the given integers, or the zero value if not found.
func NodeTypeFrom[T constraints.Integer | constraints.Float](s T) NodeType {
	return NodeType(s)
}

// NodeTypeFromValid is like NodeTypeFrom, but returns an extra boolean value to check if the conversion is valid.
func NodeTypeFromValid[T constraints.Integer | constraints.Float](s T) (NodeType, bool) {
	_, valid := NodeType_name[int32(s)]
	return NodeType(s), valid
}

// NodeTypeFromStr returns the NodeType for the given string, or the zero value if not found.
func NodeTypeFromStr(s string) NodeType {
	return NodeType(NodeType_value[s])
}

// NodeTypeFromValidStr is like NodeTypeFromStr, but returns an extra boolean value to check if the conversion is valid.
func NodeTypeFromValidStr(s string) (NodeType, bool) {
	v, valid := NodeType_value[s]
	return NodeType(v), valid
}

var _NodeType_all = []NodeType{
	NodeType_NODE_UNSPECIFIED,
	NodeType_LINE_BREAK,
	NodeType_PARAGRAPH,
	NodeType_CODE_BLOCK,
	NodeType_HEADING,
	NodeType_HORIZONTAL_RULE,
	NodeType_BLOCKQUOTE,
	NodeType_LIST,
	NodeType_ORDERED_LIST_ITEM,
	NodeType_UNORDERED_LIST_ITEM,
	NodeType_TASK_LIST_ITEM,
	NodeType_MATH_BLOCK,
	NodeType_TABLE,
	NodeType_EMBEDDED_CONTENT,
	NodeType_TEXT,
	NodeType_BOLD,
	NodeType_ITALIC,
	NodeType_BOLD_ITALIC,
	NodeType_CODE,
	NodeType_IMAGE,
	NodeType_LINK,
	NodeType_AUTO_LINK,
	NodeType_TAG,
	NodeType_STRIKETHROUGH,
	NodeType_ESCAPING_CHARACTER,
	NodeType_MATH,
	NodeType_HIGHLIGHT,
	NodeType_SUBSCRIPT,
	NodeType_SUPERSCRIPT,
	NodeType_REFERENCED_CONTENT,
	NodeType_SPOILER,
	NodeType_HTML_ELEMENT,
}
var _NodeType_allName = []string{
	"NodeType_NODE_UNSPECIFIED",
	"NodeType_LINE_BREAK",
	"NodeType_PARAGRAPH",
	"NodeType_CODE_BLOCK",
	"NodeType_HEADING",
	"NodeType_HORIZONTAL_RULE",
	"NodeType_BLOCKQUOTE",
	"NodeType_LIST",
	"NodeType_ORDERED_LIST_ITEM",
	"NodeType_UNORDERED_LIST_ITEM",
	"NodeType_TASK_LIST_ITEM",
	"NodeType_MATH_BLOCK",
	"NodeType_TABLE",
	"NodeType_EMBEDDED_CONTENT",
	"NodeType_TEXT",
	"NodeType_BOLD",
	"NodeType_ITALIC",
	"NodeType_BOLD_ITALIC",
	"NodeType_CODE",
	"NodeType_IMAGE",
	"NodeType_LINK",
	"NodeType_AUTO_LINK",
	"NodeType_TAG",
	"NodeType_STRIKETHROUGH",
	"NodeType_ESCAPING_CHARACTER",
	"NodeType_MATH",
	"NodeType_HIGHLIGHT",
	"NodeType_SUBSCRIPT",
	"NodeType_SUPERSCRIPT",
	"NodeType_REFERENCED_CONTENT",
	"NodeType_SPOILER",
	"NodeType_HTML_ELEMENT",
}
var _NodeType_allValue = []int32{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	51,
	52,
	53,
	54,
	55,
	56,
	57,
	58,
	59,
	60,
	61,
	62,
	63,
	64,
	65,
	66,
	67,
	68,
}

// NodeTypeAll returns all the values of the NodeType enum.
func NodeTypeAll() []NodeType {
	return _NodeType_all[:]
}

// NodeTypeAllName returns all the names of the NodeType enum.
func NodeTypeAllName() []string {
	return _NodeType_allName[:]
}

// NodeTypeAllValue returns all the values of the NodeType enum.
func NodeTypeAllValue() []int32 {
	return _NodeType_allValue[:]
}

func (x ListNode_Kind) Int() int       { return int(x) }
func (x ListNode_Kind) Int64() int64   { return int64(x) }
func (x ListNode_Kind) Int32() int32   { return int32(x) }
func (x ListNode_Kind) UInt() uint     { return uint(x) }
func (x ListNode_Kind) UInt64() uint64 { return uint64(x) }
func (x ListNode_Kind) UInt32() uint32 { return uint32(x) }

// ListNode_KindFrom returns the ListNode_Kind for the given integers, or the zero value if not found.
func ListNode_KindFrom[T constraints.Integer | constraints.Float](s T) ListNode_Kind {
	return ListNode_Kind(s)
}

// ListNode_KindFromValid is like ListNode_KindFrom, but returns an extra boolean value to check if the conversion is valid.
func ListNode_KindFromValid[T constraints.Integer | constraints.Float](s T) (ListNode_Kind, bool) {
	_, valid := ListNode_Kind_name[int32(s)]
	return ListNode_Kind(s), valid
}

// ListNode_KindFromStr returns the ListNode_Kind for the given string, or the zero value if not found.
func ListNode_KindFromStr(s string) ListNode_Kind {
	return ListNode_Kind(ListNode_Kind_value[s])
}

// ListNode_KindFromValidStr is like ListNode_KindFromStr, but returns an extra boolean value to check if the conversion is valid.
func ListNode_KindFromValidStr(s string) (ListNode_Kind, bool) {
	v, valid := ListNode_Kind_value[s]
	return ListNode_Kind(v), valid
}

var _ListNode_Kind_all = []ListNode_Kind{
	ListNode_KIND_UNSPECIFIED,
	ListNode_ORDERED,
	ListNode_UNORDERED,
	ListNode_DESCRIPTION,
}
var _ListNode_Kind_allName = []string{
	"ListNode_KIND_UNSPECIFIED",
	"ListNode_ORDERED",
	"ListNode_UNORDERED",
	"ListNode_DESCRIPTION",
}
var _ListNode_Kind_allValue = []int32{
	0,
	1,
	2,
	3,
}

// ListNode_KindAll returns all the values of the ListNode_Kind enum.
func ListNode_KindAll() []ListNode_Kind {
	return _ListNode_Kind_all[:]
}

// ListNode_KindAllName returns all the names of the ListNode_Kind enum.
func ListNode_KindAllName() []string {
	return _ListNode_Kind_allName[:]
}

// ListNode_KindAllValue returns all the values of the ListNode_Kind enum.
func ListNode_KindAllValue() []int32 {
	return _ListNode_Kind_allValue[:]
}