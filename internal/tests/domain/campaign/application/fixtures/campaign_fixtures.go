package fixtures

import (
	"time"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/entity"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rs/xid"
)

func MakeCampaignModelListEmpty() *[]model.CampaignModel {
	emptyList := make([]model.CampaignModel, 0)

	return &emptyList
}

func MakeCampaignModel_Fake() *model.CampaignModel {
	return &model.CampaignModel{
		Id:        xid.New().String(),
		Name:      "Name fake",
		Content:   "This is a content number",
		Status:    "Started",
		CreatedOn: time.Now(),
		Contacts:  []string{xid.New().String() + "@teste.com"},
	}
}

func MakeCampaignEntity_Fake() *entity.Campaign {
	modelCampaign := MakeCampaignModel_Fake()
	return entity.LoadCampaign(
		modelCampaign.Id,
		modelCampaign.Name,
		modelCampaign.Content,
		modelCampaign.Status,
		modelCampaign.CreatedOn,
		modelCampaign.Contacts)
}
