package service

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/entity"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type Mapper interface {
	ToCampaignModel(campaign *entity.Campaign) *model.CampaignModel
	ToEntity(campaignModel *model.CampaignModel) *entity.Campaign
	ToSearchCampaignModel(input *searchcampaign.Input) *model.SearchCampaignModel
	ToSearchCampaignOutput(models *[]model.CampaignModel) *searchcampaign.Output
}

type mapper struct{}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToCampaignModel(entity *entity.Campaign) *model.CampaignModel {
	contacts := make([]string, len(entity.Contacts))

	for i, contact := range entity.Contacts {
		contacts[i] = contact.Email
	}

	return &model.CampaignModel{
		Id:        entity.GetId(),
		Name:      entity.Name,
		Content:   entity.Content,
		Status:    entity.Status,
		CreatedOn: entity.GetCreatedOn(),
		Contacts:  contacts,
	}
}

func (m *mapper) ToSearchCampaignModel(input *searchcampaign.Input) *model.SearchCampaignModel {

	return &model.SearchCampaignModel{
		Page:    input.Page,
		PerPage: input.PerPage,
	}
}

func (m *mapper) ToEntity(model *model.CampaignModel) *entity.Campaign {

	return entity.LoadCampaign(
		model.Id,
		model.Name, model.Content,
		model.Status,
		model.CreatedOn,
		model.Contacts)
}

func (m *mapper) ToSearchCampaignOutput(param *[]model.CampaignModel) *searchcampaign.Output {
	campaigns := make([]searchcampaign.SearchCampaignItem, len(*param))

	for index, v := range *param {
		campaigns[index] = searchcampaign.SearchCampaignItem{
			Id:       v.Id,
			Name:     v.Name,
			Content:  v.Content,
			Status:   v.Status,
			Contacts: v.Contacts,
		}
	}

	return &searchcampaign.Output{
		Campaigns: campaigns,
		Total:     len(campaigns),
	}
}
