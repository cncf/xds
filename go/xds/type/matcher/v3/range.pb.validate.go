// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xds/type/matcher/v3/range.proto

package v3

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

// Validate checks the field values on Int64Range with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Int64Range) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Start

	// no validation rules for End

	return nil
}

// Int64RangeValidationError is the validation error returned by
// Int64Range.Validate if the designated constraints aren't met.
type Int64RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int64RangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int64RangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int64RangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int64RangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int64RangeValidationError) ErrorName() string { return "Int64RangeValidationError" }

// Error satisfies the builtin error interface
func (e Int64RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt64Range.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int64RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int64RangeValidationError{}

// Validate checks the field values on Int64RangeMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Int64RangeMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRanges()) < 1 {
		return Int64RangeMatcherValidationError{
			field:  "Ranges",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Int64RangeMatcherValidationError{
					field:  fmt.Sprintf("Ranges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Int64RangeMatcherValidationError is the validation error returned by
// Int64RangeMatcher.Validate if the designated constraints aren't met.
type Int64RangeMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int64RangeMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int64RangeMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int64RangeMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int64RangeMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int64RangeMatcherValidationError) ErrorName() string {
	return "Int64RangeMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e Int64RangeMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt64RangeMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int64RangeMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int64RangeMatcherValidationError{}

// Validate checks the field values on Int32Range with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Int32Range) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Start

	// no validation rules for End

	return nil
}

// Int32RangeValidationError is the validation error returned by
// Int32Range.Validate if the designated constraints aren't met.
type Int32RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int32RangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int32RangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int32RangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int32RangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int32RangeValidationError) ErrorName() string { return "Int32RangeValidationError" }

// Error satisfies the builtin error interface
func (e Int32RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt32Range.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int32RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int32RangeValidationError{}

// Validate checks the field values on Int32RangeMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Int32RangeMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRanges()) < 1 {
		return Int32RangeMatcherValidationError{
			field:  "Ranges",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Int32RangeMatcherValidationError{
					field:  fmt.Sprintf("Ranges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Int32RangeMatcherValidationError is the validation error returned by
// Int32RangeMatcher.Validate if the designated constraints aren't met.
type Int32RangeMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int32RangeMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int32RangeMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int32RangeMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int32RangeMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int32RangeMatcherValidationError) ErrorName() string {
	return "Int32RangeMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e Int32RangeMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt32RangeMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int32RangeMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int32RangeMatcherValidationError{}

// Validate checks the field values on DoubleRange with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DoubleRange) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Start

	// no validation rules for End

	return nil
}

// DoubleRangeValidationError is the validation error returned by
// DoubleRange.Validate if the designated constraints aren't met.
type DoubleRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleRangeValidationError) ErrorName() string { return "DoubleRangeValidationError" }

// Error satisfies the builtin error interface
func (e DoubleRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleRangeValidationError{}

// Validate checks the field values on DoubleRangeMatcher with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DoubleRangeMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRanges()) < 1 {
		return DoubleRangeMatcherValidationError{
			field:  "Ranges",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DoubleRangeMatcherValidationError{
					field:  fmt.Sprintf("Ranges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DoubleRangeMatcherValidationError is the validation error returned by
// DoubleRangeMatcher.Validate if the designated constraints aren't met.
type DoubleRangeMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleRangeMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleRangeMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleRangeMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleRangeMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleRangeMatcherValidationError) ErrorName() string {
	return "DoubleRangeMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e DoubleRangeMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleRangeMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleRangeMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleRangeMatcherValidationError{}
