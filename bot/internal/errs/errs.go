package errs

import "errors"

var RequestFailed = errors.New("Request failed")
var FailedMarshall = errors.New("Failed marshall")
var FailedResponse = errors.New("Failed response")
var FailedUnmarshall = errors.New("Failed unmashal response")
var FailedRead = errors.New("Failed read")
