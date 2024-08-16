package fixtures

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/ports/handler"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type searchCampaignUseCaseSpy struct {
	CallsCount int

	ResultError   error
	ResultSuccess *searchcampaign.Output
}

func (u *searchCampaignUseCaseSpy) Execute(input *searchcampaign.Input) (*searchcampaign.Output, error) {
	u.CallsCount++

	return u.ResultSuccess, u.ResultError
}

func MakeSearchCampaignInput_Valid() string {
	return `{
				"Page": 1,
				"PerPage": 20
			}`
}

func (u *searchCampaignUseCaseSpy) WithResultError() {
	u.ResultError = errors.New("usecase_error")
}

func (u *searchCampaignUseCaseSpy) WithResultSuccess() {
	u.ResultSuccess = &searchcampaign.Output{}
}

func BuildSearchCampaignHandlerSUT(request []byte) (handler.CampaignHandler, *searchCampaignUseCaseSpy, *httptest.ResponseRecorder, *http.Request) {
	usecaseSpy := &searchCampaignUseCaseSpy{
		CallsCount: 0,
	}
	sut := handler.MakeCampaignHandlers(nil, usecaseSpy, nil)

	req := httptest.NewRequest("FAKE", "/", bytes.NewReader(request))
	res := httptest.NewRecorder()

	return sut, usecaseSpy, res, req
}
