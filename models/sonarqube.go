package models

type TextRange struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartColumn int `json:"startColumn"`
	EndColumn   int `json:"endColumn"`
}
type Location struct {
	Message   string    `json:"message"`
	FilePath  string    `json:"filePath"`
	TextRange TextRange `json:"textRange"`
}

type SonarQubeIssue struct {
	EngineID           string               `json:"engineId"`
	RuleID             string               `json:"ruleId"`
	Severity           string               `json:"severity"`
	Type               string               `json:"type"`
	PrimaryLocation    Location      		`json:"primaryLocation,omitempty"`
	EffortMinutes      int                  `json:"effortMinutes,omitempty"`
	SecondaryLocations []Location 			`json:"secondaryLocations,omitempty"`
}


type SonarQubeReport struct {
	SonarQubeIssues []SonarQubeIssue `json:"issues"`
}
