// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: udpa/data/orca/v1/orca_load_report.proto

package v1

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

// Validate checks the field values on OrcaLoadReport with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrcaLoadReport) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrcaLoadReport with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrcaLoadReportMultiError,
// or nil if none found.
func (m *OrcaLoadReport) ValidateAll() error {
	return m.validate(true)
}

func (m *OrcaLoadReport) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetCpuUtilization(); val < 0 || val > 1 {
		err := OrcaLoadReportValidationError{
			field:  "CpuUtilization",
			reason: "value must be inside range [0, 1]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetMemUtilization(); val < 0 || val > 1 {
		err := OrcaLoadReportValidationError{
			field:  "MemUtilization",
			reason: "value must be inside range [0, 1]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Rps

	// no validation rules for RequestCost

	{
		sorted_keys := make([]string, len(m.GetUtilization()))
		i := 0
		for key := range m.GetUtilization() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetUtilization()[key]
			_ = val

			// no validation rules for Utilization[key]

			if val := val; val < 0 || val > 1 {
				err := OrcaLoadReportValidationError{
					field:  fmt.Sprintf("Utilization[%v]", key),
					reason: "value must be inside range [0, 1]",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return OrcaLoadReportMultiError(errors)
	}

	return nil
}

// OrcaLoadReportMultiError is an error wrapping multiple validation errors
// returned by OrcaLoadReport.ValidateAll() if the designated constraints
// aren't met.
type OrcaLoadReportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrcaLoadReportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrcaLoadReportMultiError) AllErrors() []error { return m }

// OrcaLoadReportValidationError is the validation error returned by
// OrcaLoadReport.Validate if the designated constraints aren't met.
type OrcaLoadReportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrcaLoadReportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrcaLoadReportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrcaLoadReportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrcaLoadReportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrcaLoadReportValidationError) ErrorName() string { return "OrcaLoadReportValidationError" }

// Error satisfies the builtin error interface
func (e OrcaLoadReportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrcaLoadReport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrcaLoadReportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrcaLoadReportValidationError{}
