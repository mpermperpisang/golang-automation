package support

type scenarioDetail struct {
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

type featureDetail struct {
	Location struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"location"`
	Type string `json:"type"`
	Tags []struct {
		Location struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"location"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"tags"`
	Language            string `json:"language"`
	Keyword             string `json:"keyword"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	ScenarioDefinitions []struct {
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
	} `json:"scenarioDefinitions"`
	Comments []interface{} `json:"comments"`
}
