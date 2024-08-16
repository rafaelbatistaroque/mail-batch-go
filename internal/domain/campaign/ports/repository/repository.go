package repository

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/contract"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"gorm.io/gorm"
)

const (
	TABLE_CAMPAIGNS string = "campaigns"
)

type repositoryMySql struct {
	//dependencias
	db *gorm.DB
}

func NewRepository(db *gorm.DB) contract.Repository {
	return &repositoryMySql{
		db: db,
	}
}

func (r *repositoryMySql) Save(model *model.CampaignModel) error {
	tx := r.db.Table(TABLE_CAMPAIGNS).Create(model)

	return tx.Error
}

func (r *repositoryMySql) Search(param *model.SearchCampaignModel) (*[]model.CampaignModel, error) {
	var campaigns []model.CampaignModel
	tx := r.db.Table(TABLE_CAMPAIGNS).Find(&campaigns)

	return &campaigns, tx.Error
}

func (r *repositoryMySql) GetById(id string) (*model.CampaignModel, error) {
	var campaign model.CampaignModel
	tx := r.db.Table(TABLE_CAMPAIGNS).First(&campaign, "id = ?", id)

	return &campaign, tx.Error
}

func (r *repositoryMySql) Update(model *model.CampaignModel) error {
	tx := r.db.Table(TABLE_CAMPAIGNS).Where("id=?", model.Id).Updates(model)

	return tx.Error
}
