// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: weixin-service.proto

package proto

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

// Validate checks the field values on PingReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PingReqMultiError, or nil if none found.
func (m *PingReq) ValidateAll() error {
	return m.validate(true)
}

func (m *PingReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return PingReqMultiError(errors)
	}

	return nil
}

// PingReqMultiError is an error wrapping multiple validation errors returned
// by PingReq.ValidateAll() if the designated constraints aren't met.
type PingReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingReqMultiError) AllErrors() []error { return m }

// PingReqValidationError is the validation error returned by PingReq.Validate
// if the designated constraints aren't met.
type PingReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingReqValidationError) ErrorName() string { return "PingReqValidationError" }

// Error satisfies the builtin error interface
func (e PingReqValidationError) Error() string {
	errmsg := e.Reason()

	if e.cause != nil {
		errmsg = fmt.Sprintf("%s: %s", errmsg, e.cause.Error())
	}

	return errmsg
}

var _ error = PingReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingReqValidationError{}

// Validate checks the field values on PingRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PingRspMultiError, or nil if none found.
func (m *PingRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *PingRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Errno

	// no validation rules for Errmsg

	// no validation rules for Data

	if len(errors) > 0 {
		return PingRspMultiError(errors)
	}

	return nil
}

// PingRspMultiError is an error wrapping multiple validation errors returned
// by PingRsp.ValidateAll() if the designated constraints aren't met.
type PingRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingRspMultiError) AllErrors() []error { return m }

// PingRspValidationError is the validation error returned by PingRsp.Validate
// if the designated constraints aren't met.
type PingRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingRspValidationError) ErrorName() string { return "PingRspValidationError" }

// Error satisfies the builtin error interface
func (e PingRspValidationError) Error() string {
	errmsg := e.Reason()

	if e.cause != nil {
		errmsg = fmt.Sprintf("%s: %s", errmsg, e.cause.Error())
	}

	return errmsg
}

var _ error = PingRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingRspValidationError{}
