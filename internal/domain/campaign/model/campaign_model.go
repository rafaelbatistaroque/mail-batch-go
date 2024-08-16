package model

import (
	"time"

	"github.com/lib/pq"
)

type CampaignModel struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
	Status    string         `json:"status"`
	CreatedOn time.Time      `json:"createdOn"`
	Contacts  pq.StringArray `json:"contacts" gorm:"type:text[]"`
}
