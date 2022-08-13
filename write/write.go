package write

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	// ContentLengthHeader is the name of the content-length header key
	ContentLengthHeader = "Content-Length"

	// ContentTypeHeader is the name of the content-type header key
	ContentTypeHeader = "Content-Type"

	// ContentTypeJSON is the content-type value for JSON responses
	ContentTypeJSON = "application/json; charset=utf-8"

	// ContentTypeText is the content-type for Text responses
	ContentTypeText = "text/plain; charset=utf-8"
)

// Writer is the struct that is used for building HTTP responses.
type Writer struct {
	http.ResponseWriter
	status int
}

// New creates a new instance of writer.
func New(rw http.ResponseWriter, status int) *Writer {
	return &Writer{ResponseWriter: rw, status: status}
}

// Status returns the current HTTP status
func (w *Writer) Status() int {
	return w.status
}

// SetStatus mutates the status and returns a new writer
func (w *Writer) SetStatus(status int) *Writer {
	w.status = status
	return w
}

// Bytes writes bytes to the response
func (w *Writer) Bytes(bs []byte) error {
	w.Header().Set(ContentLengthHeader, strconv.Itoa(len(bs)))
	w.WriteHeader(w.status)
	_, err := w.ResponseWriter.Write(bs)
	return err
}

// JSON sets JSON headers and writes the response.
func (w *Writer) JSON(v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	return w.Bytes(bs)
}

// Empty returns an empty HTTP response
func (w *Writer) Empty() error {
	return w.Bytes(nil)
}

// Text is used for writing text responses.
func (w *Writer) Text(v string) error {
	w.Header().Set(ContentTypeHeader, ContentTypeText)
	return w.Bytes([]byte(v))
}
