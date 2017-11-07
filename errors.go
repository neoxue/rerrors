package rerrors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

/*
	WithTypeErrors
	WithTypeErrors is errors (which is built in error)
	error types:
	use string not int, clearly
	1, config 	error
	2, sys		error
	3, runtime 	error
	4, internal resource error
	4, external resource error
	6, unexpected 		 error
	7, other 		 	 error

*/
type WithTypeErrors struct {
	data      interface{}
	errorType string
	cause     error
	code      int
}

const ErrorTypeConfig = "errorTypeConfig"
const ErrorTypeSys = "errorTypeSys"
const ErrorTypeRuntime = "errorTypeRuntime"
const ErrorTypeInternalRes = "errorTypeInternalRes"
const ErrorTypeExternalRes = "errorTypeExternalRes"
const ErrorTypeUnexpected = "errorTypeUnexpected"
const ErrorTypeOther = "errorTypeOther"

func NewErrors(message string, errorType string, errorCode int, data interface{}) *WithTypeErrors {
	err := errors.New(message)
	return &WithTypeErrors{
		cause:     err,
		errorType: errorType,
		data:      data,
		code:      errorCode,
	}
}

func WrapErrors(err error, msg string, errorType string, errorCode int, data interface{}) *WithTypeErrors {
	if err == nil {
		return nil
	}
	newerr := errors.Wrap(err, msg)
	return &WithTypeErrors{
		cause:     newerr,
		errorType: errorType,
		data:      data,
		code:      errorCode,
	}
}

func (err *WithTypeErrors) Error() string {
	return "errorType:" + err.errorType + " errMsg:" + err.cause.Error()
}

func (err *WithTypeErrors) Cause() error {
	return err.cause
}

func (err *WithTypeErrors) Type() string {
	return err.errorType
}

func (err *WithTypeErrors) Data() interface{} {
	return err.data
}

/*
 * not the same as errors
 * v|+v -> print stack exactly
 * s| print f.Error()
 */
func (f *WithTypeErrors) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		fmt.Fprintf(s, "errorType:%v \n", f.errorType)
		fmt.Fprintf(s, "errorCode:%v \nmsg:", f.code)
		fmt.Fprintf(s, "%+v", f.Cause())
		return
	case 's':
		io.WriteString(s, f.Error())
	case 'q':
		fmt.Fprintf(s, "%q", f.Error())
	}
}

/*
func (f *WithTypeErrors) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "errorType:%v \n", f.errorType)
			fmt.Fprintf(s, "errorCode:%v \nmsg:", f.code)
			fmt.Fprintf(s, "%+v", f.Cause())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, f.Error())
	case 'q':
		fmt.Fprintf(s, "%q", f.Error())
	}
}
*/
