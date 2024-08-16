package fixtures

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/ports/handler"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/getcampaignbyid"
)

type getCampaignByIdUseCaseSpy struct {
	CallsCount int

	ResultError   error
	ResultSuccess *getcampaignbyid.Output
}

func (u *getCampaignByIdUseCaseSpy) WithResultError() {
	u.ResultError = errors.New("usecase_error")
}

func (u *getCampaignByIdUseCaseSpy) WithResultSuccess(id string) {
	u.ResultSuccess = &getcampaignbyid.Output{
		Id: id,
	}
}

func (u *getCampaignByIdUseCaseSpy) Execute(id string) (*getcampaignbyid.Output, error) {
	u.CallsCount++

	return u.ResultSuccess, u.ResultError
}

func BuildGetCampaignByIdHandlerSUT(id string) (handler.CampaignHandler, *getCampaignByIdUseCaseSpy, *httptest.ResponseRecorder, *http.Request) {
	usecaseSpy := &getCampaignByIdUseCaseSpy{
		CallsCount: 0,
	}

	sut := handler.MakeCampaignHandlers(nil, nil, usecaseSpy)

	url := fmt.Sprintf("/%s", id)
	req := httptest.NewRequest("FAKE", url, nil)
	res := httptest.NewRecorder()

	return sut, usecaseSpy, res, req
}
