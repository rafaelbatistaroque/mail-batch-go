package proxy

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
)

func Test_NewHandleProxyCalled_WhenHandleInvoked_EnsureCalledOnce(t *testing.T) {
	// Arrange
	sut, res, req, handleSpy := BuildSUT()

	// Act
	sut.ServeHTTP(res, req)

	// Assert
	expect.Equal(t, handleSpy.CalledCount, 1)
}

func Test_NewHandleProxyCalled_WhenHandleResult_EnsureReturnCorrectHeaderParameters(t *testing.T) {
	// Arrange
	sut, res, req, handleSpy := BuildSUT()

	// Act
	sut.ServeHTTP(res, req)

	// Assert
	expect.Equal(t, res.Header().Get("Content-Type"), "application/json")
	expect.Equal(t, res.Code, handleSpy.Response.Code)
}

func Test_NewHandleProxyCalled_WhenHandleResultWithError_EnsureReturnCorrectObjectError(t *testing.T) {
	// Arrange
	sut, res, req, handleSpy := BuildSUT()

	// Act
	sut.ServeHTTP(res, req)

	// Assert
	var body ResponseBody
	json.NewDecoder(res.Body).Decode(&body)
	expect.False(t, body.IsOk)
	expect.Equal(t, body.Status, http.StatusText(handleSpy.Response.Code))
	expect.Equal(t, body.Data, handleSpy.Response.Err.Error())
}

func Test_NewHandleProxyCalled_WhenHandleResultSuccess_EnsureReturnCorrectObjectSuccess(t *testing.T) {
	// Arrange
	sut, res, req, handleSpy := BuildSUT()
	handleSpy.WithResponse(Response{
		Err:    nil,
		Code:   http.StatusOK,
		Result: "fake_success_result",
	})

	// Act
	sut.ServeHTTP(res, req)

	// Assert
	var body ResponseBody
	json.NewDecoder(res.Body).Decode(&body)
	expect.True(t, body.IsOk)
	expect.Equal(t, body.Status, http.StatusText(handleSpy.Response.Code))
	expect.Equal(t, body.Data, handleSpy.Response.Result)
}
