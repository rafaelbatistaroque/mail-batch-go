package createcampaign

type UseCase = createCampaign

type createCampaign interface {
	Execute(input *Input) (*Output, error)
}
