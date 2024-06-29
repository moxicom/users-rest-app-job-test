package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/moxicom/user_test/internal/models"
)

var ApiAddress string

func GetUserData(passportInfo string) (models.User, error) {
	serie := passportInfo[:4]
	number := passportInfo[5:]

	u, err := url.Parse(ApiAddress)
	if err != nil {
		return models.User{}, err
	}

	// Add query parameters
	q := u.Query()
	q.Set("passportSerie", fmt.Sprintf("%v", serie))
	q.Set("passportNumber", fmt.Sprintf("%v", number))
	u.RawQuery = q.Encode()

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return models.User{}, err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return models.User{}, fmt.Errorf("status code is not OK: %v", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, err
	}

	// Parse the JSON response
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
