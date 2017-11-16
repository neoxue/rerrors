package rerrors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

/*
	Rerrors
	Rerrors is errors (which is built in error)
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
type Rerrors struct {
	data      interface{}
	errorType string
	cause     error
	code      int
}

const (
	ErrorTypeConfig      = "errorTypeConfig"
	ErrorTypeSys         = "errorTypeSys"
	ErrorTypeRuntime     = "errorTypeRuntime"
	ErrorTypeInternalRes = "errorTypeInternalRes"
	ErrorTypeExternalRes = "errorTypeExternalRes"
	ErrorTypeUnexpected  = "errorTypeUnexpected"
	ErrorTypeOther       = "errorTypeOther"
)

func NewErrors(message string, errorType string, errorCode int, data interface{}) *Rerrors {
	err := errors.New(message)
	return &Rerrors{
		cause:     err,
		errorType: errorType,
		data:      data,
		code:      errorCode,
	}
}

func WrapErrors(err error, msg string, errorType string, errorCode int, data interface{}) *Rerrors {
	if err == nil {
		return nil
	}
	newerr := errors.Wrap(err, msg)
	return &Rerrors{
		cause:     newerr,
		errorType: errorType,
		data:      data,
		code:      errorCode,
	}
}

func (err *Rerrors) Error() string {
	return "errorType:" + err.errorType + " errMsg:" + err.cause.Error()
}

func (err *Rerrors) Cause() error {
	return err.cause
}

func (err *Rerrors) Type() string {
	return err.errorType
}

func (err *Rerrors) Data() interface{} {
	return err.data
}

/*
 * not the same as errors
 * v|+v -> print stack exactly
 * s| print f.Error()
 */
func (f *Rerrors) Format(s fmt.State, verb rune) {
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
func (f *Rerrors) Format(s fmt.State, verb rune) {
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
