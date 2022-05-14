// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xds/type/matcher/v3/cel.proto

package v3

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

// Validate checks the field values on CelMatcher with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CelMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CelMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CelMatcherMultiError, or
// nil if none found.
func (m *CelMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *CelMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetExprMatch() == nil {
		err := CelMatcherValidationError{
			field:  "ExprMatch",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetExprMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CelMatcherValidationError{
					field:  "ExprMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CelMatcherValidationError{
					field:  "ExprMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExprMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CelMatcherValidationError{
				field:  "ExprMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CelMatcherMultiError(errors)
	}

	return nil
}

// CelMatcherMultiError is an error wrapping multiple validation errors
// returned by CelMatcher.ValidateAll() if the designated constraints aren't met.
type CelMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CelMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CelMatcherMultiError) AllErrors() []error { return m }

// CelMatcherValidationError is the validation error returned by
// CelMatcher.Validate if the designated constraints aren't met.
type CelMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CelMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CelMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CelMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CelMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CelMatcherValidationError) ErrorName() string { return "CelMatcherValidationError" }

// Error satisfies the builtin error interface
func (e CelMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCelMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CelMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CelMatcherValidationError{}