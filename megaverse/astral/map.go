package astral

type GoalMap struct {
	Goal [][]string `json:"goal"`
}

type CurrentMap struct {
	Map struct {
		ID          string       `json:"_id"`
		Content     [][]*MapCell `json:"content"`
		CandidateID string       `json:"candidateId"`
		Phase       int          `json:"phase"`
		V           int          `json:"__v"`
	} `json:"map"`
}

type MapCell struct {
	Type int `json:"type"`
}
