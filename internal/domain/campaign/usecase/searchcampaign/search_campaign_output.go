package searchcampaign

type Output = searchCampaignOutput

type searchCampaignOutput struct {
	Campaigns []SearchCampaignItem
	Total     int
}

type SearchCampaignItem struct {
	Id       string
	Name     string
	Content  string
	Status   string
	Contacts []string
}
