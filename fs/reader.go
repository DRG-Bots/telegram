package fs

import (
	"encoding/json"
	"os"

	"github.com/drg-bots/telegram/models"
)

func readFromPath[T models.DRGData](path string) (*T, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data T
	err = json.NewDecoder(f).Decode(&data)
	return &data, err
}

func ReadSalutes() (*models.Salutes, error) {
	return readFromPath[models.Salutes]("./data/salutes.json")
}

func ReadTrivia() (*models.Trivia, error) {
	return readFromPath[models.Trivia]("./data/trivia.json")
}

func ReadDeepDives() (*models.DeepDives, error) {
	return readFromPath[models.DeepDives]("./data/deepdives.json")
}
