package createcampaign

type Output = createCampaignOutput

type createCampaignOutput struct {
	Id       string
	Name     string
	Content  string
	Contacts []string
}
