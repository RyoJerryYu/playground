// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/auth_service.proto

package apiv1

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

// Validate checks the field values on GetAuthStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAuthStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAuthStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAuthStatusRequestMultiError, or nil if none found.
func (m *GetAuthStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAuthStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetAuthStatusRequestMultiError(errors)
	}

	return nil
}

// GetAuthStatusRequestMultiError is an error wrapping multiple validation
// errors returned by GetAuthStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAuthStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAuthStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAuthStatusRequestMultiError) AllErrors() []error { return m }

// GetAuthStatusRequestValidationError is the validation error returned by
// GetAuthStatusRequest.Validate if the designated constraints aren't met.
type GetAuthStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAuthStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAuthStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAuthStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAuthStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAuthStatusRequestValidationError) ErrorName() string {
	return "GetAuthStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAuthStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAuthStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAuthStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAuthStatusRequestValidationError{}

// Validate checks the field values on GetAuthStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAuthStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAuthStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAuthStatusResponseMultiError, or nil if none found.
func (m *GetAuthStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAuthStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAuthStatusResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAuthStatusResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAuthStatusResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetAuthStatusResponseMultiError(errors)
	}

	return nil
}

// GetAuthStatusResponseMultiError is an error wrapping multiple validation
// errors returned by GetAuthStatusResponse.ValidateAll() if the designated
// constraints aren't met.
type GetAuthStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAuthStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAuthStatusResponseMultiError) AllErrors() []error { return m }

// GetAuthStatusResponseValidationError is the validation error returned by
// GetAuthStatusResponse.Validate if the designated constraints aren't met.
type GetAuthStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAuthStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAuthStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAuthStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAuthStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAuthStatusResponseValidationError) ErrorName() string {
	return "GetAuthStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAuthStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAuthStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAuthStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAuthStatusResponseValidationError{}

// Validate checks the field values on SignInRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInRequestMultiError, or
// nil if none found.
func (m *SignInRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for NeverExpire

	if len(errors) > 0 {
		return SignInRequestMultiError(errors)
	}

	return nil
}

// SignInRequestMultiError is an error wrapping multiple validation errors
// returned by SignInRequest.ValidateAll() if the designated constraints
// aren't met.
type SignInRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInRequestMultiError) AllErrors() []error { return m }

// SignInRequestValidationError is the validation error returned by
// SignInRequest.Validate if the designated constraints aren't met.
type SignInRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestValidationError) ErrorName() string { return "SignInRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignInRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestValidationError{}

// Validate checks the field values on SignInWithSSORequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SignInWithSSORequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInWithSSORequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SignInWithSSORequestMultiError, or nil if none found.
func (m *SignInWithSSORequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInWithSSORequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IdpId

	// no validation rules for Code

	// no validation rules for RedirectUri

	if len(errors) > 0 {
		return SignInWithSSORequestMultiError(errors)
	}

	return nil
}

// SignInWithSSORequestMultiError is an error wrapping multiple validation
// errors returned by SignInWithSSORequest.ValidateAll() if the designated
// constraints aren't met.
type SignInWithSSORequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInWithSSORequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInWithSSORequestMultiError) AllErrors() []error { return m }

// SignInWithSSORequestValidationError is the validation error returned by
// SignInWithSSORequest.Validate if the designated constraints aren't met.
type SignInWithSSORequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInWithSSORequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInWithSSORequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInWithSSORequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInWithSSORequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInWithSSORequestValidationError) ErrorName() string {
	return "SignInWithSSORequestValidationError"
}

// Error satisfies the builtin error interface
func (e SignInWithSSORequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInWithSSORequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInWithSSORequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInWithSSORequestValidationError{}

// Validate checks the field values on SignUpRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUpRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUpRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignUpRequestMultiError, or
// nil if none found.
func (m *SignUpRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUpRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return SignUpRequestMultiError(errors)
	}

	return nil
}

// SignUpRequestMultiError is an error wrapping multiple validation errors
// returned by SignUpRequest.ValidateAll() if the designated constraints
// aren't met.
type SignUpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpRequestMultiError) AllErrors() []error { return m }

// SignUpRequestValidationError is the validation error returned by
// SignUpRequest.Validate if the designated constraints aren't met.
type SignUpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpRequestValidationError) ErrorName() string { return "SignUpRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignUpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpRequestValidationError{}

// Validate checks the field values on SignOutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignOutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignOutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignOutRequestMultiError,
// or nil if none found.
func (m *SignOutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignOutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SignOutRequestMultiError(errors)
	}

	return nil
}

// SignOutRequestMultiError is an error wrapping multiple validation errors
// returned by SignOutRequest.ValidateAll() if the designated constraints
// aren't met.
type SignOutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignOutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignOutRequestMultiError) AllErrors() []error { return m }

// SignOutRequestValidationError is the validation error returned by
// SignOutRequest.Validate if the designated constraints aren't met.
type SignOutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignOutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignOutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignOutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignOutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignOutRequestValidationError) ErrorName() string { return "SignOutRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignOutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignOutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignOutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignOutRequestValidationError{}