package util

import (
	"weixin/idl/exterr"
)

func GetErrno(err error) int64 {
	if err == nil {
		return exterr.E_ErrOk
	}

	errno := exterr.E_DEFAULT_ERROR

	type coder interface {
		Code() int64
	}

	if code, ok := err.(coder); ok {
		errno = code.Code()
	}

	return errno
}

func GenOuterErrno(err error) int64 {
	if err == nil {
		return exterr.E_ErrOk
	}

	errno := exterr.E_DEFAULT_ERROR

	type coder interface {
		Code() int64
	}

	if code, ok := err.(coder); ok {
		errno = code.Code()
	}

	if errno >= minInternalErr {
		errno = exterr.E_DEFAULT_ERROR
	}

	return errno
}

func GenOuterErrmsg(err error) string {
	if err == nil {
		return exterr.Msg[exterr.E_ErrOk]
	}

	errmsg, _ := exterr.Msg[GenOuterErrno(err)]

	errno := GenOuterErrno(err)
	if errno < minInternalErr && errno != exterr.E_DEFAULT_ERROR {
		errmsg = err.Error()
	}

	return errmsg
}
