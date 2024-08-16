package getcampaignbyid

type Output = getCampaignByIdOutput

type getCampaignByIdOutput struct {
	Id       string
	Name     string
	Content  string
	Status   string
	Contacts []string
}
