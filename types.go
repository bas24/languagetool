package languagetool

const (
	URL string = `https://languagetool.org/api/v2/check`
)

type Soft struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	BuildDate  string `json:"buildDate"`
	ApiVersion int    `json:"apiVersion"`
	Status     string `json:"status"`
}

type Lang struct {
	Name string `json:"name`
	Code string `json:"code"`
}

type Replacement struct {
	Value string `json:"value"`
}

type ContextVal struct {
	Text   string `json:"text"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type RuleUrls struct {
	Value string `json:"value"`
}

type RuleCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RuleVal struct {
	Id          string       `json:"id"`
	SubId       string       `json:"subId"`
	Description string       `json:"description"`
	Urls        RuleUrls     `json:"urls"`
	IssueType   string       `json:"issueType"`
	Category    RuleCategory `json:"category"`
}

type Match struct {
	Message      string        `json:"message"`
	ShortMessage string        `json:"shortMessage"`
	Offset       int           `json:"offset"`
	Length       int           `json:"length"`
	Replacements []Replacement `json:"replacements"`
	Context      ContextVal    `json:"context"`
	Rule         RuleVal       `json:"rule"`
}

type Result struct {
	Software Soft    `json:"software"`
	Language Lang    `json:"language"`
	Matches  []Match `json:"matches"`
}

// supported languages
type Languages struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	LongCode string `json:"longCode"`
}
