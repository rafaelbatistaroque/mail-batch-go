package cancelcampaign

type UseCase = cancelCampaign

type cancelCampaign interface {
	Execute(input *Input) (*Output, error)
}
