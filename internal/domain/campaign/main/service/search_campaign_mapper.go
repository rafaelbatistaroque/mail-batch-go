package service

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type SearchCampaignMapper interface {
	ToSearchCampaignModel(input *searchcampaign.Input) *model.SearchCampaignModel
	ToSearchCampaignOutput(models *[]model.CampaignModel) *searchcampaign.Output
}

type searchCampaignMapper struct{}

func NewSearchCampaignMapper() SearchCampaignMapper {
	return &searchCampaignMapper{}
}

func (m *searchCampaignMapper) ToSearchCampaignModel(input *searchcampaign.Input) *model.SearchCampaignModel {

	return &model.SearchCampaignModel{
		Page:    input.Page,
		PerPage: input.PerPage,
	}
}

func (m *searchCampaignMapper) ToSearchCampaignOutput(param *[]model.CampaignModel) *searchcampaign.Output {
	campaigns := make([]searchcampaign.SearchCampaignItem, len(*param))

	for index, v := range *param {
		campaigns[index] = searchcampaign.SearchCampaignItem{
			Id:       v.Id,
			Name:     v.Name,
			Content:  v.Content,
			Contacts: v.Contacts,
		}
	}

	return &searchcampaign.Output{
		Campaigns: campaigns,
		Total:     len(campaigns),
	}
}
