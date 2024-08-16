package searchcampaign

type UseCase = searchCampaign

type searchCampaign interface {
	Execute(input *Input) (*Output, error)
}
