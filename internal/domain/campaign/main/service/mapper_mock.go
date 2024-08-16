package service

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type SearchCampaignMapperSpy struct {
	CalledCount int

	// ParamEntity              *entity.Campaign
	// ParamModel               *model.CampaignModel
	ParamSearchCampaignModel *searchcampaign.Input

	// ResultToModel             *model.CampaignModel
	ResultSearchCampaignModel *model.SearchCampaignModel
	// ResultToEntity            *entity.Campaign
}

func (m *SearchCampaignMapperSpy) ToSearchCampaignModel(input *searchcampaign.Input) *model.SearchCampaignModel {
	m.CalledCount++
	m.ParamSearchCampaignModel = input

	return nil
}

func (m *SearchCampaignMapperSpy) ToSearchCampaignOutput(param *[]model.CampaignModel) *searchcampaign.Output {
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

func MakeSearchCampaignMapperSpy() *SearchCampaignMapperSpy {
	return &SearchCampaignMapperSpy{
		CalledCount: 0,
	}
}
