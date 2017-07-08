package languagetool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Check(text, lang string) (Result, error) {
	apiUrl := URL
	data := url.Values{}
	data.Set("text", text)
	data.Set("language", lang)
	data.Set("enabledOnly", "false")
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))
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
