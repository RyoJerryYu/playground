// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: store/idp.proto

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

// Validate checks the field values on IdentityProvider with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IdentityProvider) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdentityProvider with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdentityProviderMultiError, or nil if none found.
func (m *IdentityProvider) ValidateAll() error {
	return m.validate(true)
}

func (m *IdentityProvider) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for IdentifierFilter

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IdentityProviderValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IdentityProviderValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IdentityProviderValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IdentityProviderMultiError(errors)
	}

	return nil
}

// IdentityProviderMultiError is an error wrapping multiple validation errors
// returned by IdentityProvider.ValidateAll() if the designated constraints
// aren't met.
type IdentityProviderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentityProviderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentityProviderMultiError) AllErrors() []error { return m }

// IdentityProviderValidationError is the validation error returned by
// IdentityProvider.Validate if the designated constraints aren't met.
type IdentityProviderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentityProviderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentityProviderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentityProviderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentityProviderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentityProviderValidationError) ErrorName() string { return "IdentityProviderValidationError" }

// Error satisfies the builtin error interface
func (e IdentityProviderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentityProvider.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentityProviderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentityProviderValidationError{}

// Validate checks the field values on IdentityProviderConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IdentityProviderConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdentityProviderConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdentityProviderConfigMultiError, or nil if none found.
func (m *IdentityProviderConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *IdentityProviderConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Config.(type) {
	case *IdentityProviderConfig_Oauth2Config:
		if v == nil {
			err := IdentityProviderConfigValidationError{
				field:  "Config",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetOauth2Config()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IdentityProviderConfigValidationError{
						field:  "Oauth2Config",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IdentityProviderConfigValidationError{
						field:  "Oauth2Config",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOauth2Config()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IdentityProviderConfigValidationError{
					field:  "Oauth2Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return IdentityProviderConfigMultiError(errors)
	}

	return nil
}

// IdentityProviderConfigMultiError is an error wrapping multiple validation
// errors returned by IdentityProviderConfig.ValidateAll() if the designated
// constraints aren't met.
type IdentityProviderConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentityProviderConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentityProviderConfigMultiError) AllErrors() []error { return m }

// IdentityProviderConfigValidationError is the validation error returned by
// IdentityProviderConfig.Validate if the designated constraints aren't met.
type IdentityProviderConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentityProviderConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentityProviderConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentityProviderConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentityProviderConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentityProviderConfigValidationError) ErrorName() string {
	return "IdentityProviderConfigValidationError"
}

// Error satisfies the builtin error interface
func (e IdentityProviderConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentityProviderConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentityProviderConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentityProviderConfigValidationError{}

// Validate checks the field values on FieldMapping with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FieldMapping) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FieldMapping with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FieldMappingMultiError, or
// nil if none found.
func (m *FieldMapping) ValidateAll() error {
	return m.validate(true)
}

func (m *FieldMapping) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identifier

	// no validation rules for DisplayName

	// no validation rules for Email

	if len(errors) > 0 {
		return FieldMappingMultiError(errors)
	}

	return nil
}

// FieldMappingMultiError is an error wrapping multiple validation errors
// returned by FieldMapping.ValidateAll() if the designated constraints aren't met.
type FieldMappingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldMappingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldMappingMultiError) AllErrors() []error { return m }

// FieldMappingValidationError is the validation error returned by
// FieldMapping.Validate if the designated constraints aren't met.
type FieldMappingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldMappingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldMappingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldMappingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldMappingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldMappingValidationError) ErrorName() string { return "FieldMappingValidationError" }

// Error satisfies the builtin error interface
func (e FieldMappingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldMapping.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldMappingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldMappingValidationError{}

// Validate checks the field values on OAuth2Config with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OAuth2Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2Config with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OAuth2ConfigMultiError, or
// nil if none found.
func (m *OAuth2Config) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	// no validation rules for ClientSecret

	// no validation rules for AuthUrl

	// no validation rules for TokenUrl

	// no validation rules for UserInfoUrl

	if all {
		switch v := interface{}(m.GetFieldMapping()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "FieldMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "FieldMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFieldMapping()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "FieldMapping",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OAuth2ConfigMultiError(errors)
	}

	return nil
}

// OAuth2ConfigMultiError is an error wrapping multiple validation errors
// returned by OAuth2Config.ValidateAll() if the designated constraints aren't met.
type OAuth2ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2ConfigMultiError) AllErrors() []error { return m }

// OAuth2ConfigValidationError is the validation error returned by
// OAuth2Config.Validate if the designated constraints aren't met.
type OAuth2ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ConfigValidationError) ErrorName() string { return "OAuth2ConfigValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Config.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ConfigValidationError{}
