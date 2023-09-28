package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var client *http.Client

func Init() {
	client = &http.Client{Timeout: 10 * time.Second}
}

func getJson(url string, target interface{}) error {

	err := godotenv.Load()

	if err != nil {
		return err
	}

	token := os.Getenv("TOKEN")

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(`Authorization`, fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
