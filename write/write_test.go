package write_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/aidenwallis/go-write/write"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	handle := func(t *testing.T, cb func(http.ResponseWriter)) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "http://test/library", nil)
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			cb(w)
		}).ServeHTTP(w, req)
		return w
	}

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		const status = http.StatusTeapot
		resp := handle(t, func(w http.ResponseWriter) {
			assertNoError(t, write.New(w, status).Empty())
		})

		bs, err := io.ReadAll(resp.Result().Body)
		assertNoError(t, err)
		assertValues(t, 0, len(bs))
		assertValues(t, "0", resp.Result().Header.Get(write.ContentLengthHeader))
		assertValues(t, "", resp.Result().Header.Get(write.ContentTypeHeader))
	})

	t.Run("json", func(t *testing.T) {
		t.Parallel()

		t.Run("error", func(t *testing.T) {
			const (
				failedStatus  = http.StatusUnauthorized
				successStatus = http.StatusTeapot
			)

			resp := handle(t, func(w http.ResponseWriter) {
				err := write.New(w, failedStatus).JSON(func() string {
					return "this is not valid, you cannot marshal funcs!"
				})
				if err == nil {
					t.Error("should fail marshalling invalid json")
				}

				assertNoError(t, write.New(w, successStatus).Empty())
			})

			assertValues(t, successStatus, resp.Result().StatusCode)
		})

		t.Run("success", func(t *testing.T) {
			t.Parallel()

			type bodyData struct {
				Foo string `json:"foo"`
				Bar string `json:"bar"`
			}

			const expected = `{"foo":"foo123","bar":"bar123"}`
			body := bodyData{
				Foo: "foo123",
				Bar: "bar123",
			}

			const status = http.StatusTeapot
			resp := handle(t, func(w http.ResponseWriter) {
				assertNoError(t, write.New(w, status).JSON(body))
			})

			bs, err := io.ReadAll(resp.Result().Body)
			assertNoError(t, err)
			assertValues(t, expected, string(bs))
			assertValues(t, strconv.Itoa(len(expected)), resp.Result().Header.Get(write.ContentLengthHeader))
			assertValues(t, write.ContentTypeJSON, resp.Result().Header.Get(write.ContentTypeHeader))
		})
	})

	t.Run("text", func(t *testing.T) {
		t.Parallel()

		const (
			body   = "this is a text body"
			status = http.StatusTeapot
		)

		resp := handle(t, func(w http.ResponseWriter) {
			assertNoError(t, write.New(w, status).Text(body))
		})

		bs, err := io.ReadAll(resp.Result().Body)
		assertNoError(t, err)
		assertValues(t, body, string(bs))
		assertValues(t, strconv.Itoa(len(body)), resp.Result().Header.Get(write.ContentLengthHeader))
		assertValues(t, write.ContentTypeText, resp.Result().Header.Get(write.ContentTypeHeader))
	})

	t.Run("bytes", func(t *testing.T) {
		t.Parallel()

		const (
			body   = "this is a text body"
			status = http.StatusTeapot
		)

		resp := handle(t, func(w http.ResponseWriter) {
			assertNoError(t, write.New(w, status).Bytes([]byte(body)))
		})

		bs, err := io.ReadAll(resp.Result().Body)
		assertNoError(t, err)
		assertValues(t, body, string(bs))
		assertValues(t, strconv.Itoa(len(body)), resp.Result().Header.Get(write.ContentLengthHeader))
	})

	t.Run("setting status", func(t *testing.T) {
		t.Parallel()

		const (
			firstStatus  = http.StatusNotFound
			secondStatus = http.StatusTeapot
		)

		resp := handle(t, func(w http.ResponseWriter) {
			writer := write.New(w, firstStatus)
			assertValues(t, firstStatus, writer.Status())
			assertValues(t, secondStatus, writer.SetStatus(secondStatus).Status())
			assertValues(t, secondStatus, writer.Status())
			assertNoError(t, writer.Empty())
		})

		assertValues(t, secondStatus, resp.Result().StatusCode)
	})
}

func assertValues[T comparable](t *testing.T, expected T, actual T) {
	if expected != actual {
		t.Error("values do not match, expected", expected, "got", actual)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("unexpected error:", err)
	}
}
