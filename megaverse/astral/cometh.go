package astral

const defaultComethEndpoint = "/api/comeths"

type DirectionType string

const (
	Up    ColorType = "up"
	Down  ColorType = "down"
	Left  ColorType = "left"
	Right ColorType = "right"
)

type Cometh struct {
	Row         int           `json:"row"`
	Column      int           `json:"column"`
	CandidateId string        `json:"candidateId"`
	Direction   DirectionType `json:"direction"`
}

func (c Cometh) GetEndpoint() string {
	return defaultComethEndpoint
}

func (c Cometh) GetPayload() interface{} {
	return c
}
