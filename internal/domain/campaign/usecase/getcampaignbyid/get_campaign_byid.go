package getcampaignbyid

type UseCase = getCampaignById

type getCampaignById interface {
	Execute(id string) (*Output, error)
}
