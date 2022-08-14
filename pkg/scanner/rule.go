package scanner

type Rule struct {
	ID          string `json:"ID,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Severity    string `json:"severity,omitempty"`
}

func NewRule(ID string, t string, description string, severity string) *Rule {
	return &Rule{ID: ID, Type: t, Description: description, Severity: severity}
}
