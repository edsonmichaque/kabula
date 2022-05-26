package spec

type Mode string

const (
	SingleAnswer    = Mode("SINGLE_ANSWER")
	MultipleAnswers = Mode("MULTIPLE_ANSWERS")
)

type Question struct {
	Src   string         `json:"Ref,omitempty"`
	Value *QuestionValue `json:"Value,omitempty"`
}

type QuestionValue struct {
	Question *Question `json:"Question,omitempty"`
	Options  []Option  `json:"Options,omitempty"`
	Mode     Mode      `json:"Mode,omitempty"`
}
