// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: link_shortener/v1/link_shortener.proto

package linkShortenerV1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on GetShortLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetShortLinkRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetLongLink()) < 1 {
		return GetShortLinkRequestValidationError{
			field:  "LongLink",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetShortLinkRequestValidationError is the validation error returned by
// GetShortLinkRequest.Validate if the designated constraints aren't met.
type GetShortLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShortLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShortLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShortLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShortLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShortLinkRequestValidationError) ErrorName() string {
	return "GetShortLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetShortLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShortLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShortLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShortLinkRequestValidationError{}

// Validate checks the field values on GetShortLinkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetShortLinkResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetShortLinkResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetShortLinkResponseValidationError is the validation error returned by
// GetShortLinkResponse.Validate if the designated constraints aren't met.
type GetShortLinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShortLinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShortLinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShortLinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShortLinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShortLinkResponseValidationError) ErrorName() string {
	return "GetShortLinkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetShortLinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShortLinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShortLinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShortLinkResponseValidationError{}

// Validate checks the field values on GetLongLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetLongLinkRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetShortLink()) < 1 {
		return GetLongLinkRequestValidationError{
			field:  "ShortLink",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetLongLinkRequestValidationError is the validation error returned by
// GetLongLinkRequest.Validate if the designated constraints aren't met.
type GetLongLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLongLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLongLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLongLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLongLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLongLinkRequestValidationError) ErrorName() string {
	return "GetLongLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLongLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLongLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLongLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLongLinkRequestValidationError{}

// Validate checks the field values on GetLongLinkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetLongLinkResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetLongLinkResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetLongLinkResponseValidationError is the validation error returned by
// GetLongLinkResponse.Validate if the designated constraints aren't met.
type GetLongLinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLongLinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLongLinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLongLinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLongLinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLongLinkResponseValidationError) ErrorName() string {
	return "GetLongLinkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLongLinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLongLinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLongLinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLongLinkResponseValidationError{}

// Validate checks the field values on GetShortLinkResponse_Result with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetShortLinkResponse_Result) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ShortLink

	return nil
}

// GetShortLinkResponse_ResultValidationError is the validation error returned
// by GetShortLinkResponse_Result.Validate if the designated constraints
// aren't met.
type GetShortLinkResponse_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShortLinkResponse_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShortLinkResponse_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShortLinkResponse_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShortLinkResponse_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShortLinkResponse_ResultValidationError) ErrorName() string {
	return "GetShortLinkResponse_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e GetShortLinkResponse_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShortLinkResponse_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShortLinkResponse_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShortLinkResponse_ResultValidationError{}

// Validate checks the field values on GetLongLinkResponse_Result with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetLongLinkResponse_Result) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for LongLink

	return nil
}

// GetLongLinkResponse_ResultValidationError is the validation error returned
// by GetLongLinkResponse_Result.Validate if the designated constraints aren't met.
type GetLongLinkResponse_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLongLinkResponse_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLongLinkResponse_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLongLinkResponse_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLongLinkResponse_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLongLinkResponse_ResultValidationError) ErrorName() string {
	return "GetLongLinkResponse_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e GetLongLinkResponse_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLongLinkResponse_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLongLinkResponse_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLongLinkResponse_ResultValidationError{}