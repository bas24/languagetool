package languagetool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// CheckWithURL runs an api call for a custom languagetool server
func CheckWithURL(text, lang, apiURL string) (Result, error) {
	data := url.Values{}
	data.Set("text", text)
	data.Set("language", lang)
	data.Set("enabledOnly", "false")
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return Result{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Result{}, err
	}

	if resp.StatusCode == 200 {
		return result, nil
	}

	// TODO handle other status codes

	return Result{}, nil
}

// Check runs an api call to the public LanguageTool API, to check the given text with the given lang
func Check(text, lang string) (Result, error) {
	return CheckWithURL(text, lang, URL)
}
