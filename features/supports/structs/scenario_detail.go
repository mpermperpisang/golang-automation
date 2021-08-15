package structs

type ScenarioDetail struct {
	Location struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"location"`
	Type    string `json:"type"`
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
	Steps   []struct {
		Location struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"location"`
		Type    string `json:"type"`
		Keyword string `json:"keyword"`
		Text    string `json:"text"`
	} `json:"steps"`
	Tags []struct {
		Location struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"location"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"tags"`
}
