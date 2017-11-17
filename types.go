package languagetool

const (
	URL string = `https://languagetool.org/api/v2/check`
)

type Warnings struct {
	IncompleteResults bool `json:"incompleteResults"`
}

type Software struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	BuildDate  string `json:"buildDate"`
	APIVersion int    `json:"apiVersion"`
	Status     string `json:"status"`
}

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Result struct {
	Software Software `json:"software"`
	Warnings Warnings `json:"warnings"`
	Language Language `json:"language"`
	Matches  []Match  `json:"matches"`
}

type Match struct {
	Message      string        `json:"message"`
	ShortMessage string        `json:"shortMessage"`
	Replacements []Replacement `json:"replacements"`
	Offset       int           `json:"offset"`
	Length       int           `json:"length"`
	Context      Context       `json:"context"`
	Rule         Rule          `json:"rule"`
}

type Replacement struct {
	Value string `json:"value"`
}

type Context struct {
	Text   string `json:"text"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type Rule struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	IssueType   string   `json:"issueType"`
	Category    Category `json:"category"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
