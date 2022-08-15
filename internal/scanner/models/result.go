package models

type ScanResult struct {
	Type     string   `json:"type"`
	RuleId   string   `json:"ruleId"`
	Location Location `json:"location"`
	Metadata struct {
		Description string `json:"description"`
		Severity    string `json:"severity"`
	} `json:"metadata"`
}

type Location struct {
	Path      string   `json:"path"`
	Positions Position `json:"positions"`
}

type Position struct {
	Begin struct {
		Line int `json:"line"`
	} `json:"begin"`
	End struct {
		Line int `json:"line"`
	} `json:"end"`
}
