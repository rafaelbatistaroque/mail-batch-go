package fixtures

import (
	"net/http"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/entity"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/cancelcampaign"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type MapperSpy struct {
	ToEntityCalledCount              int
	ToCampaignModelCalledCount       int
	ToSearchCampaignModelCalledCount int

	ParamEntity              *entity.Campaign
	ParamModel               *model.CampaignModel
	ParamSearchCampaignModel *searchcampaign.Input
	ParamCancelCampaignModel *cancelcampaign.Input

	ResultMapToModel               *model.CampaignModel
	ResultMapToSearchCampaignModel *model.SearchCampaignModel
	ResultMapToEntity              *entity.Campaign
}

func MakeMapperSpy() *MapperSpy {
	return &MapperSpy{
		ToEntityCalledCount:              0,
		ToCampaignModelCalledCount:       0,
		ToSearchCampaignModelCalledCount: 0,
	}
}

func (m *MapperSpy) WithResultModel(campaignModel *model.CampaignModel) *MapperSpy {
	m.ResultMapToModel = &model.CampaignModel{}
	if campaignModel != nil {
		m.ResultMapToModel = campaignModel
	}

	return m
}

func (m *MapperSpy) WithResultEntity(entity *entity.Campaign) *MapperSpy {
	m.ResultMapToEntity = MakeCampaignEntity_Fake()
	if entity != nil {
		m.ResultMapToEntity = entity
	}

	return m
}

func (m *MapperSpy) ToCampaignModel(entity *entity.Campaign) *model.CampaignModel {
	m.ToCampaignModelCalledCount++
	m.ParamEntity = entity

	if m.ResultMapToModel != nil {
		return m.ResultMapToModel
	}

	return &model.CampaignModel{}
}

func (m *MapperSpy) ToSearchCampaignModel(input *searchcampaign.Input) *model.SearchCampaignModel {
	m.ToSearchCampaignModelCalledCount++
	m.ParamSearchCampaignModel = input

	return nil
}

func (m *MapperSpy) ToEntity(model *model.CampaignModel) *entity.Campaign {
	m.ToEntityCalledCount++
	m.ParamModel = model

	return m.ResultMapToEntity
}

func (m *MapperSpy) ToJsonWriter(w http.ResponseWriter, status int, output map[string]interface{}) {

}

func (m *MapperSpy) ToSearchCampaignOutput(param *[]model.CampaignModel) *searchcampaign.Output {
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
