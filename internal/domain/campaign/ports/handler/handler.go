package handler

import (
	"net/http"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/createcampaign"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/getcampaignbyid"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type CampaignHandler interface {
	CreateCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	SearchCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	GetCampaignById(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type campaignhandler struct {
	createcampaign  createcampaign.UseCase
	searchcampaign  searchcampaign.UseCase
	getcampaignbyid getcampaignbyid.UseCase
}

func MakeCampaignHandlers(
	createcampaign createcampaign.UseCase,
	searchcampaign searchcampaign.UseCase,
	getcampaignbyid getcampaignbyid.UseCase) CampaignHandler {

	return &campaignhandler{
		createcampaign:  createcampaign,
		searchcampaign:  searchcampaign,
		getcampaignbyid: getcampaignbyid,
	}
}
