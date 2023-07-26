package astral

type Coordinates struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type AstralObject interface {
	GetCoordinates() interface{}
	GetEndpoint() string
	GetPayload() interface{}
}
