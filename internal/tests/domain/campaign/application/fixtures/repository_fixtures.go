package fixtures

import (
	"errors"
	"strconv"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
)

type RepositorySpy struct {
	//dependencias and properties
	CalledCount       int
	UpdateCalledCount int

	ParamModel               *model.CampaignModel
	ParamSearchCampaignModel *model.SearchCampaignModel
	ParamGetCampaignById     string

	ResultError             error
	ResultCampaignModelList *[]model.CampaignModel
	ResultCampaignModel     *model.CampaignModel
}

func MakeRepositorySpy( /*dependencias*/ ) *RepositorySpy {
	return &RepositorySpy{
		CalledCount: 0,
		ResultError: nil,
	}
}

func (r *RepositorySpy) WithResultError() *RepositorySpy {
	r.ResultError = errors.New("generic-error")

	return r
}

func (r *RepositorySpy) WithSearchResultSuccess(param *model.SearchCampaignModel) *RepositorySpy {
	var campaign = make([]model.CampaignModel, param.PerPage)

	for i := 0; i < param.PerPage; i++ {
		id := strconv.Itoa(i)
		campaign[i] = model.CampaignModel{
			Id:       "id " + id,
			Name:     "Name fake" + id,
			Content:  "This is a content number " + id,
			Contacts: []string{id + "@teste" + id + ".com"},
		}
	}

	r.ResultCampaignModelList = &campaign

	return r
}

func (r *RepositorySpy) WithResultModel(param *model.CampaignModel) *RepositorySpy {
	r.ResultCampaignModel = param

	return r
}
func (r *RepositorySpy) WithResultListModel(param *[]model.CampaignModel) *RepositorySpy {
	r.ResultCampaignModelList = param

	return r
}

func (r *RepositorySpy) Save(model *model.CampaignModel) error {
	r.CalledCount++
	r.ParamModel = model

	return r.ResultError
}

func (r *RepositorySpy) Search(model *model.SearchCampaignModel) (*[]model.CampaignModel, error) {
	r.CalledCount++
	r.ParamSearchCampaignModel = model

	return r.ResultCampaignModelList, r.ResultError
}

func (r *RepositorySpy) GetById(id string) (*model.CampaignModel, error) {
	r.CalledCount++
	r.ParamGetCampaignById = id

	return r.ResultCampaignModel, r.ResultError
}

func (r *RepositorySpy) Update(model *model.CampaignModel) error {
	r.UpdateCalledCount++
	r.ParamModel = model

	return nil
}
