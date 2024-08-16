package fixtures

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/ports/handler"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/createcampaign"
)

type createCampaignUseCaseSpy struct {
	CallsCount int

	ResultError   error
	ResultSuccess *createcampaign.Output
}

func (u *createCampaignUseCaseSpy) Execute(input *createcampaign.Input) (*createcampaign.Output, error) {
	u.CallsCount++

	return u.ResultSuccess, u.ResultError
}

func MakeCreateCampaignInput_Valid() string {
	return `{
				"name": "Rafael Batista",
				"Content": "Teste de content2",
				"Contacts": [
					"t@t.com",
					"b@b.com"
				]
			}`
}

func (u *createCampaignUseCaseSpy) WithResultError() {
	u.ResultError = errors.New("usecase_error")
}

func (u *createCampaignUseCaseSpy) WithResultSuccess() {
	u.ResultSuccess = &createcampaign.Output{
		Id: "successId",
	}
}

func BuildCreateCampaignHandlerSUT(request []byte) (handler.CampaignHandler, *createCampaignUseCaseSpy, *httptest.ResponseRecorder, *http.Request) {
	usecaseSpy := &createCampaignUseCaseSpy{
		CallsCount: 0,
	}
	sut := handler.MakeCampaignHandlers(usecaseSpy, nil, nil)

	req := httptest.NewRequest("FAKE", "/", bytes.NewReader(request))
	res := httptest.NewRecorder()

	return sut, usecaseSpy, res, req
}
