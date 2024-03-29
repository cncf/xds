// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xds/type/matcher/v3/domain.proto

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

// Validate checks the field values on ServerNameMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ServerNameMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerNameMatcher with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServerNameMatcherMultiError, or nil if none found.
func (m *ServerNameMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerNameMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDomainMatchers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerNameMatcherValidationError{
						field:  fmt.Sprintf("DomainMatchers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerNameMatcherValidationError{
						field:  fmt.Sprintf("DomainMatchers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerNameMatcherValidationError{
					field:  fmt.Sprintf("DomainMatchers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ServerNameMatcherMultiError(errors)
	}

	return nil
}

// ServerNameMatcherMultiError is an error wrapping multiple validation errors
// returned by ServerNameMatcher.ValidateAll() if the designated constraints
// aren't met.
type ServerNameMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerNameMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerNameMatcherMultiError) AllErrors() []error { return m }

// ServerNameMatcherValidationError is the validation error returned by
// ServerNameMatcher.Validate if the designated constraints aren't met.
type ServerNameMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerNameMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerNameMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerNameMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerNameMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerNameMatcherValidationError) ErrorName() string {
	return "ServerNameMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e ServerNameMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerNameMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerNameMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerNameMatcherValidationError{}

// Validate checks the field values on ServerNameMatcher_DomainMatcher with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServerNameMatcher_DomainMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerNameMatcher_DomainMatcher with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ServerNameMatcher_DomainMatcherMultiError, or nil if none found.
func (m *ServerNameMatcher_DomainMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerNameMatcher_DomainMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetDomains()) < 1 {
		err := ServerNameMatcher_DomainMatcherValidationError{
			field:  "Domains",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOnMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerNameMatcher_DomainMatcherValidationError{
					field:  "OnMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerNameMatcher_DomainMatcherValidationError{
					field:  "OnMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerNameMatcher_DomainMatcherValidationError{
				field:  "OnMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServerNameMatcher_DomainMatcherMultiError(errors)
	}

	return nil
}

// ServerNameMatcher_DomainMatcherMultiError is an error wrapping multiple
// validation errors returned by ServerNameMatcher_DomainMatcher.ValidateAll()
// if the designated constraints aren't met.
type ServerNameMatcher_DomainMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerNameMatcher_DomainMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerNameMatcher_DomainMatcherMultiError) AllErrors() []error { return m }

// ServerNameMatcher_DomainMatcherValidationError is the validation error
// returned by ServerNameMatcher_DomainMatcher.Validate if the designated
// constraints aren't met.
type ServerNameMatcher_DomainMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerNameMatcher_DomainMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerNameMatcher_DomainMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerNameMatcher_DomainMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerNameMatcher_DomainMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerNameMatcher_DomainMatcherValidationError) ErrorName() string {
	return "ServerNameMatcher_DomainMatcherValidationError"
}

// Error satisfies the builtin error interface
func (e ServerNameMatcher_DomainMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerNameMatcher_DomainMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerNameMatcher_DomainMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerNameMatcher_DomainMatcherValidationError{}
