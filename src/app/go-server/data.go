package  main


type Outcome struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Titles map[string]string `json:"titles"`
	Descriptions map[string]string `json:"descriptions"`
	SubOutcomes []Outcome `json:"suboutcomes"`
}

type Root struct {
    Outcomes []Outcome `json:"outcomes"`
}