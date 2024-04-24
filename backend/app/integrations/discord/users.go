package discord

import (
	"codejam.io/logging"
	"encoding/json"
	"net/http"
)

var logger = logging.NewLogger(logging.Options{Name: "Discord", Level: logging.DEBUG})

// GetUser calls the Discord API to get the user associated with the access token.
// Tried using a Discord API lib but would just get 401 errors back trying to call users/@me,
// so I just rolled my own call
func GetUser(token string) map[string]interface{} {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		logger.Error("error creating request: %v", err)
		return nil
	}

	req.Header.Add("Authorization", "Bearer "+token)
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
