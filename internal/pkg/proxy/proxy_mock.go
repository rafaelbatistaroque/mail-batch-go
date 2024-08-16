package proxy

import (
	"errors"
	"net/http"
	"net/http/httptest"
)

// Helpers
type Response struct {
	Result interface{}
	Code   int
	Err    error
}

type handleMock struct {
	CalledCount int
	Spy         handleFunc
	Response    Response
}

func (h *handleMock) WithResponse(response Response) *handleMock {
	h.Response = response

	return h
}

// Builder
func BuildSUT() (http.HandlerFunc, *httptest.ResponseRecorder, *http.Request, *handleMock) {
	req := httptest.NewRequest("FAKE", "/", nil)
	res := httptest.NewRecorder()
	handle := &handleMock{
		CalledCount: 0,
		Response: Response{
			Result: nil,
			Code:   http.StatusBadRequest,
			Err:    errors.New("fake_error"),
		},
	}

	handle.Spy = func(http.ResponseWriter, *http.Request) (interface{}, int, error) {
		handle.CalledCount++
		return handle.Response.Result, handle.Response.Code, handle.Response.Err
	}

	sut := New(handle.Spy)

	return sut, res, req, handle
}
