package entity

import (
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/entity"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
)

var (
	EXPECTED_NAME    = "fake_name"
	EXPECTED_STATUS  = entity.PENDING
	EXPECTED_CONTENT = "fake_content"
	EXPECT_CONTACTS  = []string{"fake1@email.com", "fake2@email.com"}
)

func Test_GivenNewCampaign_WhenNameIsNilOrEmpty_ThenReturnDomainError(t *testing.T) {
	//Arrange
	invalidName := ""

	//Act
	_, err := entity.MakeCampaign(invalidName, EXPECTED_CONTENT, EXPECT_CONTACTS)

	//Assert
	expect.NotNil(t, err)
	expect.StrictEqual(t, err, domainError.Err_NAME_NOT_NIL_OR_EMPTY)
}

func Test_GivenNewCampaign_WhenContentIsNilOrEmpty_ThenReturnDomainError(t *testing.T) {
	//Arrange
	invalidContent := ""

	//Act
	_, err := entity.MakeCampaign(EXPECTED_NAME, invalidContent, EXPECT_CONTACTS)

	//Assert
	expect.NotNil(t, err)
	expect.StrictEqual(t, err, domainError.Err_CONTENT_NOT_NIL_OR_EMPTY)
}

func Test_GivenNewCampaign_WhenContactsIsEmpty_ThenReturnDomainError(t *testing.T) {
	//Arrange
	invalidContacts := []string{}

	//Act
	_, err := entity.MakeCampaign(EXPECTED_NAME, EXPECTED_CONTENT, invalidContacts)

	//Assert
	expect.NotNil(t, err)
	expect.StrictEqual(t, err, domainError.Err_CONTACTS_NOT_EMPTY)
}

func Test_GivenNewCampaign_WhenCreatedOnIsZero_ThenEnsureTestFail(t *testing.T) {
	//Arrange &	Act
	campaign, err := entity.MakeCampaign(EXPECTED_NAME, EXPECTED_CONTENT, EXPECT_CONTACTS)

	//Assert
	expect.NoError(t, err)
	expect.False(t, campaign.GetCreatedOn().IsZero())
}

func Test_GivenCampaign_WhenCreateNewCampaign_ThenReturnInstanceOfCampaign(t *testing.T) {
	//Arrange & Act
	campaign, err := entity.MakeCampaign(EXPECTED_NAME, EXPECTED_CONTENT, EXPECT_CONTACTS)

	//Assert
	expect.NoError(t, err)
	expect.NotNil(t, campaign)
	expect.NotNil(t, campaign.GetId())
	expect.Equal(t, campaign.Name, EXPECTED_NAME)
	expect.Equal(t, campaign.Status, EXPECTED_STATUS)
	expect.Equal(t, campaign.Content, EXPECTED_CONTENT)
	expect.Len(t, campaign.Contacts, len(EXPECT_CONTACTS))
}

func Test_GivenCampaign_WhenGetContactsString_ThenReturnSliceOfContactsAsStrings(t *testing.T) {
	//Arrange & Act
	campaign, err := entity.MakeCampaign(EXPECTED_NAME, EXPECTED_CONTENT, EXPECT_CONTACTS)

	//Assert
	expect.NoError(t, err)
	expect.NotNil(t, campaign)
	expect.NotNil(t, campaign.GetId())
	expect.ContainsAll(t, campaign.GetContactsString(), EXPECT_CONTACTS)
}
