package structs

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
