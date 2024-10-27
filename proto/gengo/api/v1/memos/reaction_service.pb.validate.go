// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/memos/reaction_service.proto

package memos

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

// Validate checks the field values on Reaction with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Reaction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Reaction with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReactionMultiError, or nil
// if none found.
func (m *Reaction) ValidateAll() error {
	return m.validate(true)
}

func (m *Reaction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Creator

	// no validation rules for ContentId

	// no validation rules for ReactionType

	if len(errors) > 0 {
		return ReactionMultiError(errors)
	}

	return nil
}

// ReactionMultiError is an error wrapping multiple validation errors returned
// by Reaction.ValidateAll() if the designated constraints aren't met.
type ReactionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReactionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReactionMultiError) AllErrors() []error { return m }

// ReactionValidationError is the validation error returned by
// Reaction.Validate if the designated constraints aren't met.
type ReactionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReactionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReactionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReactionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReactionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReactionValidationError) ErrorName() string { return "ReactionValidationError" }

// Error satisfies the builtin error interface
func (e ReactionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReaction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReactionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReactionValidationError{}
