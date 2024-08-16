package contract

import "github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"

type Repository interface {
	Save(model *model.CampaignModel) error
	Search(model *model.SearchCampaignModel) (*[]model.CampaignModel, error)
	GetById(id string) (*model.CampaignModel, error)
	Update(model *model.CampaignModel) error
}
