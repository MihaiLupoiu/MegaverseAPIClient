package astral

const baseURL = ""

type GoalMap struct {
	Goal [][]string `json:"goal"`
}

type CurrentMap struct {
	ID          string      `json:"_id"`
	Content     [][]*string `json:"content"`
	CandidateID string      `json:"candidateId"`
	Phase       int         `json:"phase"`
	V           int         `json:"__v"`
}

type Coordinates struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type AstralObject interface {
	GetCoordinates() Coordinates
	GetEndpoint(candidateID string) string
	GetPayload() ([]byte, error)
}
