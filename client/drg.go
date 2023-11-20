package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drg-bots/telegram/models"
)

func FetchSalutes() (*models.Salutes, error) {
	return fetch[models.Salutes]("salutes")
}

func FetchTrivia() (*models.Trivia, error) {
	return fetch[models.Trivia]("trivia")
}

func FetchDeepDives() (*models.DeepDives, error) {
	return fetch[models.DeepDives]("deepdives")
}

func fetch[T models.DRGData](endpoint string) (*T, error) {
	res, err := http.Get(fmt.Sprintf("https://drgapi.com/v1/%s", endpoint))
	if err != nil {
		return nil, err
	}

	var t T
	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
