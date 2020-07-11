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

// CucumberReport : struct for cucumber report
type CucumberReport []struct {
	URI         string `json:"uri"`
	ID          string `json:"id"`
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Line        int    `json:"line"`
	Tags        []struct {
		Name string `json:"name"`
		Line int    `json:"line"`
	} `json:"tags"`
	Elements []struct {
		ID          string `json:"id"`
		Keyword     string `json:"keyword"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Line        int    `json:"line"`
		Type        string `json:"type"`
		Tags        []struct {
			Name string `json:"name"`
			Line int    `json:"line"`
		} `json:"tags"`
		Steps []struct {
			Keyword string `json:"keyword"`
			Name    string `json:"name"`
			Line    int    `json:"line"`
			Match   struct {
				Location string `json:"location"`
			} `json:"match"`
			Result struct {
				Status       string `json:"status"`
				ErrorMessage string `json:"error_message"`
				Duration     int64  `json:"duration"`
			} `json:"result"`
		} `json:"steps"`
	} `json:"elements,omitempty"`
}
