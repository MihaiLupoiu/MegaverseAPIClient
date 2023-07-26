package astral

const defaultSoloonEndpoint = "/api/soloons"

type ColorType string

const (
	Blue   ColorType = "blue"
	Red    ColorType = "red"
	Purple ColorType = "purple"
	White  ColorType = "white"
)

type Soloon struct {
	Row         int       `json:"row"`
	Column      int       `json:"column"`
	CandidateId string    `json:"candidateId"`
	Color       ColorType `json:"color"`
}

func (s Soloon) GetEndpoint() string {
	return defaultSoloonEndpoint
}

func (s Soloon) GetPayload() interface{} {
	return s
}
