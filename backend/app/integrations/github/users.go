package github

import (
	"codejam.io/logging"
	"encoding/json"
	"net/http"
)

var logger = logging.NewLogger(logging.Options{Name: "GitHub", Level: logging.INFO})

func GetUser(token string) map[string]interface{} {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		logger.Error("error creating request: %v", err)
		return nil
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("response error: %v", err)
		return nil
	}

	defer resp.Body.Close()

	var j map[string]interface{}
	d := json.NewDecoder(resp.Body)
	d.UseNumber() // so the ID doesn't get converted to a float64
	err = d.Decode(&j)
	if err != nil {
		logger.Error("error parsing JSON: %v", err)
	}
	return j
}
