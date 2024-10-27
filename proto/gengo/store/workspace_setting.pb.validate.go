// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: store/workspace_setting.proto

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

// Validate checks the field values on WorkspaceSetting with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WorkspaceSetting) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkspaceSetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkspaceSettingMultiError, or nil if none found.
func (m *WorkspaceSetting) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkspaceSetting) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	switch v := m.Value.(type) {
	case *WorkspaceSetting_BasicSetting:
		if v == nil {
			err := WorkspaceSettingValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetBasicSetting()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "BasicSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "BasicSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBasicSetting()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkspaceSettingValidationError{
					field:  "BasicSetting",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *WorkspaceSetting_GeneralSetting:
		if v == nil {
			err := WorkspaceSettingValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetGeneralSetting()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "GeneralSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "GeneralSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGeneralSetting()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkspaceSettingValidationError{
					field:  "GeneralSetting",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *WorkspaceSetting_StorageSetting:
		if v == nil {
			err := WorkspaceSettingValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetStorageSetting()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "StorageSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "StorageSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStorageSetting()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkspaceSettingValidationError{
					field:  "StorageSetting",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *WorkspaceSetting_MemoRelatedSetting:
		if v == nil {
			err := WorkspaceSettingValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetMemoRelatedSetting()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "MemoRelatedSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WorkspaceSettingValidationError{
						field:  "MemoRelatedSetting",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMemoRelatedSetting()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkspaceSettingValidationError{
					field:  "MemoRelatedSetting",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return WorkspaceSettingMultiError(errors)
	}

	return nil
}

// WorkspaceSettingMultiError is an error wrapping multiple validation errors
// returned by WorkspaceSetting.ValidateAll() if the designated constraints
// aren't met.
type WorkspaceSettingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceSettingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceSettingMultiError) AllErrors() []error { return m }

// WorkspaceSettingValidationError is the validation error returned by
// WorkspaceSetting.Validate if the designated constraints aren't met.
type WorkspaceSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceSettingValidationError) ErrorName() string { return "WorkspaceSettingValidationError" }

// Error satisfies the builtin error interface
func (e WorkspaceSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspaceSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceSettingValidationError{}

// Validate checks the field values on WorkspaceBasicSetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WorkspaceBasicSetting) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkspaceBasicSetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkspaceBasicSettingMultiError, or nil if none found.
func (m *WorkspaceBasicSetting) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkspaceBasicSetting) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SecretKey

	// no validation rules for SchemaVersion

	if len(errors) > 0 {
		return WorkspaceBasicSettingMultiError(errors)
	}

	return nil
}

// WorkspaceBasicSettingMultiError is an error wrapping multiple validation
// errors returned by WorkspaceBasicSetting.ValidateAll() if the designated
// constraints aren't met.
type WorkspaceBasicSettingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceBasicSettingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceBasicSettingMultiError) AllErrors() []error { return m }

// WorkspaceBasicSettingValidationError is the validation error returned by
// WorkspaceBasicSetting.Validate if the designated constraints aren't met.
type WorkspaceBasicSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceBasicSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceBasicSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceBasicSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceBasicSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceBasicSettingValidationError) ErrorName() string {
	return "WorkspaceBasicSettingValidationError"
}

// Error satisfies the builtin error interface
func (e WorkspaceBasicSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspaceBasicSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceBasicSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceBasicSettingValidationError{}

// Validate checks the field values on WorkspaceGeneralSetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WorkspaceGeneralSetting) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkspaceGeneralSetting with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkspaceGeneralSettingMultiError, or nil if none found.
func (m *WorkspaceGeneralSetting) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkspaceGeneralSetting) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisallowUserRegistration

	// no validation rules for DisallowPasswordAuth

	// no validation rules for AdditionalScript

	// no validation rules for AdditionalStyle

	if all {
		switch v := interface{}(m.GetCustomProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkspaceGeneralSettingValidationError{
					field:  "CustomProfile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkspaceGeneralSettingValidationError{
					field:  "CustomProfile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCustomProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkspaceGeneralSettingValidationError{
				field:  "CustomProfile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for WeekStartDayOffset

	// no validation rules for DisallowChangeUsername

	// no validation rules for DisallowChangeNickname

	if len(errors) > 0 {
		return WorkspaceGeneralSettingMultiError(errors)
	}

	return nil
}

// WorkspaceGeneralSettingMultiError is an error wrapping multiple validation
// errors returned by WorkspaceGeneralSetting.ValidateAll() if the designated
// constraints aren't met.
type WorkspaceGeneralSettingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceGeneralSettingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceGeneralSettingMultiError) AllErrors() []error { return m }

// WorkspaceGeneralSettingValidationError is the validation error returned by
// WorkspaceGeneralSetting.Validate if the designated constraints aren't met.
type WorkspaceGeneralSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceGeneralSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceGeneralSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceGeneralSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceGeneralSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceGeneralSettingValidationError) ErrorName() string {
	return "WorkspaceGeneralSettingValidationError"
}

// Error satisfies the builtin error interface
func (e WorkspaceGeneralSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspaceGeneralSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceGeneralSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceGeneralSettingValidationError{}

// Validate checks the field values on WorkspaceCustomProfile with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WorkspaceCustomProfile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkspaceCustomProfile with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkspaceCustomProfileMultiError, or nil if none found.
func (m *WorkspaceCustomProfile) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkspaceCustomProfile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for LogoUrl

	// no validation rules for Locale

	// no validation rules for Appearance

	if len(errors) > 0 {
		return WorkspaceCustomProfileMultiError(errors)
	}

	return nil
}

// WorkspaceCustomProfileMultiError is an error wrapping multiple validation
// errors returned by WorkspaceCustomProfile.ValidateAll() if the designated
// constraints aren't met.
type WorkspaceCustomProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceCustomProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceCustomProfileMultiError) AllErrors() []error { return m }

// WorkspaceCustomProfileValidationError is the validation error returned by
// WorkspaceCustomProfile.Validate if the designated constraints aren't met.
type WorkspaceCustomProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceCustomProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceCustomProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceCustomProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceCustomProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceCustomProfileValidationError) ErrorName() string {
	return "WorkspaceCustomProfileValidationError"
}

// Error satisfies the builtin error interface
func (e WorkspaceCustomProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspaceCustomProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceCustomProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceCustomProfileValidationError{}

// Validate checks the field values on WorkspaceStorageSetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WorkspaceStorageSetting) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkspaceStorageSetting with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkspaceStorageSettingMultiError, or nil if none found.
func (m *WorkspaceStorageSetting) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkspaceStorageSetting) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StorageType

	// no validation rules for FilepathTemplate

	// no validation rules for UploadSizeLimitMb

	if all {
		switch v := interface{}(m.GetS3Config()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkspaceStorageSettingValidationError{
					field:  "S3Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkspaceStorageSettingValidationError{
					field:  "S3Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetS3Config()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkspaceStorageSettingValidationError{
				field:  "S3Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WorkspaceStorageSettingMultiError(errors)
	}

	return nil
}

// WorkspaceStorageSettingMultiError is an error wrapping multiple validation
// errors returned by WorkspaceStorageSetting.ValidateAll() if the designated
// constraints aren't met.
type WorkspaceStorageSettingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceStorageSettingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceStorageSettingMultiError) AllErrors() []error { return m }

// WorkspaceStorageSettingValidationError is the validation error returned by
// WorkspaceStorageSetting.Validate if the designated constraints aren't met.
type WorkspaceStorageSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceStorageSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceStorageSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceStorageSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceStorageSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceStorageSettingValidationError) ErrorName() string {
	return "WorkspaceStorageSettingValidationError"
}

// Error satisfies the builtin error interface
func (e WorkspaceStorageSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspaceStorageSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceStorageSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceStorageSettingValidationError{}

// Validate checks the field values on StorageS3Config with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StorageS3Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StorageS3Config with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StorageS3ConfigMultiError, or nil if none found.
func (m *StorageS3Config) ValidateAll() error {
	return m.validate(true)
}

func (m *StorageS3Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessKeyId

	// no validation rules for AccessKeySecret

	// no validation rules for Endpoint

	// no validation rules for Region

	// no validation rules for Bucket

	if len(errors) > 0 {
		return StorageS3ConfigMultiError(errors)
	}

	return nil
}

// StorageS3ConfigMultiError is an error wrapping multiple validation errors
// returned by StorageS3Config.ValidateAll() if the designated constraints
// aren't met.
type StorageS3ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StorageS3ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StorageS3ConfigMultiError) AllErrors() []error { return m }

// StorageS3ConfigValidationError is the validation error returned by
// StorageS3Config.Validate if the designated constraints aren't met.
type StorageS3ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageS3ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageS3ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageS3ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageS3ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageS3ConfigValidationError) ErrorName() string { return "StorageS3ConfigValidationError" }

// Error satisfies the builtin error interface
func (e StorageS3ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageS3Config.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageS3ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageS3ConfigValidationError{}

// Validate checks the field values on WorkspaceMemoRelatedSetting with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WorkspaceMemoRelatedSetting) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkspaceMemoRelatedSetting with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkspaceMemoRelatedSettingMultiError, or nil if none found.
func (m *WorkspaceMemoRelatedSetting) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkspaceMemoRelatedSetting) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisallowPublicVisibility

	// no validation rules for DisplayWithUpdateTime

	// no validation rules for ContentLengthLimit

	// no validation rules for EnableAutoCompact

	// no validation rules for EnableDoubleClickEdit

	// no validation rules for EnableLinkPreview

	// no validation rules for EnableComment

	// no validation rules for EnableLocation

	// no validation rules for DefaultVisibility

	// no validation rules for DisableMarkdownShortcuts

	if len(errors) > 0 {
		return WorkspaceMemoRelatedSettingMultiError(errors)
	}

	return nil
}

// WorkspaceMemoRelatedSettingMultiError is an error wrapping multiple
// validation errors returned by WorkspaceMemoRelatedSetting.ValidateAll() if
// the designated constraints aren't met.
type WorkspaceMemoRelatedSettingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceMemoRelatedSettingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceMemoRelatedSettingMultiError) AllErrors() []error { return m }

// WorkspaceMemoRelatedSettingValidationError is the validation error returned
// by WorkspaceMemoRelatedSetting.Validate if the designated constraints
// aren't met.
type WorkspaceMemoRelatedSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceMemoRelatedSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceMemoRelatedSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceMemoRelatedSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceMemoRelatedSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceMemoRelatedSettingValidationError) ErrorName() string {
	return "WorkspaceMemoRelatedSettingValidationError"
}

// Error satisfies the builtin error interface
func (e WorkspaceMemoRelatedSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspaceMemoRelatedSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceMemoRelatedSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceMemoRelatedSettingValidationError{}