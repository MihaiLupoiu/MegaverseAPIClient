package astral

type BasicAstralObject struct {
	Row         int    `json:"row"`
	Column      int    `json:"column"`
	CandidateId string `json:"candidateId"`
}

type AstralObject interface {
	GetEndpoint() string
	GetPayload() interface{}
}
