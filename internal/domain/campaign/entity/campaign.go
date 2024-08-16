package entity

import (
	"time"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/valueObject"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/helpers/validation"

	"github.com/rs/xid"
)

const (
	PENDING  string = "Pending"
	STARTED  string = "Started"
	CANCELED string = "Canceled"
	DONE     string = "Done"
)

type Campaign struct {
	id         string
	Name       string
	createdOn  time.Time
	lastUpdate time.Time
	Content    string
	Status     string
	Contacts   []valueObject.Contact
}

func (c *Campaign) GetId() string {
	return c.id
}

func (c *Campaign) GetCreatedOn() time.Time {
	return c.createdOn
}

func (c *Campaign) GetLastUpdate() time.Time {
	return c.lastUpdate
}

func (c *Campaign) GetContactsString() []string {
	contacts := make([]string, len(c.Contacts))
	for i, contact := range c.Contacts {
		contacts[i] = contact.Email
	}

	return contacts
}

func MakeCampaign(name, content string, emails []string) (*Campaign, error) {

	err := validate(name, content, emails)
	if err != nil {
		return nil, err
	}

	contacts := toContacts(emails)

	return &Campaign{
		id:        xid.New().String(),
		Name:      name,
		Content:   content,
		createdOn: time.Now(),
		Status:    PENDING,
		Contacts:  contacts,
	}, nil
}

func (c *Campaign) Cancel() {
	c.Status = CANCELED
	c.lastUpdate = time.Now()
}

func LoadCampaign(id string, name string, content string, status string, createdOn time.Time, emails []string) *Campaign {

	contacts := toContacts(emails)

	return &Campaign{
		id:        id,
		Name:      name,
		Content:   content,
		Status:    status,
		createdOn: createdOn,
		Contacts:  contacts,
	}
}

func toContacts(emails []string) []valueObject.Contact {
	contacts := make([]valueObject.Contact, len(emails))
	for index, value := range emails {
		contacts[index].Email = value
	}

	return contacts
}

func validate(name string, content string, emails []string) error {
	if validation.IsNilOrEmpty(name) {
		return domainError.Err_NAME_NOT_NIL_OR_EMPTY
	}

	if validation.IsNilOrEmpty(content) {
		return domainError.Err_CONTENT_NOT_NIL_OR_EMPTY
	}

	if validation.IsNilOrEmpty(emails) {
		return domainError.Err_CONTACTS_NOT_EMPTY
	}

	return nil
}
