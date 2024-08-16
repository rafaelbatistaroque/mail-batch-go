package endpoint

import (
	"net/http"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/application"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/main/service"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/ports/handler"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/ports/repository"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/ports/repository/database"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/proxy"
)

type CampaignEndpoints struct {
	CreateCampaign  http.HandlerFunc
	SearchCampaign  http.HandlerFunc
	GetCampaignById http.HandlerFunc
}

func GetCampaignEndpoints() CampaignEndpoints {
	handlers := CampaignHandlersComposer()

	return CampaignEndpoints{
		CreateCampaign:  proxy.New(handlers.CreateCampaign),
		SearchCampaign:  proxy.New(handlers.SearchCampaign),
		GetCampaignById: proxy.New(handlers.GetCampaignById),
	}
}

func CampaignHandlersComposer() handler.CampaignHandler {
	mapper := service.NewMapper()
	repository := repository.NewRepository(database.NewDB())

	createCampaignUseCase := application.NewCreateCampaignUseCase(repository, mapper)
	searchCampaignUseCase := application.NewSearchCampaignUseCase(repository, mapper)
	getCampaignByIdUseCase := application.NewGetCampaignByIdUseCase(repository)

	return handler.MakeCampaignHandlers(
		createCampaignUseCase,
		searchCampaignUseCase,
		getCampaignByIdUseCase)
}
