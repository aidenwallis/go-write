// Generated by writegen. DO NOT EDIT.
package write

import "net/http"

// Accepted creates a new writer of status code 202
func Accepted(w http.ResponseWriter) *Writer {
	return New(w, http.StatusAccepted)
}

// AlreadyReported creates a new writer of status code 208
func AlreadyReported(w http.ResponseWriter) *Writer {
	return New(w, http.StatusAlreadyReported)
}

// BadGateway creates a new writer of status code 502
func BadGateway(w http.ResponseWriter) *Writer {
	return New(w, http.StatusBadGateway)
}

// BadRequest creates a new writer of status code 400
func BadRequest(w http.ResponseWriter) *Writer {
	return New(w, http.StatusBadRequest)
}

// Conflict creates a new writer of status code 409
func Conflict(w http.ResponseWriter) *Writer {
	return New(w, http.StatusConflict)
}

// Continue creates a new writer of status code 100
func Continue(w http.ResponseWriter) *Writer {
	return New(w, http.StatusContinue)
}

// Created creates a new writer of status code 201
func Created(w http.ResponseWriter) *Writer {
	return New(w, http.StatusCreated)
}

// EarlyHints creates a new writer of status code 103
func EarlyHints(w http.ResponseWriter) *Writer {
	return New(w, http.StatusEarlyHints)
}

// ExpectationFailed creates a new writer of status code 417
func ExpectationFailed(w http.ResponseWriter) *Writer {
	return New(w, http.StatusExpectationFailed)
}

// FailedDependency creates a new writer of status code 424
func FailedDependency(w http.ResponseWriter) *Writer {
	return New(w, http.StatusFailedDependency)
}

// Forbidden creates a new writer of status code 403
func Forbidden(w http.ResponseWriter) *Writer {
	return New(w, http.StatusForbidden)
}

// Found creates a new writer of status code 302
func Found(w http.ResponseWriter) *Writer {
	return New(w, http.StatusFound)
}

// GatewayTimeout creates a new writer of status code 504
func GatewayTimeout(w http.ResponseWriter) *Writer {
	return New(w, http.StatusGatewayTimeout)
}

// Gone creates a new writer of status code 410
func Gone(w http.ResponseWriter) *Writer {
	return New(w, http.StatusGone)
}

// HTTPVersionNotSupported creates a new writer of status code 505
func HTTPVersionNotSupported(w http.ResponseWriter) *Writer {
	return New(w, http.StatusHTTPVersionNotSupported)
}

// IMUsed creates a new writer of status code 226
func IMUsed(w http.ResponseWriter) *Writer {
	return New(w, http.StatusIMUsed)
}

// InsufficientStorage creates a new writer of status code 507
func InsufficientStorage(w http.ResponseWriter) *Writer {
	return New(w, http.StatusInsufficientStorage)
}

// InternalServerError creates a new writer of status code 500
func InternalServerError(w http.ResponseWriter) *Writer {
	return New(w, http.StatusInternalServerError)
}

// LengthRequired creates a new writer of status code 411
func LengthRequired(w http.ResponseWriter) *Writer {
	return New(w, http.StatusLengthRequired)
}

// Locked creates a new writer of status code 423
func Locked(w http.ResponseWriter) *Writer {
	return New(w, http.StatusLocked)
}

// LoopDetected creates a new writer of status code 508
func LoopDetected(w http.ResponseWriter) *Writer {
	return New(w, http.StatusLoopDetected)
}

// MethodNotAllowed creates a new writer of status code 405
func MethodNotAllowed(w http.ResponseWriter) *Writer {
	return New(w, http.StatusMethodNotAllowed)
}

// MisdirectedRequest creates a new writer of status code 421
func MisdirectedRequest(w http.ResponseWriter) *Writer {
	return New(w, http.StatusMisdirectedRequest)
}

// MovedPermanently creates a new writer of status code 301
func MovedPermanently(w http.ResponseWriter) *Writer {
	return New(w, http.StatusMovedPermanently)
}

// MultiStatus creates a new writer of status code 207
func MultiStatus(w http.ResponseWriter) *Writer {
	return New(w, http.StatusMultiStatus)
}

// MultipleChoices creates a new writer of status code 300
func MultipleChoices(w http.ResponseWriter) *Writer {
	return New(w, http.StatusMultipleChoices)
}

// NetworkAuthenticationRequired creates a new writer of status code 511
func NetworkAuthenticationRequired(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNetworkAuthenticationRequired)
}

// NoContent creates a new writer of status code 204
func NoContent(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNoContent)
}

// NonAuthoritativeInfo creates a new writer of status code 203
func NonAuthoritativeInfo(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNonAuthoritativeInfo)
}

// NotAcceptable creates a new writer of status code 406
func NotAcceptable(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNotAcceptable)
}

// NotExtended creates a new writer of status code 510
func NotExtended(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNotExtended)
}

// NotFound creates a new writer of status code 404
func NotFound(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNotFound)
}

// NotImplemented creates a new writer of status code 501
func NotImplemented(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNotImplemented)
}

// NotModified creates a new writer of status code 304
func NotModified(w http.ResponseWriter) *Writer {
	return New(w, http.StatusNotModified)
}

// OK creates a new writer of status code 200
func OK(w http.ResponseWriter) *Writer {
	return New(w, http.StatusOK)
}

// PartialContent creates a new writer of status code 206
func PartialContent(w http.ResponseWriter) *Writer {
	return New(w, http.StatusPartialContent)
}

// PaymentRequired creates a new writer of status code 402
func PaymentRequired(w http.ResponseWriter) *Writer {
	return New(w, http.StatusPaymentRequired)
}

// PermanentRedirect creates a new writer of status code 308
func PermanentRedirect(w http.ResponseWriter) *Writer {
	return New(w, http.StatusPermanentRedirect)
}

// PreconditionFailed creates a new writer of status code 412
func PreconditionFailed(w http.ResponseWriter) *Writer {
	return New(w, http.StatusPreconditionFailed)
}

// PreconditionRequired creates a new writer of status code 428
func PreconditionRequired(w http.ResponseWriter) *Writer {
	return New(w, http.StatusPreconditionRequired)
}

// Processing creates a new writer of status code 102
func Processing(w http.ResponseWriter) *Writer {
	return New(w, http.StatusProcessing)
}

// ProxyAuthRequired creates a new writer of status code 407
func ProxyAuthRequired(w http.ResponseWriter) *Writer {
	return New(w, http.StatusProxyAuthRequired)
}

// RequestEntityTooLarge creates a new writer of status code 413
func RequestEntityTooLarge(w http.ResponseWriter) *Writer {
	return New(w, http.StatusRequestEntityTooLarge)
}

// RequestHeaderFieldsTooLarge creates a new writer of status code 431
func RequestHeaderFieldsTooLarge(w http.ResponseWriter) *Writer {
	return New(w, http.StatusRequestHeaderFieldsTooLarge)
}

// RequestTimeout creates a new writer of status code 408
func RequestTimeout(w http.ResponseWriter) *Writer {
	return New(w, http.StatusRequestTimeout)
}

// RequestURITooLong creates a new writer of status code 414
func RequestURITooLong(w http.ResponseWriter) *Writer {
	return New(w, http.StatusRequestURITooLong)
}

// RequestedRangeNotSatisfiable creates a new writer of status code 416
func RequestedRangeNotSatisfiable(w http.ResponseWriter) *Writer {
	return New(w, http.StatusRequestedRangeNotSatisfiable)
}

// ResetContent creates a new writer of status code 205
func ResetContent(w http.ResponseWriter) *Writer {
	return New(w, http.StatusResetContent)
}

// SeeOther creates a new writer of status code 303
func SeeOther(w http.ResponseWriter) *Writer {
	return New(w, http.StatusSeeOther)
}

// ServiceUnavailable creates a new writer of status code 503
func ServiceUnavailable(w http.ResponseWriter) *Writer {
	return New(w, http.StatusServiceUnavailable)
}

// SwitchingProtocols creates a new writer of status code 101
func SwitchingProtocols(w http.ResponseWriter) *Writer {
	return New(w, http.StatusSwitchingProtocols)
}

// Teapot creates a new writer of status code 418
func Teapot(w http.ResponseWriter) *Writer {
	return New(w, http.StatusTeapot)
}

// TemporaryRedirect creates a new writer of status code 307
func TemporaryRedirect(w http.ResponseWriter) *Writer {
	return New(w, http.StatusTemporaryRedirect)
}

// TooEarly creates a new writer of status code 425
func TooEarly(w http.ResponseWriter) *Writer {
	return New(w, http.StatusTooEarly)
}

// TooManyRequests creates a new writer of status code 429
func TooManyRequests(w http.ResponseWriter) *Writer {
	return New(w, http.StatusTooManyRequests)
}

// Unauthorized creates a new writer of status code 401
func Unauthorized(w http.ResponseWriter) *Writer {
	return New(w, http.StatusUnauthorized)
}

// UnavailableForLegalReasons creates a new writer of status code 451
func UnavailableForLegalReasons(w http.ResponseWriter) *Writer {
	return New(w, http.StatusUnavailableForLegalReasons)
}

// UnprocessableEntity creates a new writer of status code 422
func UnprocessableEntity(w http.ResponseWriter) *Writer {
	return New(w, http.StatusUnprocessableEntity)
}

// UnsupportedMediaType creates a new writer of status code 415
func UnsupportedMediaType(w http.ResponseWriter) *Writer {
	return New(w, http.StatusUnsupportedMediaType)
}

// UpgradeRequired creates a new writer of status code 426
func UpgradeRequired(w http.ResponseWriter) *Writer {
	return New(w, http.StatusUpgradeRequired)
}

// UseProxy creates a new writer of status code 305
func UseProxy(w http.ResponseWriter) *Writer {
	return New(w, http.StatusUseProxy)
}

// VariantAlsoNegotiates creates a new writer of status code 506
func VariantAlsoNegotiates(w http.ResponseWriter) *Writer {
	return New(w, http.StatusVariantAlsoNegotiates)
}
