// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: store/inbox.proto

package store

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on InboxMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InboxMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InboxMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InboxMessageMultiError, or
// nil if none found.
func (m *InboxMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *InboxMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if m.ActivityId != nil {
		// no validation rules for ActivityId
	}

	if len(errors) > 0 {
		return InboxMessageMultiError(errors)
	}

	return nil
}

// InboxMessageMultiError is an error wrapping multiple validation errors
// returned by InboxMessage.ValidateAll() if the designated constraints aren't met.
type InboxMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InboxMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InboxMessageMultiError) AllErrors() []error { return m }

// InboxMessageValidationError is the validation error returned by
// InboxMessage.Validate if the designated constraints aren't met.
type InboxMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InboxMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InboxMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InboxMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InboxMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InboxMessageValidationError) ErrorName() string { return "InboxMessageValidationError" }

// Error satisfies the builtin error interface
func (e InboxMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInboxMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InboxMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InboxMessageValidationError{}