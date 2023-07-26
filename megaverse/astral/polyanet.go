package astral

const defaultPolyanetEndpoint = "/api/polyanets"

type Polyanet struct {
	Row         int    `json:"row"`
	Column      int    `json:"column"`
	CandidateId string `json:"candidateId"`
}

func (p Polyanet) GetEndpoint() string {
	return defaultPolyanetEndpoint
}

func (p Polyanet) GetPayload() interface{} {
	return p
}
