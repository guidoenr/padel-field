package errs

import "github.com/guidoenr/padel-field/logger"

/*
// general
-1 -> golang error

// register
0 -> NIL
1 -> existing username
2 -> existing email

// login
3 -> username does not exist
4 -> wrong password
5 -> wrong email
*/

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func ThrowError(err error, msg string, code int) RequestError {
	logger.Logerror.Printf(msg+": %v", err)
	return RequestError{
		StatusCode: code,
		Err:        err,
	}
}

func ThrowEmptyError() RequestError {
	return RequestError{
		StatusCode: 0,
		Err:        nil,
	}
}
